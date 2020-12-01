package api

import (
	api "Go-000/Week02/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.POST("api/vi/user/getUser", api.GetUser)
	return r
}
