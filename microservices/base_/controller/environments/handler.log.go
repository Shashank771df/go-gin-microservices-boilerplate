package environments

import (
	"app/core/errors"
	"app/core/lang"
	"app/core/utils"
	"app/microservices/base_/config"
	"app/microservices/base_/lib"
	"app/microservices/base_/log"
)

func (o Handler) ItemLog() ItemLog {
	LogLevel := config.Env.Micro.Log.LogLevel

	return ItemLog{
		LogLevel: &LogLevel,
	}
}

func (o Handler) UpdateLog(data interface{}) Update {
	var params ItemLog
	err := utils.UtilsMap{}.InterfaceToStruct(data, &params)

	if err != nil {
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: "Invalid pipe data",
			},
			Data: map[string]interface{}{
				"data":  data,
				"model": ItemAlert{},
			},
			HideData: true,
		})
	}

	before := o.ItemLog()
	var util utils.Util
	util.SetValue(&config.Env.Micro.Log.LogLevel, params.LogLevel)

	log.Log.SetLogLevel(config.Env.Micro.Log.LogLevel)
	lib.DB.SetLogger(log.Log)

	return Update{
		Before:  before,
		Current: o.ItemLog(),
	}
}
