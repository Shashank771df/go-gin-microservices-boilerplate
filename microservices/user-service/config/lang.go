package config

import (
	l "app/core/lang"
)

// Lang variables de mensajes de la aplicacion
var Lang = lang{
	NewKeyTest: l.Message{
		ID:      "TEST_KEY",
		Message: "The request was attended",
	},
}

type (
	lang struct {
		// Unimos la estructura generica del core con el de la aplicacion
		NewKeyTest l.Message
	}
)
