package v1

import (
	"github.com/abyss-w/gin_blog/model"
	"github.com/abyss-w/gin_blog/utils/errmsg"
	"github.com/abyss-w/gin_blog/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//查询用户是否存在
func UserExist(c *gin.Context) {

}

//添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)

	msg, code := validator.Validate(data)
	if code != errmsg.SUCCEED {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"message": msg,
		})
		return
	}

	code = model.ContainsUser(data.Name)
	if code == errmsg.SUCCEED {
		model.CreateUser(&data)
	} else if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetMsg(code),
	})
}

//查询单个用户

//查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, total := model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCEED
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetMsg(code),
		"total": total,
	})
}

//编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.ContainsUser(data.Name)
	if code == errmsg.SUCCEED {
		model.EditUser(id, &data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetMsg(code),
	})
}

//删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetMsg(code),
	})
}
