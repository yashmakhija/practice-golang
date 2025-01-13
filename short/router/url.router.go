package router

import (
	"github.com/yashmakhija/practice-golang/tree/main/short/controller"

	"github.com/gin-gonic/gin"
)

func urlRouter(route *gin.RouterGroup){
	
	router := route.Group("/url")

	router.GET("/create", controller.urlController)


}