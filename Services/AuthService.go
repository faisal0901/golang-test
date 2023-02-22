package Services

import (
	"context"
	"errors"
	"net/http"

	model "go-test/Model"
	repository "go-test/Repository"
	response "go-test/Response"
	security "go-test/Security"
	"time"

	bcrypt "golang.org/x/crypto/bcrypt"
)


type IAuthService interface {
	CreateCustomer(ctx context.Context, customer *model.Customer) (interface{}, error)
	LoginCustomer(ctx context.Context, customer *model.Customer)(*response.LoginResponse,  ErrorResponse)
	LogoutCustomer(ctx context.Context,token string,id uint)(interface{},error)
}
type AuthService struct {
	authRepo repository.IRepository
	UserLogService ILogService
	TokenService ITokenService
}



func NewAuthService(authRepo repository.IRepository, userLogService ILogService,tokenService ITokenService) *AuthService {
	return &AuthService{authRepo: authRepo,
	UserLogService:userLogService,
	TokenService: tokenService,
	}
}
func (u *AuthService) CreateCustomer(ctx context.Context, customer *model.Customer) (interface{}, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	
	customer.Password = string(hashedPassword)
	err = u.CreateNewLog(err, ctx, *customer,"register")
    if err != nil {
		return "", err
    }
	res, err := u.authRepo.Create(ctx, customer)

	
	customerResponse, ok := res.(*model.Customer)
	if !ok {
		return nil, errors.New("unexpected response from repository")
	}

	
	response := &response.CustomerResponse{
		ID:        int(customerResponse.ID),
		Name:      customerResponse.Name,
		Email:     customerResponse.Email,
		Phone:      customerResponse.Phone,
		Address:   customer.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

	}

	return response, err
}
type ErrorResponse struct {
    Error string `json:"error"`
	Code  int    `json:"code"`
}

func (u *AuthService) LoginCustomer(ctx context.Context, customer *model.Customer) (*response.LoginResponse,  ErrorResponse) {
	errResp := ErrorResponse{}
	

	// Get customer by email
	var cus model.Customer
	err := u.authRepo.GetByEmail(ctx, customer.Email, &cus)
	if err != nil {
		errResp.Error = "Email not found"
		errResp.Code = http.StatusBadRequest
		return nil, errResp
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(cus.Password), []byte(customer.Password))
	if err != nil {
		errResp.Error = "Incorrect password"
		errResp.Code = http.StatusBadRequest
		return nil, errResp
	}

	// Generate token
	token, err := security.GenerateToken(cus.ID)
	if err != nil {
		errResp.Error = "Failed to generate token"
		errResp.Code = http.StatusInternalServerError
		return nil, errResp
	}

	// Create new log
	err = u.CreateNewLog(err, ctx, cus, "login")
	if err != nil {
		errResp.Error = "Failed to create new log"
		errResp.Code = http.StatusInternalServerError
		return nil, errResp
	}

	// Create new token
	_, err = u.TokenService.CreateNewToken(ctx, cus.ID, token)
	if err != nil {
		errResp.Error = "Failed to create new token"
		errResp.Code = http.StatusInternalServerError
		return nil, errResp
	}

	// Return success response
	response := &response.LoginResponse{
		ID:       int(cus.ID),
		Email:    cus.Email,
		Password: cus.Password,
		Token:    token,
	}

  return response, errResp

}


func (u *AuthService) LogoutCustomer(ctx context.Context,token string,id uint) (interface{},error) {
	
	res, err := u.TokenService.DisableToken(ctx, token,id )
	if err != nil {
		// error handling
		return "",nil
	}
	_, err = u.UserLogService.CreateNewLog(ctx, id, "logout")
	return res,nil
	
}



func (u *AuthService) CreateNewLog(err error, ctx context.Context, cus model.Customer,action string) error {
	_, err = u.UserLogService.CreateNewLog(ctx, cus.ID, action)
	return err
}


