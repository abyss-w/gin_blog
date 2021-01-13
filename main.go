package main

import (
	"github.com/abyss-w/gin_blog/model"
	"github.com/abyss-w/gin_blog/routes"
)

func main() {
	model.InitDB()
	routes.InitRouter()
}
