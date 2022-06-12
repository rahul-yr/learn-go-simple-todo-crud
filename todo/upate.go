package todo

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// this function is responsible for updating a todo by id
func updateTodo(c *gin.Context) {
	var inputTodo Todo
	// bind the json body to the todo struct
	c.BindJSON(&inputTodo)
	// get the id from the url
	id := c.Param("id")
	// convert the id to an integer
	todoId, err := strconv.Atoi(id)
	inputTodo.ID = todoId
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	// check if input todo is valid
	if inputTodo.Title == "" {
		c.JSON(400, gin.H{"error": "title is required"})
		return
	}
	// loop through the todos
	for i, todo := range todos {
		// if the id matches the id in the todo
		if todo.ID == todoId {
			todos[i] = inputTodo
			c.JSON(200, inputTodo)
			return
		}
	}
	c.JSON(404, gin.H{"error": "todo not found"})
}
