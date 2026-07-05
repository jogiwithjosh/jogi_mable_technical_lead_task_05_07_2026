package auth

import (
	"context"
	"database/sql"
	"errors"

	"api/internal/database"
)

type Repository interface {
	Create(
		ctx context.Context,
		user *User,
	) error

	GetByEmail(
		ctx context.Context,
		email string,
	) (*User, error)

	GetByID(
		ctx context.Context,
		id string,
	) (*User, error)
}

type ClickHouseRepository struct {
	db *sql.DB
}

func NewClickHouseRepository(
	ch *database.ClickHouse,
) *ClickHouseRepository {
	return &ClickHouseRepository{
		db: ch.DB,
	}
}

func (r *ClickHouseRepository) Create(
	ctx context.Context,
	user *User,
) error {
	_, err := r.GetByEmail(
		ctx,
		user.Email,
	)

	if err == nil {
		return ErrUserAlreadyExists
	}

	if !errors.Is(err, ErrUserNotFound) {
		return err
	}

	query := `
	INSERT INTO users
	(
		id,
		email,
		password_hash,
		name,
		created_at
	)
	VALUES
	(
		?, ?, ?, ?, ?
	)
	`

	_, err = r.db.ExecContext(
		ctx,
		query,
		user.ID,
		user.Email,
		user.PasswordHash,
		user.Name,
		user.CreatedAt,
	)
	return err
}

func (r *ClickHouseRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*User, error) {
	query := `
	SELECT
		id,
		email,
		password_hash,
		name,
		created_at
	FROM users
	WHERE email = ?
	LIMIT 1
	`

	var user User

	err := r.db.QueryRowContext(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Name,
		&user.CreatedAt,
	)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (r *ClickHouseRepository) GetByID(
	ctx context.Context,
	id string,
) (*User, error) {
	query := `
	SELECT
		id,
		email,
		password_hash,
		name,
		created_at,
	FROM users
	WHERE id = ?
	LIMIT 1
	`

	var user User

	err := r.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Name,
		&user.CreatedAt,
	)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}
