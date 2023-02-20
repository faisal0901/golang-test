package Services

import (
	"context"
	model "go-test/Model"
	repository "go-test/Repository"
)
type ILogService interface {
	CreateNewLog(ctx context.Context,id uint,action string) (interface{},error)
}

type LogService struct {
	LogRepo repository.IRepository
}
func NewLogService(LogRepo repository.IRepository) *LogService {
	return &LogService{LogRepo: LogRepo}
}

func (u *LogService) CreateNewLog(ctx context.Context,id uint,action string)  (interface{},error){
	var userLog = model.UserLog{
		UserID: id,
		Action: action,
	}
	res,err := u.LogRepo.Create(ctx,&userLog)
	return res,err
}