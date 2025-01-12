package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var todos []Todo

type Todo struct{
	Id int `json:"id"`
	Title string `json:"title"`
	IsComplete bool `json:"isComplete"`
}


func main(){
	
	todos = []Todo{}

	router := gin.Default()
	
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/todo", func(ctx *gin.Context) {
		var todo Todo

		if err:= ctx.BindJSON(&todo); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} 

		todo.Id = len(todos) + 1

		todos = append(todos, todo)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Todo added",
		})
	})

	router.GET("/todo", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, todos)
	})

	router.PUT("/todo/:id", func(ctx *gin.Context) {

		idParam := 	ctx.Param("id")
		id, err := strconv.Atoi(idParam)

		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid id format",
			})
			return
		}

		var updatedTodo Todo

		if err := ctx.BindJSON(&updatedTodo); err != nil{
			ctx.JSON(http.StatusBadRequest,gin.H{
				"error": "Invalid JSON format",
			})
			return
		}

		for _, todo := range todos{
			if todo.Id == id{
				todos[id].IsComplete = true

				ctx.JSON(http.StatusOK, gin.H{
					"message": "Todo Marked as completed",
					"todos": todos[id],
				})
				return
			}
			
		}

		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Todo not found",
		})
		
	})

	router.DELETE("/todo/:id",func(ctx *gin.Context) {
		var todoId = ctx.Param("id")

		id,err := strconv.Atoi(todoId)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid id format",
			})
			return
		}

		for i, todo := range todos{
			if todo.Id == id {
				todos = append(todos[:i], todos[i+1:]... )
				
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Todo deleted",
				})
				return
			}
		}

		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Todo not found",
		})
	})

	router.Run(":9012")
}