package app

import (
	"github.com/labstack/echo/v4"
)

type (
	router struct {
		*echo.Echo
	}

	// Controller for a Model
	ResourceController interface {
		// Returns a list of all of these resources
		List(echo.Context) error
		// Returns a specific resource (by id).
		Show(echo.Context) error
		// Creates a new instance of the resource.
		Create(echo.Context) error
		// Updates a resource.
		Update(echo.Context) error
		// Deletes a resource.
		Delete(echo.Context) error
	}
)

var Router *router = &router{echo.New()}

func (r *router) Resource(path string, c ResourceController, middleware ...echo.MiddlewareFunc) {
	g := r.Group(path, middleware...)
	g.GET("/", c.List)
	g.GET("/:id", c.Show)
	g.POST("/", c.Create)
	g.PATCH("/:id", c.Update)
	g.DELETE("/:id", c.Delete)
}
