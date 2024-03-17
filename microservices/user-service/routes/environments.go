package routes

import (
	"app/core"
	c "app/microservices/user-service/config"
	"app/microservices/user-service/controller/environments"
	l "app/microservices/user-service/lib"
	pipes "app/microservices/user-service/pipes/environments"
	"net/http"
)

func init() {
	const subject = "environments"

	handler := environments.Handler{}
	routes := Group{
		Prefix: "/" + c.AppInfo.Version + "/" + subject,
		Routes: []Route{
			/************************************************************
			* * * * * * * * Users - Backoffice * * * * * * * *  * * *
			************************************************************/
			/************************************************************
			* * * * * * * * Api - Interservices * * * * * * * * *  * * *
			************************************************************/
			{
				Method: http.MethodPost,
				Path:   "/search-item/interservices",
				Handler: l.InterceptorAPI.Request(core.InterceptorItems{
					Guard: func(ctx core.AppContext) (interface{}, error) {
						return l.APIKey(ctx, false)
					},
					Pipe: func(ctx core.AppContext) interface{} {
						var pipe pipes.SearchItemInterSVC
						return ctx.BindPipe(&pipe)
					},
					Handler: func(ctx core.AppContext) {
						handler.SearchItemInterSVC(ctx)
					},
				}),
			},
			{
				Method: http.MethodPost,
				Path:   "/update-item/interservices",
				Handler: l.InterceptorAPI.Request(core.InterceptorItems{
					Guard: func(ctx core.AppContext) (interface{}, error) {
						return l.APIKey(ctx, false)
					},
					Pipe: func(ctx core.AppContext) interface{} {
						var pipe pipes.UpdateItemInterSVC
						return ctx.BindPipe(&pipe)
					},
					Handler: func(ctx core.AppContext) {
						handler.UpdateItemInterSVC(ctx)
					},
				}),
			},
			{
				Method: http.MethodPost,
				Path:   "/search-items/interservices",
				Handler: l.InterceptorAPI.Request(core.InterceptorItems{
					Guard: func(ctx core.AppContext) (interface{}, error) {
						return l.APIKey(ctx, false)
					},
					Handler: func(ctx core.AppContext) {
						handler.SearchItemsInterSVC(ctx)
					},
				}),
			},
		},
	}

	AppRouting = append(AppRouting, routes)
}
