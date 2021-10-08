package usecase

import (
	"context"
	"fmt"

	"github.com/ferjmc/user/internal/models"
	"github.com/ferjmc/user/internal/user"
	"github.com/ferjmc/user/pkg/logger"
	uuid "github.com/satori/go.uuid"
)

const (
	imagesExchange = "images"
	resizeKey      = "resize_image_key"
	userUUIDHeader = "user_uuid"
)

type userUseCase struct {
	userPGRepo user.PGRepository
	log        logger.Logger
}

func NewUserUseCase(
	userPGRepo user.PGRepository,
	log logger.Logger,
) *userUseCase {
	return &userUseCase{
		userPGRepo: userPGRepo,
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

	return "", nil
}

func (u *userUseCase) DeleteSession(ctx context.Context, sessionID string) error {

	return nil
}

func (u *userUseCase) GetSessionByID(ctx context.Context, sessionID string) (*models.Session, error) {

	sess := &models.Session{}

	return sess, nil
}

func (u *userUseCase) GetCSRFToken(ctx context.Context, sessionID string) (string, error) {

	return "", nil
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
