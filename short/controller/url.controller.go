package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func urlController(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "url controller",
	})
}