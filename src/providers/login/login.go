package login

import (
	"github.com/go-playground/validator/v10"
	"github.com/golang-etl/base-playwright/src/providers/context"
	logininterfaces "github.com/golang-etl/base-playwright/src/providers/login/interfaces"
	loginpreparers "github.com/golang-etl/base-playwright/src/providers/login/preparers"
	loginresponses "github.com/golang-etl/base-playwright/src/providers/login/responses"
	loginsteps "github.com/golang-etl/base-playwright/src/providers/login/steps"
	packagegeneralinterfaces "github.com/golang-etl/package-general/src/interfaces"
	packagehttpinterfaces "github.com/golang-etl/package-http/src/interfaces"
	packagehttputils "github.com/golang-etl/package-http/src/utils"
	packageusertokenmodels "github.com/golang-etl/package-user-token/src/models"
	"github.com/golang-etl/package-user-token/src/providers/usertoken"
)

type LoginProvider struct {
	CfgGoModuleName   string
	CfgDebug          bool
	Validator         *validator.Validate
	UserTokenModel    packageusertokenmodels.UserTokenModel
	ContextProvider   context.ContextProvider
	UserTokenProvider usertoken.UserTokenProvider
}

func (provider LoginProvider) GetLogin(shared *packagegeneralinterfaces.Shared, originalInputData logininterfaces.InputData) packagehttpinterfaces.Response {
	inputData := loginpreparers.DefaultInputData(originalInputData)

	err := provider.Validator.Struct(inputData)

	if err != nil {
		return packagehttputils.ValidationErrorHandlerToUnprocessableEntityResponse(err, inputData, map[string]string{
			"user.required":     "El RUT del usuario es obligatorio.",
			"password.required": "La clave es obligatoria.",
		})
	}

	extra := map[string]string{}

	context := provider.ContextProvider.NewContext(nil)
	defer provider.ContextProvider.CloseAll(context, shared)
	page := provider.ContextProvider.NewPage(context)
	loginsteps.GoToPage(page)
	loginsteps.FillAndSubmitForm(page, inputData)
	loginErrorResponse := loginsteps.WaitForLoginSuccess(page)

	if loginErrorResponse != nil {
		return *loginErrorResponse
	}

	userToken := loginsteps.CreateUserToken(provider.UserTokenModel, page, extra)

	return loginresponses.LoginSuccessResponse(loginresponses.LoginSuccessResponseBody{
		UserToken: userToken.Token,
	})
}
