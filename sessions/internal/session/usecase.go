package session

import (
	"context"

	"github.com/ferjmc/api_ddd/sessions/internal/models"
	uuid "github.com/satori/go.uuid"
)

// Session SessUseCase
type SessUseCase interface {
	CreateSession(ctx context.Context, userID uuid.UUID) (*models.Session, error)
	GetSessionByID(ctx context.Context, sessID string) (*models.Session, error)
	DeleteSession(ctx context.Context, sessID string) error
}
