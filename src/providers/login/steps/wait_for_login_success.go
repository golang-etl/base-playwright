package loginsteps

import (
	"fmt"
	"strings"

	loginresponses "github.com/golang-etl/base-playwright/src/providers/login/responses"
	packagehttpinterfaces "github.com/golang-etl/package-http/src/interfaces"
	"github.com/playwright-community/playwright-go"
)

func WaitForLoginSuccess(page playwright.Page) *packagehttpinterfaces.Response {
	defaultErrorMessageSelector := "#mensajeError"
	userNameSelector := "#userName"
	currentMessageLocator := page.Locator(defaultErrorMessageSelector + ", " + userNameSelector).First()

	err := currentMessageLocator.WaitFor(playwright.LocatorWaitForOptions{State: playwright.WaitForSelectorStateVisible})

	if err != nil {
		panic(fmt.Errorf("error al esperar que el mensaje para reconocer que ocurrió esté visible: %w", err))
	}

	classAttr, err := currentMessageLocator.GetAttribute("class")

	if err != nil {
		panic(fmt.Errorf("error al obtener el id del mensaje para reconocer que ocurrió: %w", err))
	}

	if strings.Contains(classAttr, defaultErrorMessageSelector[1:]) {
		locator := page.Locator(defaultErrorMessageSelector).First()
		content, err := locator.TextContent()

		if err != nil {
			panic(fmt.Errorf("error al obtener el contenido del mensaje de error: %w", err))
		}

		if strings.Contains(content, "Detectamos que tiene una sesión activa") {
			response := loginresponses.LoginActiveSessionErrorResponse()

			return &response
		}

		if strings.Contains(content, "Los datos ingresados no corresponden") {
			response := loginresponses.LoginInvalidCredentialsResponse()

			return &response
		}

		response := loginresponses.LoginUnknownErrorResponse()

		return &response
	}

	if strings.Contains(classAttr, userNameSelector[1:]) {
		page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
			State: playwright.LoadStateNetworkidle,
		})

		return nil
	}

	panic(fmt.Errorf("error al esperar el inicio de sesión: %w", err))
}
