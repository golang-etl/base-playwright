package loginsteps

import (
	"encoding/json"
	"fmt"

	packagegeneralutils "github.com/golang-etl/package-general/src/utils"
	packageplaywrightutils "github.com/golang-etl/package-playwright/src/utils"
	packageusertokenmodels "github.com/golang-etl/package-user-token/src/models"
	"github.com/playwright-community/playwright-go"
)

func CreateUserToken(userTokenModel packageusertokenmodels.UserTokenModel, page playwright.Page, extra map[string]string) packageusertokenmodels.UserToken {
	_, err := page.WaitForFunction(`() => sessionStorage.getItem("token") !== null`, playwright.PageWaitForFunctionOptions{
		Timeout: playwright.Float(2000),
	})

	if err != nil {
		panic(fmt.Errorf("La clave 'token' no se estableció en sessionStorage dentro del tiempo límite: %w", err))
	}

	storageState, err := page.Context().StorageState()

	if err != nil {
		panic(fmt.Errorf("error al obtener el estado de la sesion: %w", err))
	}

	storageStateJSON, err := json.Marshal(storageState)

	if err != nil {
		panic(fmt.Errorf("error al transformar el estado de la sesion: %w", err))
	}

	_, sessionStorage, err := packageplaywrightutils.GetStorage(page)

	if err != nil {
		panic(fmt.Errorf("error al obtener el 'sessionStorage': %w", err))
	}

	token := packagegeneralutils.GenerateRandToken(128)

	userToken := packageusertokenmodels.UserToken{
		Token:          token,
		Context:        string(storageStateJSON),
		SessionStorage: sessionStorage,
		Extra:          extra,
	}

	err = userTokenModel.Insert(userToken)

	if err != nil {
		panic(fmt.Errorf("error al insertar el token de usuario: %w", err))
	}

	return userToken
}
