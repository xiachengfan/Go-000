package v1

import (
	"Go-000/Week02"
	"Go-000/Week02/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
)

//返回太丑了，请原谅。
func GetUser(c *gin.Context) {
	param := service.GetAuthRequest{}
	err := c.ShouldBind(&param)
	if err != nil {
		c.JSON(200, service.Resp{
			CODE: 400,
			ERR:  err,
			MSG:  nil,
		})
		return
	}
	srv := service.New(c)
	user, err := srv.CheckAuth(&param)
	r := service.Resp{
		CODE: 200,
		ERR:  nil,
		MSG:  user,
	}
	if err != nil && errors.Is(err, Week02.ErrUserNotFound) {
		r.CODE = 404
		r.ERR = err
	}
	if err != nil {
		r.CODE = 500
		r.ERR = err
	}
	c.JSON(200, r)
	return
}
