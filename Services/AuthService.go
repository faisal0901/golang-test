package Services

import (
	"context"
	"errors"

	model "go-test/Model"
	repository "go-test/Repository"
	security "go-test/Security"
	"time"

	bcrypt "golang.org/x/crypto/bcrypt"
)

type CustomerResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type LoginResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"phone"`
	Token   string    `json:"address"`
}
type IAuthService interface {
	CreateCustomer(ctx context.Context, customer *model.Customer) (interface{}, error)
	LoginCustomer(ctx context.Context, customer *model.Customer) (interface{}, error)
	LogoutCustomer(ctx context.Context,token string)error
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

	
	response := map[string]interface{}{
		"id":         customerResponse.ID,
		"name":       customerResponse.Name,
		"email":      customerResponse.Email,
		"phone":      customerResponse.Phone,
		"address":    customerResponse.Address,
		"created_at": customerResponse.CreatedAt,
		"updated_at": customerResponse.UpdatedAt,
		"message":    "customer created successfully",
	}

	return response, err
}
func (u *AuthService) LoginCustomer(ctx context.Context, customer *model.Customer) (interface{}, error) {
	
	var cus model.Customer
    err := u.authRepo.GetByEmail(ctx, customer.Email, &cus)
    if err != nil {
        return "", err
    }
	err = bcrypt.CompareHashAndPassword([]byte(cus.Password), []byte(customer.Password))
    if err != nil {
        return "", err
    }

	token,err := security.GenerateToken(cus.ID)
	
	
	err = u.CreateNewLog(err, ctx, cus,"login")
    if err != nil {
		return "", err
    }
	_,err = u.TokenService.CreateNewToken(ctx, cus.ID,token)
    if err != nil {
		return "", err
    }
		
	
	response := map[string]interface{}{
		"id":         cus.ID,
		"name":       cus.Name,
		"email":      cus.Email,
		"password":   cus.Password,
		"token":    token,
		"message":    "login succesfuly",
	}
	return response,nil
}
func (u *AuthService) LogoutCustomer(ctx context.Context,token string) (error) {
	u.TokenService.DisableToken(ctx,token)
	return nil
}



func (u *AuthService) CreateNewLog(err error, ctx context.Context, cus model.Customer,action string) error {
	_, err = u.UserLogService.CreateNewLog(ctx, cus.ID, action)
	return err
}


