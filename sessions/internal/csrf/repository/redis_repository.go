package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// CsrfRepository
type CsrfRepository struct {
	redis    *redis.Client
	prefix   string
	duration time.Duration
}

// NewRedisRepository
func NewCsrfRepository(redis *redis.Client, prefix string, duration time.Duration) *CsrfRepository {
	return &CsrfRepository{redis: redis, prefix: prefix, duration: duration}
}

// Create csrf token
func (r *CsrfRepository) Create(ctx context.Context, token string) error {

	if err := r.redis.SetEX(ctx, r.createKey(token), token, r.duration).Err(); err != nil {
		return fmt.Errorf("CsrfRepository.Create.redis.SetEX: %w", err)
	}

	return nil
}

// Check csrf token
func (r *CsrfRepository) GetToken(ctx context.Context, token string) (string, error) {

	token, err := r.redis.Get(ctx, r.createKey(token)).Result()
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *CsrfRepository) createKey(token string) string {
	return fmt.Sprintf("%s: %s", r.prefix, token)
}
