package application

import (
	"context"
	"errors"

	"github.com/bryanriosb/stock-info/internal/user/domain"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserUseCase interface {
	Create(ctx context.Context, req CreateUserRequest) (*domain.User, error)
	GetByID(ctx context.Context, id int64) (*domain.User, error)
	GetAll(ctx context.Context) ([]*domain.User, error)
	Update(ctx context.Context, id int64, req UpdateUserRequest) (*domain.User, error)
	Delete(ctx context.Context, id int64) error
	Authenticate(ctx context.Context, username, password string) (*domain.User, error)
}

type CreateUserRequest struct {
	Username string
	Email    string
	Password string
}

type UpdateUserRequest struct {
	Username string
	Email    string
	Password string
}

type userUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}

func (uc *userUseCase) Create(ctx context.Context, req CreateUserRequest) (*domain.User, error) {
	existing, _ := uc.repo.FindByUsername(ctx, req.Username)
	if existing != nil {
		return nil, ErrUserAlreadyExists
	}

	existing, _ = uc.repo.FindByEmail(ctx, req.Email)
	if existing != nil {
		return nil, ErrUserAlreadyExists
	}

	user := &domain.User{
		Username: req.Username,
		Email:    req.Email,
		Role:     domain.RoleUser,
	}

	if err := user.SetPassword(req.Password); err != nil {
		return nil, err
	}

	if err := uc.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUseCase) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	user, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (uc *userUseCase) GetAll(ctx context.Context) ([]*domain.User, error) {
	return uc.repo.FindAll(ctx)
}

func (uc *userUseCase) Update(ctx context.Context, id int64, req UpdateUserRequest) (*domain.User, error) {
	user, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrUserNotFound
	}

	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		if err := user.SetPassword(req.Password); err != nil {
			return nil, err
		}
	}

	if err := uc.repo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUseCase) Delete(ctx context.Context, id int64) error {
	_, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return ErrUserNotFound
	}
	return uc.repo.Delete(ctx, id)
}

func (uc *userUseCase) Authenticate(ctx context.Context, username, password string) (*domain.User, error) {
	user, err := uc.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if !user.CheckPassword(password) {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
