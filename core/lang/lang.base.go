package lang

var Base base = base{
	USER_ALREADY_EXISTS: Message{
		ID:      "USER_ALREADY_EXISTS",
		Message: "User is already taken",
	},
	USER_INFORMATION_IN_USE: Message{
		ID:      "USER_INFORMATION_IN_USE",
		Message: "User information is in use",
	},
	USER_NOT_EXISTS: Message{
		ID:      "USER_NOT_EXISTS",
		Message: "User does not exists",
	},
}

type base struct {
	USER_ALREADY_EXISTS     Message
	USER_INFORMATION_IN_USE Message
	USER_NOT_EXISTS         Message
}
