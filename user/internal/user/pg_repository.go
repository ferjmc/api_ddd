package user

import (
	"context"

	"github.com/ferjmc/user/internal/models"
	uuid "github.com/satori/go.uuid"
)

// PGRepository
type PGRepository interface {
	Create(ctx context.Context, user *models.User) (*models.UserResponse, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*models.UserResponse, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, user *models.UserUpdate) (*models.UserResponse, error)
	//	UpdateAvatar(ctx context.Context, msg models.UploadedImageMsg) (*models.UserResponse, error)
	GetUsersByIDs(ctx context.Context, userIDs []string) ([]*models.UserResponse, error)
}
