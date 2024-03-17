package environments

import (
	"app/core/alert"
	"app/core/errors"
	"app/core/lang"
	"app/core/utils"
	"app/microservices/user-service/config"
	"app/microservices/user-service/lib"
)

func (o Handler) ItemAlert() ItemAlert {
	Enable := config.Env.Micro.Alert.Enable
	UrlBot := config.Env.Micro.Alert.UrlBot
	ChatId := config.Env.Micro.Alert.ChatId

	return ItemAlert{
		Enable: &Enable,
		Url:    &UrlBot,
		ChatId: &ChatId,
	}
}

func (o Handler) UpdateAlert(data interface{}) Update {
	var params ItemAlert
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
			HideData: false,
		})
	}

	before := o.ItemAlert()
	var util utils.Util
	util.SetValue(&config.Env.Micro.Alert.Enable, params.Enable)
	util.SetValue(&config.Env.Micro.Alert.UrlBot, params.Url)
	util.SetValue(&config.Env.Micro.Alert.ChatId, params.ChatId)

	lib.Alert.Initialize(
		alert.AlertProps{
			Url:    config.Env.Micro.Alert.UrlBot,
			ChatId: config.Env.Micro.Alert.ChatId,
			Enable: config.Env.Micro.Alert.Enable,
		},
	)

	return Update{
		Before:  before,
		Current: o.ItemAlert(),
	}
}
