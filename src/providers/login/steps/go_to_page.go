package loginsteps

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

func GoToPage(page playwright.Page) {
	_, err := page.Goto("https://example.com")

	if err != nil {
		panic(fmt.Errorf("error al ir a la pagina: %w", err))
	}
}
