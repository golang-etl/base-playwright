package loginsteps

import (
	"fmt"

	logininterfaces "github.com/golang-etl/base-playwright/src/providers/login/interfaces"
	"github.com/playwright-community/playwright-go"
)

func FillAndSubmitForm(page playwright.Page, inputData logininterfaces.InputData) {
	userInput := page.Locator("#user").First()
	passwordInput := page.Locator("#password").First()
	submitButton := page.Locator("#submit").First()

	if err := userInput.WaitFor(playwright.LocatorWaitForOptions{State: playwright.WaitForSelectorStateVisible}); err != nil {
		panic(fmt.Errorf("error al esperar que el campo usuario esté visible: %w", err))
	}

	if err := userInput.Click(); err != nil {
		panic(fmt.Errorf("error al clickear el campo usuario: %w", err))
	}

	if err := userInput.Fill(inputData.User); err != nil {
		panic(fmt.Errorf("error al rellenar el campo usuario: %w", err))
	}

	if err := passwordInput.WaitFor(playwright.LocatorWaitForOptions{State: playwright.WaitForSelectorStateVisible}); err != nil {
		panic(fmt.Errorf("error al esperar que el campo clave esté visible: %w", err))
	}

	if err := passwordInput.Click(); err != nil {
		panic(fmt.Errorf("error al clickear el campo clave: %w", err))
	}

	if err := passwordInput.Fill(inputData.Password); err != nil {
		panic(fmt.Errorf("error al rellenar el campo clave: %w", err))
	}

	if err := submitButton.WaitFor(playwright.LocatorWaitForOptions{State: playwright.WaitForSelectorStateVisible}); err != nil {
		panic(fmt.Errorf("error al esperar que el botón de enviar esté visible: %w", err))
	}

	if err := submitButton.Click(); err != nil {
		panic(fmt.Errorf("error al clickear el botón de enviar: %w", err))
	}
}
