package core

import (
	"app/core/errors"
	"app/core/lang"
	"app/core/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Middleware .
type (
	Middleware struct{}
)

/*func (m Middleware) AppContext() gin.HandlerFunc {
	return func() (next gin.HandlerFunc) {
		return func(c *gin.Context) error {
			appContext := &AppContext{Context: c}
			return next(appContext)
		}
	}
}*/

func (m Middleware) AddCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header(HeaderAccessControlAllowOrigin, "*")
		c.Header(HeaderAccessControlAllowCredentials, "true")

		if strings.ToLower(c.Request.Method) != "options" {
			c.Next()
			return
		}

		requestHeaders := c.Request.Header.Get(HeaderAccessControlRequestHeaders)
		if requestHeaders != "" {
			c.Header(HeaderAccessControlAllowHeaders, requestHeaders)
		}

		requestMethods := c.Request.Header.Get(HeaderAccessControlRequestMethod)
		if requestMethods != "" {
			c.Header(HeaderAccessControlAllowMethods, requestMethods)
			c.Header(HeaderAllow, requestMethods)
		}

		c.Header(HeaderAccessControlMaxAge, "content-type, content-length, etag")
		c.Header(HeaderAccessControlExposeHeaders, "600")

		c.AbortWithStatus(http.StatusOK)
	}
}

func (i Middleware) logErrorResponse(data interface{}) {
	err, _ := json.MarshalIndent(data, "", "\t")
	fmt.Printf("ERROR: %v\n", string(err))
}

func (i Middleware) ErrorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, e := range c.Errors {
			// Si es un error de echo
			if err, ok := e.Err.(*gin.Error); ok {
				i.logErrorResponse(map[string]interface{}{
					"message": lang.Message{
						ID:      lang.FrameworkInternalError.ID,
						Message: e.Error(),
					},
					"backTrace": errors.ErrorStackTrace(2),
				})
				c.JSON(http.StatusInternalServerError, errors.HTTPError{
					MessageID: lang.FrameworkInternalError.ID,
					Message:   fmt.Sprintf("%v", err.Error()),
				})
				return
			}

			// Si es un error de la aplicacion
			httpError := errors.HTTPErrorPanic{}
			httpErrorReport := errors.HTTPError{}
			utilMap := utils.UtilsMap{}
			err := utilMap.ErrorToStruct(e, &httpError)

			if err == nil {
				i.logErrorResponse(map[string]interface{}{
					"message": lang.Message{
						ID:      lang.FrameworkInternalError.ID,
						Message: e.Error(),
					},
					"backTrace": errors.ErrorStackTrace(2),
				})

				utilMap.InterfaceToStruct(httpError, &httpErrorReport)

				c.JSON(httpError.StatusCode, httpErrorReport)
				return
			}

			//Default error
			i.logErrorResponse(map[string]interface{}{
				"message": lang.Message{
					ID:      lang.GeneralInternalError.ID,
					Message: e.Error(),
				},
				"backTrace": errors.ErrorStackTrace(2),
			})
			c.JSON(http.StatusInternalServerError, errors.HTTPError{
				MessageID: lang.GeneralInternalError.ID,
				Message:   e.Error(),
			})
		}
	}
}

func (m Middleware) EchoRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				c.Error(err)
			}
		}()
		c.Next()
	}
}
