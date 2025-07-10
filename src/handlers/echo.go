package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/golang-etl/base-playwright/src/config"
	"github.com/golang-etl/base-playwright/src/controllers/web"
	"github.com/golang-etl/base-playwright/src/database"
	"github.com/golang-etl/base-playwright/src/providers/browser"
	"github.com/golang-etl/base-playwright/src/providers/context"
	"github.com/golang-etl/base-playwright/src/providers/health"
	"github.com/golang-etl/base-playwright/src/providers/login"
	packagegeneralutils "github.com/golang-etl/package-general/src/utils"
	packageusertokenmodels "github.com/golang-etl/package-user-token/src/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		panic(fmt.Errorf("error loading config: %w", err))
	}

	mainDB := database.MainDB{}
	mainDB.Connect(cfg.MongoDBURI)
	mainDB.Ping(cfg.MongoDBDatabaseName)
	defer mainDB.Disconnect()

	e := echo.New()
	e.Use(middleware.Recover())

	mainValidator := packagegeneralutils.ValidatorNewWithTagNameInJson(validator.New())

	userTokenModel := packageusertokenmodels.UserTokenModel{Client: mainDB.Client, Secret: cfg.SecretKeyUserTokenData, Database: cfg.MongoDBDatabaseName}

	browserProvider := browser.BrowserProvider{Cfg: cfg}
	contextProvider := context.ContextProvider{Cfg: cfg, BrowserProvider: &browserProvider}
	healthProvider := health.HealthProvider{CfgGoModuleName: cfg.GoModuleName, CfgDebug: cfg.Debug, MongoClient: mainDB.Client}
	loginProvider := login.LoginProvider{CfgGoModuleName: cfg.GoModuleName, CfgDebug: cfg.Debug, Validator: mainValidator, UserTokenModel: userTokenModel, ContextProvider: contextProvider}

	browserProvider.OpenBrowser()
	defer browserProvider.CloseAll()

	e.GET("/health", web.GetHealth(healthProvider))
	e.POST("/login", web.Login(loginProvider))

	e.Logger.Fatal(e.Start(cfg.EchoAddress))
}
