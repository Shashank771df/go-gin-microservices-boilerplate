package core

import (
	"app/core/errors"
	"app/core/lang"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppContext struct {
	gin.Context
	Data interface{}
	Sec  *SecuritySessionValidate
}

func (c AppContext) StatusOK(data interface{}) error {
	c.JSON(http.StatusOK, data)

	return nil
}

func (c AppContext) BindPipe(i interface{}) interface{} {
	var data []byte
	var err error

	if c.Request.Body != nil {
		data, err = io.ReadAll(c.Request.Body)
	}

	if err != nil {
		return err
	}

	if len(data) > 0 {
		err = json.Unmarshal(data, i)

		if err != nil {
			errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
				Message: lang.Message{
					Message: err.Error(),
				},
			})
		}
	}

	err = c.BindUri(i)

	if err != nil {
		return nil
	}

	err = c.BindQuery(i)

	if err != nil {
		return nil
	}

	err = c.BindHeader(i)

	if err != nil {
		return nil
	}

	return i
}
