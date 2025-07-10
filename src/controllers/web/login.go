package web

import (
	"fmt"

	"github.com/golang-etl/base-playwright/src/providers/login"
	logininterfaces "github.com/golang-etl/base-playwright/src/providers/login/interfaces"
	packagegeneralinterfaces "github.com/golang-etl/package-general/src/interfaces"
	packagehttputils "github.com/golang-etl/package-http/src/utils"
	"github.com/labstack/echo/v4"
)

func Login(loginProvider login.LoginProvider) func(c echo.Context) error {
	return func(c echo.Context) error {
		var shared *packagegeneralinterfaces.Shared = &packagegeneralinterfaces.Shared{}
		var bindData struct {
			User     string `json:"user"`
			Password string `json:"password"`
		}

		defer packagehttputils.InternalServerErrorResponse(c, shared, loginProvider.CfgGoModuleName, loginProvider.CfgDebug, loginProvider.CfgDebug)

		if err := c.Bind(&bindData); err != nil {
			panic(fmt.Errorf("error binding request body: %w", err))
		}

		return packagehttputils.AdaptEchoResponse(c, shared, loginProvider.GetLogin(shared, logininterfaces.InputData{
			User:     bindData.User,
			Password: bindData.Password,
		}))
	}
}
