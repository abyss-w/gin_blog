package v1

import (
	"github.com/abyss-w/gin_blog/model"
	"github.com/abyss-w/gin_blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")

	fileSize := fileHeader.Size

	url, code := model.UploadFile(file, fileSize)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetMsg(code),
		"url": url,
	})
}
