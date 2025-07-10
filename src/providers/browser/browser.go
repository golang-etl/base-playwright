package browser

import (
	"fmt"

	"github.com/golang-etl/base-playwright/src/config"

	"github.com/playwright-community/playwright-go"
)

type BrowserProvider struct {
	Cfg     *config.Config
	PW      *playwright.Playwright
	Browser playwright.Browser
}

func (provider *BrowserProvider) OpenBrowser() (*playwright.Playwright, playwright.Browser) {
	pw, err := playwright.Run()

	if err != nil {
		panic(fmt.Errorf("error running playwright: %w", err))
	}

	browserLaunchOptions := playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(!provider.Cfg.DisabledHeadlessInBrowser),
		Devtools: playwright.Bool(provider.Cfg.EnabledDevtoolsInBrowser),
	}

	if provider.Cfg.ProxyServer != "" {
		browserLaunchOptions.Proxy = &playwright.Proxy{
			Server:   provider.Cfg.ProxyServer,
			Username: playwright.String(provider.Cfg.ProxyUsername),
			Password: playwright.String(provider.Cfg.ProxyPassword),
		}
	}

	browser, err := pw.Chromium.Launch(browserLaunchOptions)

	if err != nil {
		panic(fmt.Errorf("error launching browser: %w", err))
	}

	provider.PW = pw
	provider.Browser = browser

	return pw, browser
}

func (provider *BrowserProvider) CloseAll() {
	if provider.Browser != nil {
		provider.Browser.Close()
	}

	if provider.PW != nil {
		provider.PW.Stop()
	}
}
