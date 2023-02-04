package main

import (
	"fmt"
	"net/http"

	"github.com/andyantrim/apiutil"
	"github.com/andyantrim/coreapi/transport"
	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()
	router.RouteNotFound("/*", func(c echo.Context) error {
		return c.String(http.StatusNotFound, "Not found base")
	})
	apiGroup := router.Group("/api/v1")

	transport.InitV1(apiGroup)

	for _, route := range router.Routes() {
		fmt.Printf("%s %s \n", route.Method, route.Path)
	}

	// Custom error handler
	router.HTTPErrorHandler = apiutil.CustomerHTTPErrorHandler

	router.Logger.Fatal(router.Start(":1234"))
}
