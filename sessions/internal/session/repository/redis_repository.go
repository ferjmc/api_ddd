package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ferjmc/api_ddd/sessions/internal/models"
	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
)

// sessionRedisRepo
type sessionRedisRepo struct {
	redis      *redis.Client
	prefix     string
	expiration time.Duration
}

func NewSessionRedisRepo(redis *redis.Client, prefix string, expiration time.Duration) *sessionRedisRepo {
	return &sessionRedisRepo{redis: redis, prefix: prefix, expiration: expiration}
}

func (s *sessionRedisRepo) CreateSession(ctx context.Context, userID uuid.UUID) (*models.Session, error) {

	sess := &models.Session{
		SessionID: uuid.NewV4().String(),
		UserID:    userID,
	}

	sessBytes, err := json.Marshal(&sess)
	if err != nil {
		return nil, fmt.Errorf("sessionRepo.CreateSession.json.Marshal: %w", err)
	}

	if err := s.redis.SetEX(ctx, s.createKey(sess.SessionID), string(sessBytes), s.expiration).Err(); err != nil {
		return nil, fmt.Errorf("sessionRepo.CreateSession.redis.SetEX: %w", err)
	}

	return sess, nil
}

func (s *sessionRedisRepo) GetSessionByID(ctx context.Context, sessID string) (*models.Session, error) {
	result, err := s.redis.Get(ctx, s.createKey(sessID)).Result()
	if err != nil {
		return nil, fmt.Errorf("sessionRepo.GetSessionByID.redis.Get: %w", err)
	}

	var sess models.Session
	if err := json.Unmarshal([]byte(result), &sess); err != nil {
		return nil, fmt.Errorf("sessionRepo.GetSessionByID.json.Unmarshal: %w", err)
	}
	return &sess, nil
}

func (s *sessionRedisRepo) DeleteSession(ctx context.Context, sessID string) error {

	if err := s.redis.Del(ctx, s.createKey(sessID)).Err(); err != nil {
		return fmt.Errorf("sessionRepo.DeleteSession.redis.Del: %w", err)
	}
	return nil
}

func (s *sessionRedisRepo) createKey(sessionID string) string {
	return fmt.Sprintf("%s: %s", s.prefix, sessionID)
}
