package lib

import (
	base "app/core/sdk.base"
	c "app/microservices/user-service/config"
)

var SDKBase base.BaseSDK

func init() {
	SDKBase = base.New(
		c.Env.Micro.APIServices.Base.Host,
		c.Env.Micro.App.APIKey,
	)
}
