package http

import (
	"net/http"
	"strconv"

	"github.com/andyantrim/apiutil/messages"
	"github.com/andyantrim/coreapi/service"
	"github.com/labstack/echo/v4"
)

// ResourceController has all the http controllers
// for the /resource route
type ResourceController struct {
	srv service.ResourceInterface
}

// NewResourceController builds a new ResourceController
func NewResourceController(srv service.ResourceInterface) *ResourceController {
	return &ResourceController{
		srv: srv,
	}
}

// Routes contains all endpoints
func (r *ResourceController) Routes(g *echo.Group) {
	g.GET("/resource", r.List)
	g.GET("/resource/:id", r.Get)
}

func (r *ResourceController) List(c echo.Context) error {
	resp, err := r.srv.List()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (r *ResourceController) Get(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return messages.NewRequiredFieldError("id must be a number")
	}
	resp, err := r.srv.Get(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
