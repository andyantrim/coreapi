package transport

import (
	"github.com/andyantrim/apiutil"
	"github.com/andyantrim/coreapi/connectors/mysql"
	"github.com/andyantrim/coreapi/service"
	apihttp "github.com/andyantrim/coreapi/transport/http"
	"github.com/labstack/echo/v4"
)

// InitV1 builds V1 resources
func InitV1(g *echo.Group) {
	resourceRepo := mysql.NewResourceRepo()
	resourceSrv := service.NewResourceService(resourceRepo)
	resourceController := apihttp.NewResourceController(resourceSrv)
	apiutil.RegisterRoutes(g, resourceController)
}
