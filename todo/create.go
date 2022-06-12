package todo

import "github.com/gin-gonic/gin"

// this variable is used to generate unique ids for todos
var idCounter int = 1

// This is a function to create a new todo
func createTodo(c *gin.Context) {
	var todo Todo
	// bind the json body to the todo struct
	c.BindJSON(&todo)
	// check if the title is empty
	if todo.Title == "" {
		c.JSON(400, gin.H{"error": "title is required"})
		return
	} else {
		// set the id
		todo.ID = idCounter
		todos = append(todos, todo)
		// increment the idCounter
		idCounter++
		c.JSON(200, todo)
	}
}
