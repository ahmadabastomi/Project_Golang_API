package controllers

import (
	"fmt"
	"golangApi/models"
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
)

var ORM orm.Ormer

func init() {
	models.ConnectToDb()
	ORM = models.GetOrmObject()
}

var CreateTodo = func(c *gin.Context) {
	var newTodo models.Todos
	c.BindJSON(&newTodo)

	_, err := ORM.Insert(&newTodo)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Created New Todo Success"})
	} else {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to create new Todo"})
	}
}

var ReadTodos = func(c *gin.Context) {
	var todos []models.Todos
	fmt.Println(ORM)
	_, err := ORM.QueryTable("todos").All(&todos)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "todos": &todos})
	} else {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to read the todos"})
	}
}

var ReadTodo = func(c *gin.Context) {
	var todo models.Todos
	id := c.Param("id")
	fmt.Println(ORM)
	err := ORM.QueryTable("todos").Filter("id", id).One(&todo)
	if err == orm.ErrNoRows {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to read the todos"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &todo})
	}
}

var UpdateUser = func(c *gin.Context) {
	var updateTodo models.Todos
	id := c.Param("id")
	c.BindJSON(&updateTodo)
	_, err := ORM.QueryTable("todos").Filter("id", id).Update(
		orm.Params{"todo": updateTodo.Todo})
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	} else {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to update the users"})
	}
}

var DeleteTodo = func(c *gin.Context) {
	var delTodo models.Todos
	id := c.Param("id")
	c.BindJSON(&delTodo)
	_, err := ORM.QueryTable("todos").Filter("id", id).Delete()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	} else {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to delete the todo"})
	}
}
