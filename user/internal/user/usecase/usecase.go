package usecase

import (
	"context"
	"fmt"

	"github.com/ferjmc/api_ddd/user/internal/models"
	"github.com/ferjmc/api_ddd/user/internal/user"
	"github.com/ferjmc/api_ddd/user/pkg/logger"
	sessionService "github.com/ferjmc/api_ddd/user/proto/session"
	uuid "github.com/satori/go.uuid"
)

const (
	imagesExchange = "images"
	resizeKey      = "resize_image_key"
	userUUIDHeader = "user_uuid"
)

type userUseCase struct {
	userPGRepo user.PGRepository
	sessClient sessionService.AuthorizationServiceClient
	log        logger.Logger
}

func NewUserUseCase(
	userPGRepo user.PGRepository,
	sessClient sessionService.AuthorizationServiceClient,
	log logger.Logger,
) *userUseCase {
	return &userUseCase{
		userPGRepo: userPGRepo,
		sessClient: sessClient,
		log:        log,
	}
}

func (u *userUseCase) GetByID(ctx context.Context, userID uuid.UUID) (*models.UserResponse, error) {

	userResponse, err := u.userPGRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("userUseCase.userPGRepo.GetByID: %w", err)
	}

	return userResponse, nil
}

func (u *userUseCase) Register(ctx context.Context, user *models.User) (*models.UserResponse, error) {

	if err := user.PrepareCreate(); err != nil {
		return nil, fmt.Errorf("user.PrepareCreate: %w", err)
	}

	created, err := u.userPGRepo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("userPGRepo.Create: %w", err)
	}

	return created, err
}

func (u *userUseCase) Login(ctx context.Context, login models.Login) (*models.User, error) {

	userByEmail, err := u.userPGRepo.GetByEmail(ctx, login.Email)
	if err != nil {
		return nil, fmt.Errorf("userPGRepo.GetByEmail: %w", err)
	}

	if err := userByEmail.ComparePasswords(login.Password); err != nil {
		return nil, fmt.Errorf("userUseCase.ComparePasswords: %w", err)
	}

	userByEmail.SanitizePassword()

	return userByEmail, nil
}

func (u *userUseCase) CreateSession(ctx context.Context, userID uuid.UUID) (string, error) {
	session, err := u.sessClient.CreateSession(ctx, &sessionService.CreateSessionRequest{UserID: userID.String()})
	if err != nil {
		return "", fmt.Errorf("sessClient.CreateSession: %w", err)
	}

	return session.GetSession().GetSessionID(), err
}

func (u *userUseCase) DeleteSession(ctx context.Context, sessionID string) error {
	_, err := u.sessClient.DeleteSession(ctx, &sessionService.DeleteSessionRequest{SessionID: sessionID})
	if err != nil {
		return fmt.Errorf("sessClient.DeleteSession: %w", err)
	}

	return nil
}

func (u *userUseCase) GetSessionByID(ctx context.Context, sessionID string) (*models.Session, error) {

	sessionByID, err := u.sessClient.GetSessionByID(ctx, &sessionService.GetSessionByIDRequest{SessionID: sessionID})
	if err != nil {
		return nil, fmt.Errorf("sessClient.GetSessionByID: %w", err)
	}

	sess := &models.Session{}
	sess, err = sess.FromProto(sessionByID.GetSession())
	if err != nil {
		return nil, fmt.Errorf("sess.FromProto: %w", err)
	}

	return sess, nil
}

func (u *userUseCase) GetCSRFToken(ctx context.Context, sessionID string) (string, error) {
	csrfToken, err := u.sessClient.CreateCsrfToken(
		ctx,
		&sessionService.CreateCsrfTokenRequest{CsrfTokenInput: &sessionService.CsrfTokenInput{SessionID: sessionID}},
	)
	if err != nil {
		return "", fmt.Errorf("sessClient.CreateCsrfToken: %w", err)
	}

	return csrfToken.GetCsrfToken().GetToken(), nil
}

func (u *userUseCase) Update(ctx context.Context, user *models.UserUpdate) (*models.UserResponse, error) {

	userResponse, err := u.userPGRepo.Update(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("userUseCase.Update.userPGRepo.Update: %w", err)
	}

	return userResponse, nil
}

func (u *userUseCase) GetUsersByIDs(ctx context.Context, userIDs []string) ([]*models.UserResponse, error) {

	return u.userPGRepo.GetUsersByIDs(ctx, userIDs)
}
