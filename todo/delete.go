package todo

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// this function is responsible for deleting a todo
func deleteTodo(c *gin.Context) {
	// get the id from the url
	id := c.Param("id")
	// convert the id to an integer
	todoId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	// loop through the todos
	for i, todo := range todos {
		// if the id matches the id in the todo
		if todo.ID == todoId {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(200, gin.H{"message": "todo deleted"})
			return
		}
	}
	c.JSON(404, gin.H{"error": "todo not found"})
}
