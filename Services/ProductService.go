package Services

//customer id
//product id
import (
	"context"
	"fmt"
	model "go-test/Model"
	repository "go-test/Repository"
)

type IProductService interface {
	GetProductById(ctx context.Context,id uint) (*model.Product, error)
}

type ProductService struct {
	productRepo repository.IRepository
}

func NewProductService(productRepo repository.IRepository) * ProductService {
	return &ProductService{productRepo: productRepo}
}

func (u *ProductService)GetProductById(ctx context.Context,id uint) (*model.Product, error){
	var p model.Product
	fmt.Println(id)
    err := u.productRepo.GetByID(ctx, id, &p)
    if err != nil {
        return &p, err
    }
	
	return &p,err
}
