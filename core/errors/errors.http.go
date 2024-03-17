package errors

import (
	"app/core/lang"
	"encoding/json"
	"net/http"
)

var (
	// HTTPErrors permite crear HTTPError precreados
	HTTPErrors = httpErrors{}
)

type (
	// HTTPError .
	HTTPError struct {
		StatusCode int                    `json:"statusCode"`
		Success    bool                   `json:"success"`
		MessageID  string                 `json:"messageId"`
		Data       map[string]interface{} `json:"data,omitempty"`
		Message    string                 `json:"message"`
	}

	HTTPErrorPanic struct {
		StatusCode  int                    `json:"statusCode"`
		MessageID   string                 `json:"messageId"`
		ReportError bool                   `json:"reportError"`
		Data        map[string]interface{} `json:"data,omitempty"`
		HideData    bool                   `json:"hideData"`
		Message     string                 `json:"message"`
		Frames      []Frame                `json:"frames,omitempty"`
	}
	// httpErrors permite crear HTTPError precreados
	httpErrors struct{}
	// HTTPErrorConfig .
	HTTPErrorConfig struct {
		Message     lang.Message
		Data        map[string]interface{}
		HideData    bool
		ReportError bool
	}
)

// InternalServerError .
func (o httpErrors) InternalServerError(conf HTTPErrorConfig) {
	messageDefault := lang.GeneralInternalError

	if conf.Message.ID == "" {
		conf.Message.ID = messageDefault.ID
	}

	if conf.Message.Message == "" {
		conf.Message.Message = messageDefault.Message
	}

	httpError := HTTPErrorPanic{
		StatusCode:  http.StatusInternalServerError,
		MessageID:   conf.Message.ID,
		Message:     conf.Message.Message,
		ReportError: true,
		HideData:    conf.HideData,
		Data:        conf.Data,
		Frames:      ErrorStackTrace(2),
	}

	if !conf.HideData {
		httpError.Data = conf.Data
		httpError.Frames = ErrorStackTrace(2)
	}

	data, _ := json.Marshal(httpError)
	panic(string(data))
}

// BadRequest .
func (o httpErrors) BadRequest(conf HTTPErrorConfig) {
	messageDefault := lang.RequestBadRequest

	if conf.Message.ID == "" {
		conf.Message.ID = messageDefault.ID
	}

	if conf.Message.Message == "" {
		conf.Message.Message = messageDefault.Message
	}

	httpError := HTTPErrorPanic{
		StatusCode:  http.StatusBadRequest,
		MessageID:   conf.Message.ID,
		Message:     conf.Message.Message,
		ReportError: conf.ReportError,
		HideData:    conf.HideData,
		Data:        conf.Data,
		Frames:      ErrorStackTrace(2),
	}

	data, _ := json.Marshal(httpError)
	panic(string(data))
}

// NotFound .
func (o httpErrors) NotFound(conf HTTPErrorConfig) {
	messageDefault := lang.RequestFailedNotFound

	if conf.Message.ID == "" {
		conf.Message.ID = messageDefault.ID
	}

	if conf.Message.Message == "" {
		conf.Message.Message = messageDefault.Message
	}

	httpError := HTTPErrorPanic{
		StatusCode:  http.StatusNotFound,
		MessageID:   conf.Message.ID,
		Message:     conf.Message.Message,
		ReportError: conf.ReportError,
		HideData:    conf.HideData,
		Data:        conf.Data,
		Frames:      ErrorStackTrace(2),
	}

	data, _ := json.Marshal(httpError)
	panic(string(data))
}

// Unauthorized .
func (o httpErrors) Unauthorized(conf HTTPErrorConfig) {
	messageDefault := lang.RequestUnathorized

	if conf.Message.ID == "" {
		conf.Message.ID = messageDefault.ID
	}

	if conf.Message.Message == "" {
		conf.Message.Message = messageDefault.Message
	}

	httpError := HTTPErrorPanic{
		StatusCode:  http.StatusUnauthorized,
		MessageID:   conf.Message.ID,
		Message:     conf.Message.Message,
		ReportError: conf.ReportError,
		HideData:    conf.HideData,
		Data:        conf.Data,
		Frames:      ErrorStackTrace(2),
	}

	data, _ := json.Marshal(httpError)
	panic(string(data))
}

// Conflict .
func (o httpErrors) Conflict(conf HTTPErrorConfig) {
	messageDefault := lang.RequestConflict

	if conf.Message.ID == "" {
		conf.Message.ID = messageDefault.ID
	}

	if conf.Message.Message == "" {
		conf.Message.Message = messageDefault.Message
	}

	httpError := HTTPErrorPanic{
		StatusCode:  http.StatusConflict,
		MessageID:   conf.Message.ID,
		Message:     conf.Message.Message,
		ReportError: conf.ReportError,
		HideData:    conf.HideData,
		Data:        conf.Data,
		Frames:      ErrorStackTrace(2),
	}

	data, _ := json.Marshal(httpError)
	panic(string(data))
}
