package Services

//customer id
//product id
import (
	"context"
	model "go-test/Model"
	repository "go-test/Repository"
)

type ITransactionService interface {
	CreateNewTransaction(ctx context.Context,transaction *model.Transaction) (interface{}, error)
}

type TransactionService struct {
	transRepo repository.IRepository
	ProductService IProductService
}

func NewTransactionService(transRepo repository.IRepository,productService IProductService) *TransactionService {
	return &TransactionService{
		transRepo: transRepo,
		ProductService: productService,
	}
}

//get product
//get customer id
//insert transaction
func (u *TransactionService) CreateNewTransaction(ctx context.Context,transaction *model.Transaction)   (interface{},error){
	
	p,err:=u.ProductService.GetProductById(ctx,transaction.ProductID)
	if err != nil {
        return p, err
    }
	var transactioninsert = model.Transaction{
		Qty:        transaction.Qty,
		Price:      transaction.Qty * float32(p.Price),
		CustomerID: transaction.CustomerID,
		ProductID:  transaction.ProductID,
	}
	res,err := u.transRepo.Create(ctx,&transactioninsert)
	
	return res,err
}