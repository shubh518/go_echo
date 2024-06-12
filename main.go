package main

import (
	"github.com/labstack/echo"
	"main.go/cmd/api/handlers"
)

func main() {

	// Echo is a web framework for the Go programming language.
	// It is a high performance, extensible, minimalist web framework for building web applications and APIs.
	// It features a robust set of features including routing, middleware, templating, and more.

	// To create a new instance of the echo package, you can use the New() function:
	e := echo.New()

	// GET registers a new GET route for a path with matching handler in the router with optional route-level middleware.

	e.GET("/movies", handlers.MoviesHandler)
	e.GET("/posts", handlers.PostIndexHandler)
	e.GET("/posts/:id", handlers.PostSingleHandler)

	// This code will start a server on port 1234 and log any errors that occur.
	//If an error occurs that is fatal, the server will be stopped and the error will be logged.

	e.Logger.Fatal(e.Start(":1234"))

}
