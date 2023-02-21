package Repository

import (
	"context"

	"gorm.io/gorm"
)

type IRepository interface {
	GetAll(ctx context.Context, results interface{}) error
	GetByID(ctx context.Context, id uint, result interface{}) error
	Create(ctx context.Context, data interface{})(interface{}, error)
	CreateData(ctx context.Context, data interface{}) error
	Update(ctx context.Context, id uint, data interface{}) (interface{}, error)
	UpdateToken(ctx context.Context, token string, data interface{}) (interface{}, error)
	Delete(ctx context.Context, id uint) error
	GetByEmail(ctx context.Context, email string, result interface{}) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return &repository{db}
}

func (r *repository) GetAll(ctx context.Context, results interface{}) error {
	return r.db.Find(results).Error
}

func (r *repository) GetByID(ctx context.Context, id uint, result interface{}) error {
	return r.db.First(result, id).Error
}
func (r *repository) GetByEmail(ctx context.Context, email string, result interface{}) error {
    return r.db.Where("email = ?", email).First(result).Error
}
func (r *repository) Create(ctx context.Context, data interface{}) (interface{}, error) {
    err := r.db.Create(data).Error
    if err != nil {
        return nil, err
    }

    // return the inserted data
    return data, nil 
}
func (r *repository) CreateData(ctx context.Context, data interface{})  error {
   return r.db.Create(data).Error
  
}
func (r *repository) Update(ctx context.Context, id uint, data interface{}) (interface{}, error)  {
	err:= r.db.Model(data).Where("id = ?", id).Updates(data).Error
	if err != nil {
        return nil, err
    }

  
    return data, nil 
}
func (r *repository) UpdateToken(ctx context.Context, token string, data interface{}) (interface{}, error)  {
	err:= r.db.Model(data).Where("token = ?", token).Updates(data).Error
	if err != nil {
        return nil, err
    }

  
    return data, nil 
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	return r.db.Delete(nil, id).Error
}