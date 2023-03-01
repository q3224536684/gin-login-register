package services

import (
	"context"
	"frank/gin-login-register/models"
	"gorm.io/gorm"
	"log"
)

type UserService interface {
	InitTable()
	Register(info *models.UserInfo) (*models.UserInfo, error)
	Login(username string, password string) (*models.UserInfo, error)
	FindByUsername(username string) (*models.UserInfo, error)
	UsernameHasRegister(username string) bool
	UserList() ([]*models.UserInfo, error)
}

type UserServiceImpl struct {
	db  *gorm.DB
	ctx context.Context
}

func NewUserServiceImpl(db *gorm.DB, ctx context.Context) *UserServiceImpl {
	return &UserServiceImpl{db: db, ctx: ctx}
}

func (u UserServiceImpl) InitTable() {
	err := u.db.AutoMigrate(&models.UserInfo{})
	if err != nil {
		log.Fatal("初始化表失败, error: ", err)
	}
}

func (u UserServiceImpl) Register(info *models.UserInfo) (*models.UserInfo, error) {
	result := u.db.Create(&info)
	if result.Error != nil {
		return nil, result.Error
	}
	return info, nil
}

func (u UserServiceImpl) Login(username string, password string) (*models.UserInfo, error) {
	var user *models.UserInfo
	result := u.db.Where("username=?", username).Where("password=?", password).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u UserServiceImpl) UserList() ([]*models.UserInfo, error) {
	var mds []*models.UserInfo
	result := u.db.Find(&mds)
	if result.Error != nil {
		return nil, result.Error
	}
	return mds, nil
}

func (u UserServiceImpl) FindByUsername(username string) (*models.UserInfo, error) {
	var user *models.UserInfo
	result := u.db.Where("username=?", username).First(&user)
	log.Printf("result.RowsAffected: %+v", result)
	if result.RowsAffected < 1 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u UserServiceImpl) UsernameHasRegister(username string) bool {
	result := u.db.Where("username=?", username).First(&models.UserInfo{})
	log.Printf("result.RowsAffected: %+v", result.RowsAffected)
	return result.RowsAffected > 0
}
