package logger

import (
	"app/core/date"
	"encoding/json"
	"fmt"
	"os"
)

type DefaultLogger struct {
	logEnabled LogEnabled
}

func (obj *DefaultLogger) Init(props LogProps) {
	obj.SetLogLevel(props.LogLevel)
}

func (obj *DefaultLogger) SetLogLevel(value string) {
	obj.logEnabled.SetLogLevel(value)
}

func (obj DefaultLogger) Error(log LogInfo) {
	if !obj.logEnabled.Error {
		return
	}

	obj.write(RED, log)
}

func (obj DefaultLogger) Warn(log LogInfo) {
	if !obj.logEnabled.Warning {
		return
	}

	obj.write(YELLOW, log)
}

func (obj *DefaultLogger) Fatal(log LogInfo) {
	obj.write(RED, log)
	os.Exit(1)
}

func (obj DefaultLogger) Info(log LogInfo) {
	if !obj.logEnabled.Info {
		return
	}

	obj.write(CYAN, log)
}

func (obj DefaultLogger) Debug(log LogInfo) {
	if !obj.logEnabled.Debug {
		return
	}

	obj.write(GREEN, log)
}

func (obj DefaultLogger) write(keyColor string, log LogInfo) {
	var udate date.UtilDate
	date := udate.FormatMilliseconds(udate.Now())

	keyStr := fmt.Sprintf("%s%s%s", keyColor, log.Key, RESET)

	var extraData string = ""
	extra, _ := json.MarshalIndent(log.Extra, "", "\t")

	if len(log.Extra) > 0 {
		extraData = string(extra)
	}

	if extraData != "" {
		fmt.Printf("%s | %s | %s\n%s\n", keyStr, date, log.Value, extraData)
	} else {
		fmt.Printf("%s | %s | %s\n", keyStr, date, log.Value)
	}
}
