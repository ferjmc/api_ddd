package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ferjmc/user/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
	uuid "github.com/satori/go.uuid"
)

type userPGRepository struct {
	db *pgxpool.Pool
}

func NewUserPGRepository(db *pgxpool.Pool) *userPGRepository {
	return &userPGRepository{db: db}
}

// Create new user
func (u *userPGRepository) Create(ctx context.Context, user *models.User) (*models.UserResponse, error) {

	var created models.UserResponse
	if err := u.db.QueryRow(
		ctx,
		createUserQuery,
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Avatar,
		&user.Role,
	).Scan(&created.ID, &created.FirstName, &created.LastName, &created.Email,
		&created.Avatar, &created.Role, &created.UpdatedAt, &created.CreatedAt,
	); err != nil {
		return nil, fmt.Errorf("Create.Scan: %w", err)
	}

	return &created, nil
}

// Get user by id
func (u *userPGRepository) GetByID(ctx context.Context, userID uuid.UUID) (*models.UserResponse, error) {

	var res models.UserResponse
	if err := u.db.QueryRow(ctx, getUserByIDQuery, userID).Scan(
		&res.ID,
		&res.FirstName,
		&res.LastName,
		&res.Email,
		&res.Avatar,
		&res.Role,
		&res.UpdatedAt,
		&res.CreatedAt,
	); err != nil {
		return nil, fmt.Errorf("GetByID.Scan: %w", err)
	}

	return &res, nil
}

// GetByEmail
func (u *userPGRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {

	var res models.User
	if err := u.db.QueryRow(ctx, getUserByEmail, email).Scan(
		&res.ID,
		&res.FirstName,
		&res.LastName,
		&res.Email,
		&res.Password,
		&res.Avatar,
		&res.Role,
		&res.UpdatedAt,
		&res.CreatedAt,
	); err != nil {
		return nil, fmt.Errorf("GetByEmail.Scan: %w", err)
	}

	return &res, nil
}

// Update
func (u *userPGRepository) Update(ctx context.Context, user *models.UserUpdate) (*models.UserResponse, error) {

	var res models.UserResponse
	if err := u.db.QueryRow(ctx, updateUserQuery, &user.FirstName, &user.LastName, &user.Email, &user.Role, &user.ID).
		Scan(
			&res.ID,
			&res.FirstName,
			&res.LastName,
			&res.Email,
			&res.Role,
			&res.Avatar,
			&res.UpdatedAt,
			&res.CreatedAt,
		); err != nil {
		return nil, fmt.Errorf("Update.Scan: %w", err)
	}

	return &res, nil
}

func (u *userPGRepository) UpdateAvatar(ctx context.Context, msg models.UploadedImageMsg) (*models.UserResponse, error) {

	var res models.UserResponse
	if err := u.db.QueryRow(ctx, updateAvatarQuery, &msg.ImageURL, &msg.UserID).Scan(
		&res.ID,
		&res.FirstName,
		&res.LastName,
		&res.Email,
		&res.Role,
		&res.Avatar,
		&res.UpdatedAt,
		&res.CreatedAt,
	); err != nil {
		return nil, fmt.Errorf("UpdateAvatar.Scan: %w", err)
	}

	return &res, nil
}

func (u *userPGRepository) GetUsersByIDs(ctx context.Context, userIDs []string) ([]*models.UserResponse, error) {

	placeholders := CreateSQLPlaceholders(len(userIDs))
	query := fmt.Sprintf("SELECT user_id, first_name, last_name, email, avatar, role, updated_at, created_at FROM users WHERE user_id IN (%v)", placeholders)

	args := make([]interface{}, len(userIDs))
	for i, id := range userIDs {
		args[i] = id
	}

	rows, err := u.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("GetUserByUDs.db.Query: %w", err)
	}
	defer rows.Close()

	users := make([]*models.UserResponse, 0, len(userIDs))
	for rows.Next() {
		var res models.UserResponse
		if err := rows.Scan(
			&res.ID,
			&res.FirstName,
			&res.LastName,
			&res.Email,
			&res.Avatar,
			&res.Role,
			&res.UpdatedAt,
			&res.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("GetUserByUDs.db.Query: %w", err)
		}
		users = append(users, &res)
	}

	return users, nil
}

// CreateSQLPlaceholders Generate postgres $ placeholders
func CreateSQLPlaceholders(length int) string {
	var placeholders string
	for i := 0; i < length; i++ {
		placeholders += `$` + strconv.Itoa(i+1) + `,`

	}
	placeholders = placeholders[:len(placeholders)-1]
	return placeholders
}
