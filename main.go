package main

import (
	"os"
	"rahul-yr/learn-go-simple-todo-crud/todo"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

// main web server
func main() {
	// set a gin mode to release
	gin.SetMode(gin.ReleaseMode)
	// create a gin router
	r := gin.Default()
	// enable cors for all
	r.Use(cors.Default())
	// create a todo service group
	router := r.Group("/todo")
	// add a todo service
	todo.MountTodoRouter(router)
	// start the web server
	// listen and serve on
	// address:port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
