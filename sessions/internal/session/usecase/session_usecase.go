package usecase

import (
	"context"

	"github.com/ferjmc/sessions/internal/models"
	"github.com/ferjmc/sessions/internal/session"
	uuid "github.com/satori/go.uuid"
)

type sessionUseCase struct {
	sessRepo session.RedisRepository
}

func NewSessionUseCase(sessRepo session.RedisRepository) *sessionUseCase {
	return &sessionUseCase{sessRepo: sessRepo}
}

func (s *sessionUseCase) CreateSession(ctx context.Context, userID uuid.UUID) (*models.Session, error) {
	return s.sessRepo.CreateSession(ctx, userID)
}

func (s *sessionUseCase) GetSessionByID(ctx context.Context, sessID string) (*models.Session, error) {
	return s.sessRepo.GetSessionByID(ctx, sessID)
}

func (s *sessionUseCase) DeleteSession(ctx context.Context, sessID string) error {
	return s.sessRepo.DeleteSession(ctx, sessID)
}
