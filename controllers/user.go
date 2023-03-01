package controllers

import (
	"context"
	"frank/gin-login-register/models"
	"frank/gin-login-register/services"
	"github.com/gin-gonic/gin"
)

type UserControllers interface {
	Register(c *gin.Context)
	Login(c *gin.Context) (interface{}, error)
	UserList(c *gin.Context)
}

type UserControllersImpl struct {
	ctx         context.Context
	userService services.UserService
}

func NewUserControllersImpl(ctx context.Context, userService services.UserService) *UserControllersImpl {
	return &UserControllersImpl{ctx: ctx, userService: userService}
}

func (u UserControllersImpl) Register(c *gin.Context) {
	var registerInput *models.UserReq
	err := c.ShouldBindJSON(&registerInput)
	if err != nil {
		models.Fail(500, err.Error(), c)
		return
	}
	exist := u.userService.UsernameHasRegister(registerInput.Username)
	if exist {
		models.Fail(500, "用户名已注册", c)
		return
	}
	info, err := u.userService.Register(&models.UserInfo{
		Username: registerInput.Username,
		Password: registerInput.Password,
		CreateAt: models.GetNowTime(),
		UpdateAt: models.GetNowTime(),
	})
	if err != nil {
		models.Fail(500, err.Error(), c)
		return
	}
	models.Success(info, c)
}

func (u UserControllersImpl) Login(c *gin.Context) (interface{}, error) {
	var loginInput *models.UserReq
	err := c.ShouldBindJSON(&loginInput)
	if err != nil {
		return nil, err
	}
	res, err := u.userService.Login(loginInput.Username, loginInput.Password)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u UserControllersImpl) UserList(c *gin.Context) {
	res, err := u.userService.UserList()
	if err != nil {
		models.Result(500, err.Error(), nil, c)
		return
	}
	models.Result(200, "SUCCESS", res, c)
}
