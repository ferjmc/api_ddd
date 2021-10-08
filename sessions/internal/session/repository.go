package session

import (
	"context"

	"github.com/ferjmc/sessions/internal/models"
	uuid "github.com/satori/go.uuid"
)

// Session RedisRepository
type RedisRepository interface {
	CreateSession(ctx context.Context, userID uuid.UUID) (*models.Session, error)
	GetSessionByID(ctx context.Context, sessID string) (*models.Session, error)
	DeleteSession(ctx context.Context, sessID string) error
}
