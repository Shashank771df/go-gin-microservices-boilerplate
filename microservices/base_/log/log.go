package log

import (
	"app/core/logger"
	"app/microservices/base_/config"
)

var Log logger.Logger

func init() {
	Log = logger.New(
		config.Env.Micro.Log.Type, logger.LogProps{
			Name:           config.Env.Micro.Log.Name,
			ProjectID:      config.Env.Micro.Log.ProjectId,
			MaxFiles:       config.Env.Micro.Log.MaxFiles,
			FilePrefix:     config.Env.Micro.Log.FilePrefix,
			FileSize:       config.Env.Micro.Log.FileSize,
			Path:           config.Env.Micro.Log.Path,
			ConsoleEnabled: config.Env.Micro.Log.ConsoleEnabled,
			FileEnabled:    config.Env.Micro.Log.FileEnabled,
			MaxQueue:       config.Env.Micro.Log.MaxQueue,
			LogLevel:       config.Env.Micro.Log.LogLevel,
		},
	)
}
