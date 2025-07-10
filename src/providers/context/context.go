package context

import (
	"fmt"

	"github.com/golang-etl/base-playwright/src/config"
	"github.com/golang-etl/base-playwright/src/providers/browser"
	packagegeneralinterfaces "github.com/golang-etl/package-general/src/interfaces"
	packagegeneralutils "github.com/golang-etl/package-general/src/utils"
	packageplaywrightutils "github.com/golang-etl/package-playwright/src/utils"
	packageusertokenmodels "github.com/golang-etl/package-user-token/src/models"

	"github.com/playwright-community/playwright-go"
)

type ContextProvider struct {
	Cfg             *config.Config
	BrowserProvider *browser.BrowserProvider
}

func (provider *ContextProvider) NewContext(userTokenInstance *packageusertokenmodels.UserToken) playwright.BrowserContext {
	newContextOptions := playwright.BrowserNewContextOptions{
		UserAgent: playwright.String(provider.Cfg.UserAgentHeader),
		Viewport: &playwright.Size{
			Width:  1280,
			Height: 720,
		},
		StorageState: provider.GetOptionalStorageState(userTokenInstance),
	}

	context, err := provider.BrowserProvider.Browser.NewContext(newContextOptions)

	if err != nil {
		panic(fmt.Errorf("error al crear un nuevo contexto: %w", err))
	}

	err = context.Route("**/*", func(route playwright.Route) {
		headers := route.Request().Headers()
		headers["sec-ch-ua"] = provider.Cfg.SecChUaHeader

		route.Continue(playwright.RouteContinueOptions{
			Headers: headers,
		})
	})

	if err != nil {
		panic(fmt.Errorf("error al agregar cabeceras al contexto: %w", err))
	}

	if provider.Cfg.TraceEnabled {
		packageplaywrightutils.StartTracing(context)
	}

	provider.LoadSessionStorage(context, userTokenInstance)
	provider.LoadScripts(context)

	return context
}

func (provider *ContextProvider) GetOptionalStorageState(userTokenInstance *packageusertokenmodels.UserToken) *playwright.OptionalStorageState {
	if userTokenInstance == nil {
		return nil
	}

	return packageplaywrightutils.GenerateOptionalStorageState(userTokenInstance.Context)
}

func (provider *ContextProvider) LoadSessionStorage(context playwright.BrowserContext, userTokenInstance *packageusertokenmodels.UserToken) {
	if userTokenInstance == nil {
		return
	}

	sessionStorageInitScript := packageplaywrightutils.GenerateSessionStorageInitScript(userTokenInstance.SessionStorage)

	err := context.AddInitScript(playwright.Script{Content: playwright.String(sessionStorageInitScript)})

	if err != nil {
		panic(fmt.Errorf("error al añadir un script de inicialización (sessionStorage): %w", err))
	}
}

func (provider *ContextProvider) LoadScripts(context playwright.BrowserContext) {
	err := context.AddInitScript(playwright.Script{Content: playwright.String("Object.defineProperty(navigator,'webdriver',{get: () => false})")})

	if err != nil {
		panic(fmt.Errorf("error al añadir un script de inicialización (webdriver): %w", err))
	}
}

func (provider *ContextProvider) NewPage(context playwright.BrowserContext) playwright.Page {
	page, err := context.NewPage()

	if err != nil {
		panic(fmt.Errorf("error al abrir una nueva página: %w", err))
	}

	return page
}

func (provider *ContextProvider) CloseAll(context playwright.BrowserContext, shared *packagegeneralinterfaces.Shared) {
	if provider.Cfg.TraceEnabled {
		traceToken, fileName, fileAbsolutePath := packageplaywrightutils.StopTracing(context)
		shared.TraceToken = &traceToken

		if !packagegeneralutils.IsRuntimeEnvironmentLocal() {
			packagegeneralutils.UploadFileToBucketEnvironment(provider.Cfg.TraceBucket, fileAbsolutePath, fileName)
		}
	}

	if context != nil {
		context.Close()
	}
}
