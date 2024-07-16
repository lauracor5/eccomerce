package response

import (
	"ecommerce/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const (
	BindFailed      = "bind_failed"
	Ok              = "ok"
	RecordCreated   = "record_created"
	RecordUpdated   = "record_updated"
	RecordDeleted   = "record_deleted"
	UnexpextedError = "unexpected_error"
	AuthError       = "authorization_error"
)

type API struct {
}

func New() API {
	return API{}
}

func (a API) OK(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responnses{{Code: Ok, Message: "listo!"}},
	}
}

func (a API) Created(data interface{}) (int, model.MessageResponse) {
	return http.StatusCreated, model.MessageResponse{
		Data:     data,
		Messages: model.Responnses{{Code: RecordCreated, Message: "listo!"}},
	}
}

func (a API) Updated(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responnses{{Code: RecordUpdated, Message: "listo!"}},
	}
}

func (a API) Deleted(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responnses{{Code: RecordDeleted, Message: "listo!"}},
	}

}

func (a API) BindFailed(err error) error {
	log := logrus.New()
	e := model.NewError()
	e.Err = err
	e.Code = BindFailed
	e.StatusHTTP = http.StatusBadRequest
	e.Who = "c.Bind"

	log.Warnf("%s", e.Error())
	return &e
}

func (a API) Error(c echo.Context, who string, err error) *model.Error {
	log := logrus.New()
	e := model.NewError()
	e.Err = err
	e.APIMessage = "uy metimos la pata, disculpa lo solucionaremos"
	e.Code = UnexpextedError
	e.StatusHTTP = http.StatusInternalServerError
	e.Who = who

	userID, ok := c.Get("userID").(uuid.UUID)
	if !ok {
		log.Errorf("cannot get/parse uuid from userID")
	}

	e.UserID = userID.String()
	log.Errorf("%s", e.Error())
	return &e
}
