package loginresponses

import (
	"net/http"

	"github.com/golang-etl/package-http/src/consts"
	packagehttpinterfaces "github.com/golang-etl/package-http/src/interfaces"
)

type LoginSuccessResponseBody struct {
	UserToken string `json:"userToken"`
}

func LoginSuccessResponse(body LoginSuccessResponseBody) packagehttpinterfaces.Response {
	return packagehttpinterfaces.Response{
		StatusCode: http.StatusOK,
		Headers:    consts.HeaderContentType.JSON,
		Body:       body,
	}
}

func LoginInvalidCredentialsResponse() packagehttpinterfaces.Response {
	return packagehttpinterfaces.Response{
		StatusCode: http.StatusUnauthorized,
		Headers:    consts.HeaderContentType.JSON,
		Body: packagehttpinterfaces.ResponseBodyError{
			Message:   "Los datos de autenticación son incorrectos.",
			ErrorCode: "INVALID_CREDENTIALS",
		},
	}
}

func LoginInvalidCompanyResponse() packagehttpinterfaces.Response {
	return packagehttpinterfaces.Response{
		StatusCode: http.StatusUnauthorized,
		Headers:    consts.HeaderContentType.JSON,
		Body: packagehttpinterfaces.ResponseBodyError{
			Message:   "La empresa a la que intentas iniciar sesión no esta vinculada a esta cuenta de usuario.",
			ErrorCode: "UNLINKED_COMPANY_ERROR",
		},
	}
}

func LoginUnknownErrorResponse() packagehttpinterfaces.Response {
	return packagehttpinterfaces.Response{
		StatusCode: http.StatusInternalServerError,
		Headers:    consts.HeaderContentType.JSON,
		Body: packagehttpinterfaces.ResponseBodyError{
			Message:   "Error desconocido. Intente nuevamente.",
			ErrorCode: "UNKNOWN_LOGIN_ERROR",
		},
	}
}

func LoginActiveSessionErrorResponse() packagehttpinterfaces.Response {
	return packagehttpinterfaces.Response{
		StatusCode: http.StatusUnauthorized,
		Headers:    consts.HeaderContentType.JSON,
		Body: packagehttpinterfaces.ResponseBodyError{
			Message:   "Ya existe una sesión activa en esta cuenta. Intente iniciar sesión en unos minutos.",
			ErrorCode: "ALREADY_ACTIVE_SESSION_ERROR",
		},
	}
}
