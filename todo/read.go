package todo

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// this function is responsible for fetching a todo by id
func readTodo(c *gin.Context) {
	// get the id from the url
	id := c.Param("id")
	// convert the id to an integer
	todoId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	// loop through the todos
	for _, todo := range todos {
		// if the id matches the id in the todo
		if todo.ID == todoId {
			c.JSON(200, todo)
			return
		}
	}
	c.JSON(404, gin.H{"error": "todo not found"})
}

// this function is responsible for fetching all todos
func listTodos(c *gin.Context) {
	c.JSON(200, todos)
}
