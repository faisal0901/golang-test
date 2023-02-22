package Test

import (
	"context"
	"errors"
	"go-test/Model"
	"go-test/Services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TokenRepoMock struct {
	mock.Mock
}

// CreateData implements Repository.IRepository
func (*TokenRepoMock) CreateData(ctx context.Context, data interface{}) error {
	panic("unimplemented")
}

// Delete implements Repository.IRepository
func (*TokenRepoMock) Delete(ctx context.Context, id uint) error {
	panic("unimplemented")
}

// GetAll implements Repository.IRepository
func (*TokenRepoMock) GetAll(ctx context.Context, results interface{}) error {
	panic("unimplemented")
}

// GetAllJoin implements Repository.IRepository
func (*TokenRepoMock) GetAllJoin(ctx context.Context, results interface{}) error {
	panic("unimplemented")
}

// GetAllProduct implements Repository.IRepository
func (*TokenRepoMock) GetAllProduct(ctx context.Context, results interface{}) error {
	panic("unimplemented")
}

// GetByEmail implements Repository.IRepository
func (*TokenRepoMock) GetByEmail(ctx context.Context, email string, result interface{}) error {
	panic("unimplemented")
}

// GetByID implements Repository.IRepository
func (*TokenRepoMock) GetByID(ctx context.Context, id uint, result interface{}) error {
	panic("unimplemented")
}

// Update implements Repository.IRepository
func (*TokenRepoMock) Update(ctx context.Context, id uint, data interface{}) (interface{}, error) {
	panic("unimplemented")
}

func (m *TokenRepoMock) Create(ctx context.Context, data interface{}) (interface{}, error) {
	args := m.Called(ctx, data)
	return args.Get(0), args.Error(1)
}

func (m *TokenRepoMock) UpdateToken(ctx context.Context, token string, data interface{}) (interface{}, error) {
	args := m.Called(ctx, token, data)
	return args.Get(0), args.Error(1)
}

func TestCreateNewToken(t *testing.T) {
	// Setup
	mockTokenRepo := new(TokenRepoMock)
	tokenService := Services.NewTokenService(mockTokenRepo)

	ctx := context.Background()
	id := uint(1)
	token := "token"

	// Expectation
	mockTokenRepo.On("Create", ctx, &Model.Token{
		Token:      token,
		IsValid:    1,
		CustomerID: id,
	}).Return(&Model.Token{ID: 1, Token: token, IsValid: 1, CustomerID: id}, nil)

	// Execution
	result, err := tokenService.CreateNewToken(ctx, id, token)

	// Assertion
	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockTokenRepo.AssertExpectations(t)
}
func TestDisableToken(t *testing.T) {
	// Setup
	mockTokenRepo := new(TokenRepoMock)
	tokenService := Services.NewTokenService(mockTokenRepo)

	ctx := context.Background()
	id := uint(1)
	token := "token"

	// Expectation
	mockTokenRepo.On("UpdateToken", ctx, token, &Model.Token{}).Return(&Model.Token{ID: 1, Token: token, IsValid: 0, CustomerID: id}, nil)

	// Execution
	result, err := tokenService.DisableToken(ctx, token, id)

	// Assertion
	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockTokenRepo.AssertExpectations(t)
}
func TestDisableToken_Error(t *testing.T) {
	// Setup
	mockTokenRepo := new(TokenRepoMock)
	tokenService := Services.NewTokenService(mockTokenRepo)

	ctx := context.Background()
	id := uint(1)
	token := "token"

	// Expectation
	mockTokenRepo.On("UpdateToken", ctx, token, &Model.Token{}).Return(nil, errors.New("failed to update token"))

	// Execution
	result, err := tokenService.DisableToken(ctx, token, id)

	// Assertion
	assert.Error(t, err)
	assert.Nil(t, result)
	mockTokenRepo.AssertExpectations(t)
}