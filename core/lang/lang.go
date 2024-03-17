package lang

type Message struct {
	ID      string	`json:"messageId"`
	Message string	`json:"message"`
}

var RequestCompleted = Message {
	ID:      "REQUEST_SUCCESS_COMPLETED",
	Message: "The request was attended",
}

var RequestUpdCompleted = Message{
	ID:      "REQUEST_SUCCESS_UPD_COMPLETE",
	Message: "The request was attended",
}

var RequestBadRequest = Message {
	ID:      "REQUEST_BAD_REQUEST",
	Message: "Request had invalid information",
}

var RequestUnathorized = Message{
	ID:      "REQUEST_SECURITY_UNAUTHORIZED",
	Message: "Request was not processed, try again",
}

var RequestConflict = Message{
	ID:      "REQUEST_GENERAL_CONFLICT",
	Message: "Request has a conflict",
}

var RequestFailedTryAgain = Message{
	ID:      "REQUEST_FAILED_TRY_AGAIN",
	Message: "Request was not processed, try again",
}

var RequestFailedTryAgainSupport = Message{
	ID:      "REQUEST_FAILED_TRY_AGAIN_SUPPORT",
	Message: "Request uncompleted, try again and if you keep viewing this message, look for support",
}

var RequestFailedTryAgainValidate = Message{
	ID:      "REQUEST_FAILED_TRY_AGAIN_VALIDATE",
	Message: "Request uncompleted, check the information and try again",
}

var RequestFailedAlreadyExists = Message{
	ID:      "REQUEST_FAILED_ALREADY_EXISTS",
	Message: "Request uncompleted, there is already an item with the same information",
}

var RequestFailedExternalEndpoint = Message{
	ID:      "REQUEST_FAILED_EXTERNAL_ENDPOINT",
	Message: "There was a problem with the request / connection to the external endpoint",
}

var RequestFailedNotFound = Message{
	ID:      "REQUEST_FAILED_ID_NOT_FOUND",
	Message: "There was a problem with the request / connection to the external endpoint",
}

var GeneralInternalError = Message{
	ID:      "GENERAL_INTERNAL_SERVER_ERROR",
	Message: "There was a general problem with the request",
}

var FrameworkInternalError = Message{
	ID:      "FRAMEWORK_INTERNAL_SERVER_ERROR",
	Message: "There was a general problem with the request in framework",
}

var FileFailedBadExtension = Message{
	ID:      "FILE_FAILED_BAD_EXTENSION",
	Message: "There was a problem with the request, file does not have a valid extension",
}


