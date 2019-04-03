package main

import (
	"golangApi/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/createTodo", controllers.CreateTodo)
	router.GET("/readTodos", controllers.ReadTodos)
	router.GET("/readTodo/:id", controllers.ReadTodo)
	router.PUT("/updateTodo/:id", controllers.UpdateUser)
	router.DELETE("/deleteTodo/:id", controllers.DeleteTodo)
	router.Run(":3001")
}
