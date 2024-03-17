package lib

import (
	"app/core"
	securitysdk "app/core/security.sdk"
	"app/microservices/user-service/config"
	"errors"
)

var Guard securitysdk.Security

func init() {
	Guard = securitysdk.Security{}
}

// APIKey valida el apikey de las consultas a la aplicacion
func APIKey(c core.AppContext, check bool) (interface{}, error) {
	ret := Guard.APIKey(c, securitysdk.CheckAPIKeyIn{
		APIKey: config.Env.Micro.App.APIKey,
		Check:  check,
	})

	if ret.IsOk {
		return ret.Data, nil
	}

	return ret.Data, errors.New(ret.Error)
}

// WhiteList valida las ips que pueden consultar al endpoint
func WhiteList(c core.AppContext, check bool) (interface{}, error) {
	ret := Guard.WhiteList(c, securitysdk.WhiteListIn{
		WhiteList: config.Env.Micro.API.WhiteList,
		Check:     check,
	})

	if ret.IsOk {
		return ret.Data, nil
	}

	return ret.Data, errors.New(ret.Error)
}

// CheckUserSession valida la sesion de un usuario backoffice
func CheckUserSession(c core.AppContext, check bool) (interface{}, error) {
	ret := Guard.CheckUserSession(c, securitysdk.CheckUserSessionIn{
		SecurityUrl: config.Env.Micro.APIServices.SecurityAdmin.Host,
		Check:       check,
	})

	if ret.IsOk {
		return ret.Data, nil
	}

	return ret.Data, errors.New(ret.Error)
}

// CheckCustomerSession valida la sesion de un user frontoffice
func CheckCustomerSession(c core.AppContext, trace string, check bool) (interface{}, error) {
	ret := Guard.CheckUserSession(c, securitysdk.CheckUserSessionIn{
		SecurityUrl: config.Env.Micro.APIServices.SecurityCustomer.Host,
		Trace:       trace,
		Check:       check,
	})

	if ret.IsOk {
		return ret.Data, nil
	}

	return ret.Data, errors.New(ret.Error)
}
