package response

import (
	"ecommerce/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HTTPErrorHandler(err error, c echo.Context) {
	e, ok := err.(*model.Error)
	if ok {
		_ = c.JSON(getResponseError(e))
		return
	}

	// check error
	if echoErr, ok := err.(*echo.HTTPError); ok {
		msg, ok := echoErr.Message.(string)
		if !ok {
			msg = "Ups algo inesperado ocurrio"
		}

		_ = c.JSON(echoErr.Code, model.MessageResponse{
			Errors: model.Responnses{
				{Code: UnexpextedError, Message: msg},
			},
		})
		return
	}
}

func getResponseError(err *model.Error) (int, model.MessageResponse) {
	outputstatus := 0
	outputResponse := model.MessageResponse{}

	if !err.HasCode() {
		err.Code = UnexpextedError
	}

	if err.HasData() {
		outputResponse.Data = err.Data
	}

	if !err.HasStatus() {
		err.StatusHTTP = http.StatusInternalServerError

	}

	outputstatus = err.StatusHTTP
	outputResponse.Errors = model.Responnses{model.Response{
		Code:    err.Code,
		Message: err.APIMessage,
	}}

	return outputstatus, outputResponse

}
