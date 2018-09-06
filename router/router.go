package router

import (
	"github.com/gin-gonic/gin"
	"github.com/konojunya/musi/controller"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	// static
	r.Static("/js", "./public/js")
	r.Static("/css", "./public/css")
	r.Static("/images", "./public/images")

	r.LoadHTMLGlob("views/*")

	r.GET("/", controller.Index)
	r.GET("/login", controller.Login)
	r.GET("/oauth", controller.OAuth)

	return r
}
