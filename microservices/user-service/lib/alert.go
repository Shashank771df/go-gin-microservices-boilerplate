package lib

import (
	"app/core/alert"
	c "app/microservices/user-service/config"
)

var Alert alert.Alert = nil

func init() {
	Alert = alert.New(c.Env.Micro.Alert.Agent, alert.AlertProps{
		Enable: c.Env.Micro.Alert.Enable,
		Url:    c.Env.Micro.Alert.UrlBot,
		ChatId: c.Env.Micro.Alert.ChatId,
	})
}
