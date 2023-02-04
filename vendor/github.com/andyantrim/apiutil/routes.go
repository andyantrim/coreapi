package apiutil // apiutil has untility functions for APIs

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// RouteInterface is an interface that defines
// Routes(g *echo.Group)
// And can be used to pass slices of structs to Register
// all routes tied to those structs
type RouteInterface interface {
	Routes(g *echo.Group)
}

// RegisterRoutes takes a slice of RouteInterface
// And calls Routes(...) on them, Registering all
// the provided routes to the given group
func RegisterRoutes(g *echo.Group, rr ...RouteInterface) {
	for _, r := range rr {
		fmt.Println("Registering")
		r.Routes(g)
	}
}
