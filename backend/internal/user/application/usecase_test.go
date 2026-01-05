package application

import (
	"context"
	"errors"
	"testing"

	"github.com/bryanriosb/stock-info/internal/user/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByID(ctx context.Context, id int64) (*domain.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserRepository) FindAll(ctx context.Context) ([]*domain.User, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.User), args.Error(1)
}

func (m *MockUserRepository) CountByRole(ctx context.Context, role domain.Role) (int64, error) {
	args := m.Called(ctx, role)
	return args.Get(0).(int64), args.Error(1)
}

func TestCreate_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)

	mockRepo.On("FindByUsername", mock.Anything, "testuser").Return(nil, errors.New("not found"))
	mockRepo.On("FindByEmail", mock.Anything, "test@example.com").Return(nil, errors.New("not found"))
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)

	uc := NewUserUseCase(mockRepo)
	user, err := uc.Create(context.Background(), CreateUserRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	})

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "test@example.com", user.Email)
	mockRepo.AssertExpectations(t)
}

func TestCreate_UsernameExists(t *testing.T) {
	mockRepo := new(MockUserRepository)

	existingUser := &domain.User{ID: 1, Username: "testuser"}
	mockRepo.On("FindByUsername", mock.Anything, "testuser").Return(existingUser, nil)

	uc := NewUserUseCase(mockRepo)
	user, err := uc.Create(context.Background(), CreateUserRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	})

	assert.Error(t, err)
	assert.Equal(t, ErrUserAlreadyExists, err)
	assert.Nil(t, user)
	mockRepo.AssertExpectations(t)
}

func TestCreate_EmailExists(t *testing.T) {
	mockRepo := new(MockUserRepository)

	existingUser := &domain.User{ID: 1, Email: "test@example.com"}
	mockRepo.On("FindByUsername", mock.Anything, "newuser").Return(nil, errors.New("not found"))
	mockRepo.On("FindByEmail", mock.Anything, "test@example.com").Return(existingUser, nil)

	uc := NewUserUseCase(mockRepo)
	user, err := uc.Create(context.Background(), CreateUserRequest{
		Username: "newuser",
		Email:    "test@example.com",
		Password: "password123",
	})

	assert.Error(t, err)
	assert.Equal(t, ErrUserAlreadyExists, err)
	assert.Nil(t, user)
	mockRepo.AssertExpectations(t)
}

func TestGetByID_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)

	expectedUser := &domain.User{ID: 1, Username: "testuser", Email: "test@example.com"}
	mockRepo.On("FindByID", mock.Anything, int64(1)).Return(expectedUser, nil)

	uc := NewUserUseCase(mockRepo)
	user, err := uc.GetByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, int64(1), user.ID)
	mockRepo.AssertExpectations(t)
}

func TestGetByID_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)

	mockRepo.On("FindByID", mock.Anything, int64(999)).Return(nil, errors.New("not found"))

	uc := NewUserUseCase(mockRepo)
	user, err := uc.GetByID(context.Background(), 999)

	assert.Error(t, err)
	assert.Equal(t, ErrUserNotFound, err)
	assert.Nil(t, user)
	mockRepo.AssertExpectations(t)
}

func TestGetAll_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)

	users := []*domain.User{
		{ID: 1, Username: "user1"},
		{ID: 2, Username: "user2"},
	}
	mockRepo.On("FindAll", mock.Anything).Return(users, nil)

	uc := NewUserUseCase(mockRepo)
	result, err := uc.GetAll(context.Background())

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

func TestUpdate_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)

	existingUser := &domain.User{ID: 1, Username: "olduser", Email: "old@example.com"}
	mockRepo.On("FindByID", mock.Anything, int64(1)).Return(existingUser, nil)
	mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)

	uc := NewUserUseCase(mockRepo)
	user, err := uc.Update(context.Background(), 1, UpdateUserRequest{
		Username: "newuser",
		Email:    "new@example.com",
	})

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "newuser", user.Username)
	assert.Equal(t, "new@example.com", user.Email)
	mockRepo.AssertExpectations(t)
}

func TestUpdate_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)

	mockRepo.On("FindByID", mock.Anything, int64(999)).Return(nil, errors.New("not found"))

	uc := NewUserUseCase(mockRepo)
	user, err := uc.Update(context.Background(), 999, UpdateUserRequest{Username: "newuser"})

	assert.Error(t, err)
	assert.Equal(t, ErrUserNotFound, err)
	assert.Nil(t, user)
	mockRepo.AssertExpectations(t)
}

func TestUpdate_PartialUpdate(t *testing.T) {
	mockRepo := new(MockUserRepository)

	existingUser := &domain.User{ID: 1, Username: "olduser", Email: "old@example.com"}
	mockRepo.On("FindByID", mock.Anything, int64(1)).Return(existingUser, nil)
	mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)

	uc := NewUserUseCase(mockRepo)
	user, err := uc.Update(context.Background(), 1, UpdateUserRequest{
		Username: "newuser", // Only update username
	})

	assert.NoError(t, err)
	assert.Equal(t, "newuser", user.Username)
	assert.Equal(t, "old@example.com", user.Email) // Email unchanged
	mockRepo.AssertExpectations(t)
}

