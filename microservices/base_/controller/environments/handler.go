package environments

import (
	"app/core"
	"app/core/errors"
	"app/core/lang"
	"app/microservices/base_/config"
	pipes "app/microservices/base_/pipes/environments"
	"net/http"
)

type Handler struct {
}

func (o Handler) SearchItemInterSVC(c core.AppContext) {
	params, ok := c.Data.(*pipes.SearchItemInterSVC)

	if !ok {
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: "Invalid pipe data",
			},
			Data: map[string]interface{}{
				"data": c.Data,
			},
			HideData: true,
		})
	}

	var ret interface{}

	switch *params.KeyId {
	case config.Etc.EnvKeys.Alert:
		ret = o.ItemAlert()
	case config.Etc.EnvKeys.Database:
		ret = o.ItemDatabase()
	case config.Etc.EnvKeys.Guard:
		ret = o.ItemGuard()
	case config.Etc.EnvKeys.Log:
		ret = o.ItemLog()
	case config.Etc.EnvKeys.SdkBase:
		ret = o.ItemSdkBase()
	default:
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: "Invalid keyId",
			},
			Data: map[string]interface{}{
				"data": c.Data,
			},
			HideData: true,
		})
	}

	c.JSON(http.StatusOK, Item{
		KeyId: *params.KeyId,
		Item:  ret,
	})
}

func (o Handler) SearchItemsInterSVC(c core.AppContext) {
	c.JSON(http.StatusOK, Items{
		Alert:    o.ItemAlert(),
		Database: o.ItemDatabase(),
		Guard:    o.ItemGuard(),
		SdkBase:  o.ItemSdkBase(),
		Log:      o.ItemLog(),
	})
}

func (o Handler) UpdateItemInterSVC(c core.AppContext) {
	params, ok := c.Data.(*pipes.UpdateItemInterSVC)
	if !ok {
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: "Invalid pipe data",
			},
			Data: map[string]interface{}{
				"data": c.Data,
			},
			HideData: true,
		})
	}

	var ret interface{}

	switch *params.KeyId {
	case config.Etc.EnvKeys.Alert:
		ret = o.UpdateAlert(params.Data)
	case config.Etc.EnvKeys.Database:
		ret = o.UpdateDatabase(params.Data)
	case config.Etc.EnvKeys.Guard:
		ret = o.UpdateGuard(params.Data)
	case config.Etc.EnvKeys.Log:
		ret = o.UpdateLog(params.Data)
	case config.Etc.EnvKeys.SdkBase:
		ret = o.UpdateSdkBase(params.Data)
	default:
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: "Invalid keyId",
			},
			Data: map[string]interface{}{
				"data": c.Data,
			},
			HideData: true,
		})
	}

	c.JSON(http.StatusOK, Item{
		KeyId: *params.KeyId,
		Item:  ret,
	})
}
