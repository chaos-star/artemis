package UserService

import (
	"Artemis/Common/Code"
	"Artemis/Global"
	"Artemis/Model"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func New() *UserService {
	return &UserService{
		Global.DB.Instance("db-alias"),
	}
}

func (u *UserService) Example(params ExampleParam) (ex *Code.Exception, data Model.User) {
	ex = &Code.Exception{}
	err := u.db.Where("id = ?", params.Id).First(&data).Error
	if err != nil {
		ex.SetCode(Code.DbSelectError, err)
	}
	return
}