func TestDelete_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)

	existingUser := &domain.User{ID: 1, Username: "testuser"}
	mockRepo.On("FindByID", mock.Anything, int64(1)).Return(existingUser, nil)
	mockRepo.On("Delete", mock.Anything, int64(1)).Return(nil)

	uc := NewUserUseCase(mockRepo)
	err := uc.Delete(context.Background(), 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDelete_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)

	mockRepo.On("FindByID", mock.Anything, int64(999)).Return(nil, errors.New("not found"))

	uc := NewUserUseCase(mockRepo)
	err := uc.Delete(context.Background(), 999)

	assert.Error(t, err)
	assert.Equal(t, ErrUserNotFound, err)
	mockRepo.AssertExpectations(t)
}

func TestDelete_LastAdmin(t *testing.T) {
	mockRepo := new(MockUserRepository)

	adminUser := &domain.User{ID: 1, Username: "admin", Role: domain.RoleAdmin}
	mockRepo.On("FindByID", mock.Anything, int64(1)).Return(adminUser, nil)
	mockRepo.On("CountByRole", mock.Anything, domain.RoleAdmin).Return(int64(1), nil)

	uc := NewUserUseCase(mockRepo)
	err := uc.Delete(context.Background(), 1)

	assert.Error(t, err)
	assert.Equal(t, ErrLastAdmin, err)
	mockRepo.AssertExpectations(t)
}

func TestDelete_AdminWithOtherAdmins(t *testing.T) {
	mockRepo := new(MockUserRepository)

	adminUser := &domain.User{ID: 1, Username: "admin", Role: domain.RoleAdmin}
	mockRepo.On("FindByID", mock.Anything, int64(1)).Return(adminUser, nil)
	mockRepo.On("CountByRole", mock.Anything, domain.RoleAdmin).Return(int64(2), nil)
	mockRepo.On("Delete", mock.Anything, int64(1)).Return(nil)

	uc := NewUserUseCase(mockRepo)
	err := uc.Delete(context.Background(), 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdate_ChangeRoleLastAdmin(t *testing.T) {
	mockRepo := new(MockUserRepository)

	adminUser := &domain.User{ID: 1, Username: "admin", Email: "admin@example.com", Role: domain.RoleAdmin}
	mockRepo.On("FindByID", mock.Anything, int64(1)).Return(adminUser, nil)
	mockRepo.On("CountByRole", mock.Anything, domain.RoleAdmin).Return(int64(1), nil)

	uc := NewUserUseCase(mockRepo)
	user, err := uc.Update(context.Background(), 1, UpdateUserRequest{
		Role: string(domain.RoleUser),
	})

	assert.Error(t, err)
	assert.Equal(t, ErrLastAdmin, err)
	assert.Nil(t, user)
	mockRepo.AssertExpectations(t)
}

func TestUpdate_ChangeRoleWithOtherAdmins(t *testing.T) {
	mockRepo := new(MockUserRepository)

	adminUser := &domain.User{ID: 1, Username: "admin", Email: "admin@example.com", Role: domain.RoleAdmin}
	mockRepo.On("FindByID", mock.Anything, int64(1)).Return(adminUser, nil)
	mockRepo.On("CountByRole", mock.Anything, domain.RoleAdmin).Return(int64(2), nil)
	mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)

	uc := NewUserUseCase(mockRepo)
	user, err := uc.Update(context.Background(), 1, UpdateUserRequest{
		Role: string(domain.RoleUser),
	})

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, domain.RoleUser, user.Role)
	mockRepo.AssertExpectations(t)
}

func TestAuthenticate_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)

	user := &domain.User{ID: 1, Username: "testuser"}
	user.SetPassword("password123")

	mockRepo.On("FindByUsername", mock.Anything, "testuser").Return(user, nil)

	uc := NewUserUseCase(mockRepo)
	result, err := uc.Authenticate(context.Background(), "testuser", "password123")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "testuser", result.Username)
	mockRepo.AssertExpectations(t)
}

func TestAuthenticate_UserNotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)

	mockRepo.On("FindByUsername", mock.Anything, "nonexistent").Return(nil, errors.New("not found"))

	uc := NewUserUseCase(mockRepo)
	result, err := uc.Authenticate(context.Background(), "nonexistent", "password123")

	assert.Error(t, err)
	assert.Equal(t, ErrInvalidCredentials, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestAuthenticate_WrongPassword(t *testing.T) {
	mockRepo := new(MockUserRepository)

	user := &domain.User{ID: 1, Username: "testuser"}
	user.SetPassword("correctpassword")

	mockRepo.On("FindByUsername", mock.Anything, "testuser").Return(user, nil)

	uc := NewUserUseCase(mockRepo)
	result, err := uc.Authenticate(context.Background(), "testuser", "wrongpassword")

	assert.Error(t, err)
	assert.Equal(t, ErrInvalidCredentials, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
