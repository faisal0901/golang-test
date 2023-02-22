package Services

//customer id
//product id
import (
	"context"
	model "go-test/Model"
	repository "go-test/Repository"
	response "go-test/Response"
)

type IProductService interface {
	GetProductById(ctx context.Context,id uint) (*model.Product, error)
	CreateNewProduct(ctx context.Context,product *model.Product) (interface{}, error)
}

type ProductService struct {
	productRepo repository.IRepository
}

func NewProductService(productRepo repository.IRepository) * ProductService {
	return &ProductService{productRepo: productRepo}
}

func (u *ProductService)GetProductById(ctx context.Context,id uint) (*model.Product, error){
	var p model.Product
	
    err := u.productRepo.GetByID(ctx, id, &p)
    if err != nil {
        return &p, err
    }
	
	return &p,err
}
func (u *ProductService)CreateNewProduct(ctx context.Context,product *model.Product) (interface{}, error){


    res,err := u.productRepo.Create(ctx, product)
    if err != nil {
        return res, err
    }
	respone := &response.ProductResponse{
        ID:          product.ID,
        Name:        product.Name,
        Description: product.Description,
        MerchantID:  product.MerchantID,
        Price:       uint(product.Price),
        CreatedAt:   product.CreatedAt,
        UpdatedAt:   product.UpdatedAt,
    }
	
	return respone,err
}

