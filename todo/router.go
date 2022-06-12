package todo

import "github.com/gin-gonic/gin"

// Todo is a struct that represents a todo
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// This variable is responsible for storing all todos
var todos []Todo = []Todo{}

// MountTodoRouter is responsible for mounting the todo router
func MountTodoRouter(c *gin.RouterGroup) {
	c.POST("/create", createTodo)
	c.GET("/read/:id", readTodo)
	c.POST("/update/:id", updateTodo)
	c.DELETE("/delete/:id", deleteTodo)
	c.GET("/list", listTodos)
}
