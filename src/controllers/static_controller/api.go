package static_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"wkey-core/src/data/errors"
)

func (controller *Controller) Health(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "OK")
}

func (controller *Controller) RouteNotFound(ctx echo.Context) error {
	return controller.NotFound(ctx, errors.RouteNotFound)
}
