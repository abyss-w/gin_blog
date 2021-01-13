package v1

import (
	"github.com/abyss-w/gin_blog/model"
	"github.com/abyss-w/gin_blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	model.AddArticle(&data)
	code := errmsg.SUCCEED
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetMsg(code),
		"data": data,
	})
}

//TODO 查询单个文章
func GetArticleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArticleInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetMsg(code),
		"data": data,
	})
}

//TODO 查询分类下的所有文章
func GetCateArticle(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("cid"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	articles, code, total := model.GetCateArticle(cid, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetMsg(code),
		"data": articles,
		"total": total,
	})
}

//查询文章列表
func GetArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	articles, code, total := model.GetArticles(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetMsg(code),
		"data": articles,
		"total": total,
	})
}

//编辑文章
func EditArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var article model.Article
	_ = c.ShouldBindJSON(&article)
	code := model.EditArticle(id, &article)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetMsg(code),
	})
}

//删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetMsg(code),
	})
}
