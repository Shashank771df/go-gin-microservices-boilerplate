package securitysdk

import (
	"app/core"
	"app/core/network"
	"app/core/utils"
	"net/http"
)

type (
	Security struct{}
)

// APIKey valida el apikey de las consultas a la aplicacion
func (g Security) APIKey(c core.AppContext, in CheckAPIKeyIn) GuardOut {
	if !in.Check {
		return GuardOut{IsOk: true}
	}

	headerAPIKey := c.Request.Header.Get("apiKey")
	if headerAPIKey == "" {
		return GuardOut{IsOk: false, Error: "Empty API Key"}
	}

	if headerAPIKey != in.APIKey {
		return GuardOut{IsOk: false, Error: "Invalid API Key"}
	}

	return GuardOut{IsOk: true}
}

func (g Security) WhiteList(c core.AppContext, in WhiteListIn) GuardOut {
	if !in.Check {
		return GuardOut{IsOk: true}
	}

	WhiteList := in.WhiteList

	if len(WhiteList) == 0 || WhiteList[0] == "" {
		return GuardOut{IsOk: true}
	}

	remoteIP := network.UtilNetwork{}.RemoteIPAddress(c)
	founded := false
	var err string = "IP not found"

	for _, ip := range WhiteList {
		if remoteIP == ip {
			founded = true
			err = ""
		}
	}

	return GuardOut{IsOk: founded, Error: err, Data: remoteIP}
}

// UsrSession valida la session del cliente
func (g Security) CheckUserSession(c core.AppContext, in CheckUserSessionIn) GuardOut {
	// captura de datos
	headerUsrID := c.Request.Header.Get("userId")
	headerUsrToken := c.Request.Header.Get("userToken")

	var ok bool
	userId := utils.UtilString{Value: headerUsrID}.ToUint32(&ok)

	if !in.Check {
		return GuardOut{IsOk: true, Data: map[string]interface{}{
			"success": true,
			"info": map[string]interface{}{
				"userId": userId,
				"token":  headerUsrToken,
			},
			"userActive": 1,
		}}
	}

	if headerUsrID == "" {
		return GuardOut{IsOk: false, Error: "userId is not detected"}
	}

	if headerUsrToken == "" {
		return GuardOut{IsOk: false, Error: "userToken is not detected"}
	}

	if in.SecurityUrl == "" {
		return GuardOut{IsOk: false, Error: "Empty securityURL"}
	}

	//+
	client := network.HttpClient{}

	result := client.Call(network.HttpClientRequest{
		Method: "POST",
		Url:    in.SecurityUrl + "/session/validate",
		Headers: map[string]string{
			"userId":    headerUsrID,
			"userToken": headerUsrToken,
		},
		Data: map[string]string{
			"trace": in.Trace,
		},
	})

	if result.Error != nil {
		return GuardOut{IsOk: false, Error: result.Error.Error(), Data: result}
	}

	if result.Status == http.StatusUnauthorized || result.Status == http.StatusBadRequest {
		return GuardOut{IsOk: false, Error: "Unauthorized", Data: result}
	}

	return GuardOut{IsOk: true, Data: result.Data}
}

// UsrSession valida la session del cliente
func (g Security) CheckUserIsAuthorized(c core.AppContext, in CheckUserIsAuthorizedIn) GuardOut {
	if !in.Check {
		return GuardOut{IsOk: true}
	}

	// captura de datos
	headerUsrID := c.Request.Header.Get("userId")
	if headerUsrID == "" {
		return GuardOut{IsOk: false, Error: "userId is not detected"}
	}
	headerUsrToken := c.Request.Header.Get("userToken")
	if headerUsrToken == "" {
		return GuardOut{IsOk: false, Error: "userToken is not detected"}
	}

	if in.SecurityUrl == "" {
		return GuardOut{IsOk: false, Error: "Empty securityURL"}
	}

	//+
	client := network.HttpClient{}

	result := client.Call(network.HttpClientRequest{
		Method: "POST",
		Url:    in.SecurityUrl,
		Data: map[string]string{
			"userId":     headerUsrID,
			"token":      headerUsrToken,
			"endpointId": in.EndpointId,
		},
	})

	if !result.IsStatusIn2xx {
		return GuardOut{IsOk: false, Error: "Unauthorized", Data: result}
	}

	return GuardOut{IsOk: true, Data: result.Data}
}

// UsrSession valida la session del cliente
func (g Security) CheckUserPrivilege(c core.AppContext, in CheckUserPrivilegeIn) GuardOut {
	if !in.Check {
		return GuardOut{IsOk: true}
	}

	// captura de datos
	headerUsrID := c.Request.Header.Get("userId")
	if headerUsrID == "" {
		return GuardOut{IsOk: false, Error: "userId is not detected"}
	}
	headerUsrToken := c.Request.Header.Get("userToken")
	if headerUsrToken == "" {
		return GuardOut{IsOk: false, Error: "userToken is not detected"}
	}

	if in.SecurityUrl == "" {
		return GuardOut{IsOk: false, Error: "Empty securityURL"}
	}

	//+
	client := network.HttpClient{}

	result := client.Call(network.HttpClientRequest{
		Method: "POST",
		Url:    in.SecurityUrl,
		Headers: map[string]string{
			"userId":    headerUsrID,
			"userToken": headerUsrToken,
		},
		Data: map[string]string{
			"endpointId": in.EndpointId,
			"trace":      in.Trace,
		},
	})

	if !result.IsStatusIn2xx {
		return GuardOut{IsOk: false, Error: "Unauthorized", Data: result}
	}

	return GuardOut{IsOk: true, Data: result.Data}
}
