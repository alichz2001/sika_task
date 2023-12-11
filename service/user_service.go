package service

import (
	"errors"
	"gorm.io/gorm"
	"sika/model"
)

type UserService interface {
	GetByID(ID uint) (*model.User, error)
	GetPage(page int, count int) ([]*model.User, error)
}

type UserServiceImpl struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserServiceImpl {
	return &UserServiceImpl{db}
}

func (userSrv *UserServiceImpl) GetByID(ID uint) (*model.User, error) {
	var user model.User
	userSrv.db.Where("id = ?", ID).First(&user)
	if user.ID == 0 {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (userSrv *UserServiceImpl) GetPage(page int, count int) ([]*model.User, error) {
	var users []*model.User
	userSrv.db.Limit(count).Offset((page - 1) * count).Find(&users)

	//x := userSrv.db.Session(&gorm.Session{DryRun: true}).Find(&users).Limit(count).Offset((page - 1) * count)
	//log.Print(x.Statement.SQL.String())
	if users == nil {
		return nil, errors.New("page not found")
	}
	return users, nil
}
