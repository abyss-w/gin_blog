package v1

import (
	"github.com/abyss-w/gin_blog/middleware"
	"github.com/abyss-w/gin_blog/model"
	"github.com/abyss-w/gin_blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)

	var (
		token string
		code int
	)

	code = model.CheckLogin(data.Name, data.Password)
	if code == errmsg.SUCCEED {
		token, code = middleware.SetToken(data.Name)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetMsg(code),
		"token": token,
	})
}
