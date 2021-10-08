package models

import uuid "github.com/satori/go.uuid"

// Session
type Session struct {
	UserID    uuid.UUID `json:"user_id"`
	SessionID string    `json:"session_id"`
}
