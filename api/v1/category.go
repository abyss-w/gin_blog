package v1

import (
	"github.com/abyss-w/gin_blog/model"
	"github.com/abyss-w/gin_blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加分类
func AddCategory(c *gin.Context) {
	var category model.Category
	_ = c.ShouldBindJSON(&category)
	code := model.ContainsCategory(category.Name)
	if code == errmsg.SUCCEED {
		model.AddCategory(&category)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": category,
		"message": errmsg.GetMsg(code),
	})
}

//查询单个分类

//查询分类列表
func GetCategories(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, total := model.GetCategories(pageSize, pageNum)
	code := errmsg.SUCCEED
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"total":total,
		"message": errmsg.GetMsg(code),
	})
}

//编辑分类
func EditCategory(c *gin.Context) {
	var category model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&category)
	code := model.ContainsCategory(category.Name)
	if code == errmsg.SUCCEED {
		code = model.EditCategory(id, &category)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetMsg(code),
	})
}

//删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetMsg(code),
	})
}

//TODO 查询分类下的所有文章
