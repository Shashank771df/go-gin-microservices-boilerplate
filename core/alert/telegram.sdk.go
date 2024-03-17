package alert

import (
	"app/core/network"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type TelegramSDK struct {
	urlBot     string
	chatId     string
	parse_mode string
	enable     bool
}

const TELEGRAM_MAXLEN = 4096

func (obj *TelegramSDK) Initialize(props AlertProps) error {
	obj.urlBot = props.Url
	obj.chatId = props.ChatId
	obj.parse_mode = props.ParseMode
	obj.enable = props.Enable

	if !obj.enable {
		return nil
	}

	if obj.urlBot == "" {
		return errors.New("urlBot is needed by TelegramSDK")
	}

	if obj.chatId == "" {
		return errors.New("chatId is needed by TelegramSDK")
	}

	if obj.parse_mode == "" {
		obj.parse_mode = "Markdown"
	}

	return nil
}

func (obj TelegramSDK) SendMessage(log AlertInfo) error {
	if !obj.enable {
		return nil
	}

	var extraData string = ""
	extra, _ := json.Marshal(log.Extra)

	if log.Extra != nil {
		extraData = string(extra)
	}

	message := fmt.Sprintf("%s```\n%s\n%s```", log.Key, log.Value, extraData)

	if len(message) > TELEGRAM_MAXLEN {
		message = fmt.Sprintf("%s```\n%s\n%s```", log.Key, log.Value, "")
	}

	message = strings.ReplaceAll(message, "_", "")
	message = strings.ReplaceAll(message, "*", "")

	response := network.HttpClient{}.Call(network.HttpClientRequest{
		Method: "POST",
		Url:    obj.urlBot,

		Data: map[string]interface{}{
			"chat_id":    obj.chatId,
			"parse_mode": obj.parse_mode,
			"text":       message,
		},
	})

	if response.Error != nil {
		return response.Error
	}

	if !response.IsStatusIn2xx {
		errMessage := fmt.Sprintf("%v", response.Data)
		return errors.New("telegram " + response.StatusMsg + " error: " + errMessage)
	}

	return nil
}
