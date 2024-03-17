package lib

import (
	base "app/core/sdk.base"
	c "app/microservices/base_/config"
)

var SDKBase base.BaseSDK

func init() {
	SDKBase = base.New(
		c.Env.Micro.APIServices.Base.Host,
		c.Env.Micro.App.APIKey,
	)
}
