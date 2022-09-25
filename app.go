package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saisesai/ancient-database-backend/config"
	"github.com/saisesai/ancient-database-backend/controller"
)

func main() {
	fmt.Println("http listen at:", config.C.HttpListenAddress)

	app := gin.Default()

	app.POST("/api/char/add", controller.CharAddHandler)
	app.POST("/api/char/del", controller.CharDeleteHandler)
	app.POST("/api/char/get", controller.CharGetHandler)
	app.POST("/api/char/mod", controller.CharModifyHandler)

	app.POST("/api/page/add", controller.PageAddHandler)
	app.POST("/api/page/del", controller.PageDeleteHandler)
	app.POST("/api/page/get", controller.PageGetHandler)
	app.POST("/api/page/mod", controller.PageModifyHandler)

	app.POST("/api/artwork/add", controller.ArtworkAddHandler)
	app.POST("/api/artwork/del", controller.ArtworkDeleteHandler)
	app.POST("/api/artwork/get", controller.ArtworkGetHandler)
	app.POST("/api/artwork/mod", controller.ArtworkModifyHandler)

	panic(app.Run(config.C.HttpListenAddress))
}
