package apiutil

import (
	"github.com/andyantrim/apiutil/messages"
	"github.com/labstack/echo/v4"
)

func CustomerHTTPErrorHandler(err error, c echo.Context) {
	serr, ok := err.(*messages.GenericError)
	if !ok {
		serr = messages.NewServerError(err.Error())
	}

	c.JSON(serr.HTTPCode, serr)
}
