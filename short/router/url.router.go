package router

import (
	"short/controller"

	"github.com/gin-gonic/gin"
)

func urlRouter(route *gin.RouterGroup){
	
	router := route.Group("/url")

	router.GET("/create", controller.urlController)


}