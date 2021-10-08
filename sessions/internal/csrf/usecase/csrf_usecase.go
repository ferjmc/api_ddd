package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/ferjmc/sessions/internal/csrf"
)

const (
	CSRFHeader = "X-CSRF-Token"
	// 32 bytes
	csrfSalt = "KbWaoi5xtDC3GEfBa9ovQdzOzXsuVU9I"
)

// CSRF usecase
type CsrfUseCase struct {
	csrfRepo csrf.RedisRepository
}

// NewCsrfUC
func NewCsrfUseCase(csrfRepo csrf.RedisRepository) *CsrfUseCase {
	return &CsrfUseCase{csrfRepo: csrfRepo}
}

// Create new CSRF token
func (c *CsrfUseCase) GetCSRFToken(ctx context.Context, sesID string) (string, error) {

	token, err := c.makeToken(sesID)
	if err != nil {
		return "", fmt.Errorf("CsrfUseCase.CreateToken.c.makeToken: %w", err)
	}

	if err := c.csrfRepo.Create(ctx, token); err != nil {
		return "", fmt.Errorf("CsrfUseCase.CreateToken.csrfRepo.Create: %w", err)
	}

	return token, nil
}

// Validate csrf token using session id and token
func (c *CsrfUseCase) ValidateCSRFToken(ctx context.Context, sesID string, token string) (bool, error) {
	existsToken, err := c.csrfRepo.GetToken(ctx, token)
	if err != nil {
		return false, err
	}

	return c.validateToken(existsToken, sesID), nil
}

func (c *CsrfUseCase) makeToken(sessionID string) (string, error) {
	hash := sha256.New()
	_, err := io.WriteString(hash, csrfSalt+sessionID)
	if err != nil {
		return "", err
	}
	token := base64.RawStdEncoding.EncodeToString(hash.Sum(nil))
	return token, nil
}

// Validate CSRF token
func (c *CsrfUseCase) validateToken(token string, sessionID string) bool {
	trueToken, err := c.makeToken(sessionID)
	if err != nil {
		return false
	}
	return token == trueToken
}
