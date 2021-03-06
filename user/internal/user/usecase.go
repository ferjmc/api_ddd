package user

import (
	"context"

	"github.com/ferjmc/api_ddd/user/internal/models"
	uuid "github.com/satori/go.uuid"
)

// UseCase
type UseCase interface {
	Register(ctx context.Context, user *models.User) (*models.UserResponse, error)
	Login(ctx context.Context, login models.Login) (*models.User, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*models.UserResponse, error)
	CreateSession(ctx context.Context, userID uuid.UUID) (string, error)
	GetSessionByID(ctx context.Context, sessionID string) (*models.Session, error)
	GetCSRFToken(ctx context.Context, sessionID string) (string, error)
	DeleteSession(ctx context.Context, sessionID string) error
	Update(ctx context.Context, user *models.UserUpdate) (*models.UserResponse, error)
	//UpdateUploadedAvatar(ctx context.Context, delivery amqp.Delivery) error
	//UpdateAvatar(ctx context.Context, data *models.UpdateAvatarMsg) error
	GetUsersByIDs(ctx context.Context, userIDs []string) ([]*models.UserResponse, error)
}
