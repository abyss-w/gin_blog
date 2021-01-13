package routes

import (
	v1 "github.com/abyss-w/gin_blog/api/v1"
	"github.com/abyss-w/gin_blog/middleware"
	"github.com/abyss-w/gin_blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	auth := r.Group("/api/v1")
	auth.Use(middleware.JwtToken())
	{
		//user模块路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		//category模块路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		auth.PUT("category/:id", v1.EditCategory)

		//article模块路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		auth.PUT("article/:id", v1.EditArticle)

		// upload file
		auth.POST("upload", v1.Upload)
	}

	router := r.Group("/api/v1")
	{
		//user模块路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)

		//category模块路由接口
		router.GET("categories", v1.GetCategories)

		//article模块路由接口
		router.GET("articles", v1.GetArticles)
		router.GET("article/cate/:cid", v1.GetCateArticle)
		router.GET("article/info/:id", v1.GetArticleInfo)

		//Login
		router.POST("login", v1.Login)
	}

	_ = r.Run(utils.HttpPort)
}
