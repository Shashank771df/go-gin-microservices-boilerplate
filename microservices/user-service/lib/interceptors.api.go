package lib

import (
	"app/core"
	"app/core/alert"
	"app/core/errors"
	"app/core/logger"
	"app/microservices/user-service/config"
	"app/microservices/user-service/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type interceptorAPI struct {
	i core.Interceptor
}

var InterceptorAPI interceptorAPI

func init() {
	InterceptorAPI = interceptorAPI{
		i: core.Interceptor{},
	}
}

func (o interceptorAPI) ErrorFunc() core.ErrorFunc {
	return func(c core.AppContext, e errors.HTTPError) {
		data := map[string]interface{}{
			"Error": "Implement custom error here",
		}
		c.JSON(http.StatusInternalServerError, data)
	}
}

func (o interceptorAPI) TapFunc() core.TapFunc {
	return func(c core.AppContext, e errors.HTTPErrorPanic) {
		log.Log.Error(logger.LogInfo{
			Key:   config.Env.Micro.App.Name,
			Value: e.Message,
			Extra: map[string]interface{}{
				"Data":      e.Data,
				"Backtrace": e.Frames,
				"Context":   c.Data,
				"Security":  c.Sec,
				"URL":       c.Request.URL.Path,
				"Host":      c.RemoteIP(),
			},
		})

		if e.ReportError {
			Alert.SendMessage(alert.AlertInfo{
				Key:   config.Env.Micro.App.Name,
				Value: e.Message,
				Extra: map[string]interface{}{
					"Backtrace": e.Frames,
					"Data":      e.Data,
				},
			})
		}
	}
}

func (o interceptorAPI) Request(items core.InterceptorItems) gin.HandlerFunc {
	items.Error = core.ExceptionItems{
		TapFunc: o.TapFunc(),
		//Error: o.ErrorFunc(), //Uncomment for custom error
	}
	return o.i.Request(items)
}
