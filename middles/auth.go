package middles

import (
	"frank/gin-login-register/config"
	"frank/gin-login-register/controllers"
	"frank/gin-login-register/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitAuthMiddlewares(controllers controllers.UserControllers) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		IdentityKey:      "id",
		Realm:            "test-jwt",
		SigningAlgorithm: "HS256",
		Key:              []byte(config.GetConfig().Jwt.Key),
		Timeout:          time.Hour * time.Duration(config.GetConfig().Jwt.AccessAge),
		MaxRefresh:       time.Hour * time.Duration(config.GetConfig().Jwt.RefreshAge),
		TokenLookup:      "header: Authorization",
		TokenHeadName:    "Bearer",
		TimeFunc:         time.Now,
		Authenticator:    controllers.Login,
		Authorizator:     authorizedFunc,
		PayloadFunc:      payloadFunc,
		LoginResponse:    loginResponse,
		LogoutResponse:   logoutResponse,
		Unauthorized:     unauthorizedFunc,
		IdentityHandler:  identityHandler,
	})
}

func authorizedFunc(data interface{}, c *gin.Context) bool {
	claims := jwt.ExtractClaims(c)
	return data == claims["id"]
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return claims["id"]
}
func payloadFunc(data interface{}) jwt.MapClaims {
	return jwt.MapClaims{
		"id":       data.(*models.UserInfo).ID,
		"username": data.(*models.UserInfo).Username,
	}
}

func unauthorizedFunc(c *gin.Context, code int, message string) {
	c.JSON(http.StatusUnauthorized, models.Resp{
		RequestId: requestid.Get(c),
		Code:      code,
		Msg:       message,
	})
}
func loginResponse(c *gin.Context, code int, token string, expires time.Time) {
	c.JSON(http.StatusOK, models.Resp{
		RequestId: requestid.Get(c),
		Code:      code,
		Data: models.Token{
			Token:   "Bearer " + token,
			Expired: expires,
		},
		Msg: "SUCCESS",
	})
}

func logoutResponse(c *gin.Context, code int) {
	c.JSON(http.StatusOK, models.Resp{
		RequestId: requestid.Get(c),
		Code:      code,
		Data:      nil,
		Msg:       "SUCCESS",
	})
}
