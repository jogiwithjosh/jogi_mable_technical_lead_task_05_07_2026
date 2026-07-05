package auth

import (
	"context"
	"time"

	"api/internal/dto"

	"github.com/google/uuid"
)

type Service interface {
	Signup(
		ctx context.Context,
		req dto.SignupRequest,
	) (*User, error)

	Login(
		ctx context.Context,
		req dto.LoginRequest,
	) (*User, string, error)

	GetUser(
		ctx context.Context,
		id string,
	) (*User, error)
}

type AuthService struct {
	repository Repository
	jwt        *JWTManager
}

func NewService(
	repository Repository,
	jwt *JWTManager,
) Service {
	return &AuthService{
		repository: repository,
		jwt:        jwt,
	}
}

func (s *AuthService) Signup(
	ctx context.Context,
	req dto.SignupRequest,
) (*User, error) {
	_, err := s.repository.GetByEmail(
		ctx,
		req.Email,
	)

	if err == nil {
		return nil, ErrUserAlreadyExists
	}

	passwordHash, err := Hash(
		req.Password,
	)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:           uuid.NewString(),
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now().UTC(),
	}

	if err := s.repository.Create(
		ctx,
		user,
	); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(
	ctx context.Context,
	req dto.LoginRequest,
) (*User, string, error) {
	user, err := s.repository.GetByEmail(
		ctx,
		req.Email,
	)
	if err != nil {
		return nil, "", ErrInvalidCredentials
	}

	if err := Verify(
		req.Password,
		user.PasswordHash,
	); err != nil {
		return nil, "", ErrInvalidCredentials
	}

	token, err := s.jwt.Generate(
		user.ID,
		user.Email,
	)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *AuthService) GetUser(
	ctx context.Context,
	id string,
) (*User, error) {
	return s.repository.GetByID(
		ctx,
		id,
	)
}
