package Services

import (
	"context"
	model "go-test/Model"
	repository "go-test/Repository"
)


type IMerchantService interface {
	GetAllMerchant(ctx context.Context) ([]*model.Merchant, error)
	CreateMerchant(ctx context.Context ,merchant *model.Merchant) (interface{}, error)
}
type MerchantService struct {
	merchantRepo repository.IRepository
}
func NewMerchantService(merchantRepo repository.IRepository) * MerchantService {
	return &MerchantService{merchantRepo: merchantRepo}
}

func (u *MerchantService) GetAllMerchant(ctx context.Context) ([]*model.Merchant, error){
	var results []*model.Merchant
	
	err := u.merchantRepo.GetAllProduct(ctx, &results)
    if err != nil {
        return nil, err
    }
	return results,nil
}

func (u *MerchantService)	CreateMerchant(ctx context.Context,merchant *model.Merchant) (interface{}, error){

	results,err := u.merchantRepo.Create(ctx,merchant)
    if err != nil {
        return nil, err
    }
	return results,nil
}

