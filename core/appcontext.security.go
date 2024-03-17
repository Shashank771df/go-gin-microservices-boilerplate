package core

type (
	// SecuritySessionValidate .
	SecuritySessionValidate struct {
		Success bool `json:"success"`
		Info    struct {
			Sessions      int8   `json:"sessions"`
			AliveSessions int8   `json:"aliveSessions"`
			UserActive    int8   `json:"userActive"`
			UserID        uint32 `json:"userId"`
			Token         string `json:"token"`
			ExpDatetime   string `json:"expDatetime"`
		} `json:"info"`
		TokenExpired bool `json:"tokenExpired"`
	}
)
