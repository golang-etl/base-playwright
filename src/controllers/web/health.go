package web

import (
	"github.com/golang-etl/base-playwright/src/providers/health"
	packagegeneralinterfaces "github.com/golang-etl/package-general/src/interfaces"
	packagehttputils "github.com/golang-etl/package-http/src/utils"
	"github.com/labstack/echo/v4"
)

func GetHealth(healthProvider health.HealthProvider) func(c echo.Context) error {
	return func(c echo.Context) error {
		var shared *packagegeneralinterfaces.Shared = &packagegeneralinterfaces.Shared{}

		defer packagehttputils.InternalServerErrorResponse(c, shared, healthProvider.CfgGoModuleName, healthProvider.CfgDebug, healthProvider.CfgDebug)

		return packagehttputils.AdaptEchoResponse(c, shared, healthProvider.GetHealth(shared))
	}
}
