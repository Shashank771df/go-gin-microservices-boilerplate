package environments

import (
	"app/core/errors"
	"app/core/lang"
	"app/core/utils"
	"app/microservices/base_/config"
)

func (o Handler) ItemGuard() ItemGuard {
	WhiteList := config.Env.Micro.API.WhiteList

	return ItemGuard{
		WhiteList: &WhiteList,
	}
}

func (o Handler) UpdateGuard(data interface{}) Update {
	var params ItemGuard
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

	before := o.ItemGuard()
	var util utils.Util
	util.SetValue(&config.Env.Micro.API.WhiteList, params.WhiteList)

	return Update{
		Before:  before,
		Current: o.ItemGuard(),
	}
}
