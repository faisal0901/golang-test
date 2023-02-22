package Services

import (
	"context"
	model "go-test/Model"
	repository "go-test/Repository"
)
type ITokenService interface {
	CreateNewToken(ctx context.Context,id uint,token string) (interface{},error)
	DisableToken(ctx context.Context,token string,id uint)  (interface{},error)
}
type TokenService struct {
	tokenRepo repository.IRepository
	
}
func NewTokenService(tokenRepo repository.IRepository) * TokenService {
	return &TokenService{tokenRepo: tokenRepo}
}
func (u *TokenService) CreateNewToken(ctx context.Context,id uint,token string)  (interface{},error){
	var tokenService = model.Token{
	Token: token,
	IsValid: 1,
	CustomerID: id,
	}
	res,err := u.tokenRepo.Create(ctx,&tokenService)
	return res,err
}
func (u *TokenService) DisableToken(ctx context.Context,token string,id uint)  (interface{},error){
	
	res, err := u.tokenRepo.UpdateToken(ctx, token, &model.Token{})
	if err != nil {
		return res,err
	}
	return res, nil
}
