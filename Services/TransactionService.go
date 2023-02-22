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
	GetAllTransaction(ctx context.Context) ([]CustomerResponse, error)
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
func (u *TransactionService) GetAllTransaction(ctx context.Context) ([]CustomerResponse, error) {
    var results []*model.Customer
	
	err := u.transRepo.GetAllJoin(ctx, &results)
    if err != nil {
        return nil, err
    }
	var response []CustomerResponse
    for _, customer := range results {
        var transactions []TransactionResponse
        for _, trans := range customer.Transaction {
            transactions = append(transactions, TransactionResponse{
                ID:        trans.ID,
                Qty:       int(trans.Qty),
                Price:     float64(trans.Price),
                ProductID: trans.ProductID,
            })
        }
		
        response = append(response, CustomerResponse{
            ID:          customer.ID,
            Name:        customer.Name,
            Email:       customer.Email,
            Phone:       customer.Phone,
            Address:     customer.Address,
            Transaction: transactions,
        })
    }
	
    return response, nil
}
type CustomerResponse struct {
    ID         uint            `json:"id"`
    Name       string          `json:"name"`
    Email      string          `json:"email"`
    Phone      string          `json:"phone"`
    Address    string          `json:"address"`
    Transaction []TransactionResponse `json:"transactions"`
}

type TransactionResponse struct {
    ID        uint    `json:"id"`
    Qty       int     `json:"qty"`
    Price     float64 `json:"price"`
    ProductID uint    `json:"productId"`
}