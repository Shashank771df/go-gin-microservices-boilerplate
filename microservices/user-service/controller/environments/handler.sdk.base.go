package environments

import (
	"app/core/errors"
	"app/core/lang"
	"app/core/utils"
	"app/microservices/user-service/config"
	"app/microservices/user-service/lib"
)

func (o Handler) ItemSdkBase() ItemSdkBase {
	Host := config.Env.Micro.APIServices.Base.Host

	return ItemSdkBase{
		Host: &Host,
	}
}

func (o Handler) UpdateSdkBase(data interface{}) Update {
	var params ItemSdkBase
	err := utils.UtilsMap{}.InterfaceToStruct(data, &params)

	if err != nil {
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: "Invalid pipe data",
			},
			Data: map[string]interface{}{
				"data":  data,
				"model": ItemAlert{},
			},
			HideData: true,
		})
	}

	before := o.ItemSdkBase()
	var util utils.Util
	util.SetValue(&config.Env.Micro.APIServices.Base.Host, params.Host)
	lib.SDKBase.SetHost(config.Env.Micro.APIServices.Base.Host)

	return Update{
		Before:  before,
		Current: o.ItemSdkBase(),
	}
}
