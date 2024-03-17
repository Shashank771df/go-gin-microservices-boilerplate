package core

import (
	"app/core/errors"
	"app/core/lang"
	"app/core/utils"
	"app/core/validator"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Interceptor .
//var Interceptor = interceptor{}

type (
	Interceptor struct{}

	ErrorItems struct {
		AppContext
		errors.HTTPErrorPanic
		TapFunc func(AppContext, errors.HTTPErrorPanic) error
	}

	HandlerAppFunc func(AppContext)
	PipeFunc       func(AppContext) interface{}
	GuardFunc      func(AppContext) (interface{}, error)
	ErrorFunc      func(AppContext, errors.HTTPError)
	TapFunc        func(AppContext, errors.HTTPErrorPanic)

	ExceptionItems struct {
		Error   ErrorFunc
		TapFunc TapFunc
	}

	InterceptorItems struct {
		Pipe    PipeFunc
		Guard   GuardFunc
		Handler HandlerAppFunc
		Error   ExceptionItems
	}

	InterceptorError struct {
		Error func(AppContext, errors.HTTPErrorPanic) error
	}

	ErrorsConfig struct {
		AppStage string
	}
)

// APIKey valida el apikey de las consultas a la aplicacion
func (i Interceptor) Request(items InterceptorItems) gin.HandlerFunc {
	return func(c *gin.Context) {
		appCtx := AppContext{
			Context: *c,
		}
		isPanic := true

		defer func() {
			if !isPanic {
				return
			}

			var err error
			var ok bool

			if r := recover(); r != nil {
				err, ok = r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
			}

			i.exceptionIterceptor(appCtx, items.Error, err)
		}()

		if items.Pipe != nil {
			appCtx.Data = items.Pipe(appCtx)
			val := validator.New()
			err := val.Struct(appCtx.Data)
			if err != nil {
				errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
					Message: lang.Message{
						Message: err.Error(),
					},
				})
			}
		}

		if items.Guard != nil {
			data, err := items.Guard(appCtx)

			if err != nil {
				errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
					Message: lang.Message{
						Message: err.Error(),
					},
				})
			}
			utils.UtilsMap{}.InterfaceToStruct(data, &appCtx.Sec)
		}

		if appCtx.Sec == nil {
			appCtx.Sec = &SecuritySessionValidate{}
		}

		items.Handler(appCtx)
		isPanic = false
	}
}

func (o Interceptor) exceptionIterceptor(c AppContext, ex ExceptionItems, e error) {
	// Si es un error de la aplicacion
	httpError := errors.HTTPErrorPanic{}
	httpErrorReport := errors.HTTPError{}
	utilMap := utils.UtilsMap{}
	err := utilMap.ErrorToStruct(e, &httpError)

	//Validamos si es un httpError o un error generico
	if err == nil {
		utilMap.InterfaceToStruct(httpError, &httpErrorReport)
		if httpError.HideData {
			httpErrorReport.Data = nil
		}
	} else {
		httpErrorReport = errors.HTTPError{
			StatusCode: http.StatusInternalServerError,
			MessageID:  lang.GeneralInternalError.ID,
			Message:    e.Error(),
		}
	}

	if ex.TapFunc != nil {
		if err != nil {
			httpError.Message = e.Error()
		}
		go ex.TapFunc(c, httpError)
	}

	if ex.Error != nil {
		ex.Error(c, httpErrorReport)
	} else {
		o.defaultError(c, httpErrorReport)
	}
}

func (o Interceptor) defaultError(c AppContext, httpError errors.HTTPError) {
	c.JSON(httpError.StatusCode, httpError)
}
