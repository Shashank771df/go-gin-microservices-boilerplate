package routes

import (
	"app/core"
	c "app/microservices/base_/config"
	"app/microservices/base_/controller/common"
	l "app/microservices/base_/lib"
	"net/http"
)

func init() {
	const subject = "common"

	group := Group{
		Prefix: "/" + c.AppInfo.Version + "/" + subject,
		Routes: []Route{
			{
				Method: http.MethodGet,
				Path:   "/health",
				Handler: l.InterceptorAPI.Request(
					core.InterceptorItems{
						Handler: func(ctx core.AppContext) {
							common.Handler{}.Health(ctx)
						},
					},
				),
			},
			{
				Method: http.MethodPost,
				Path:   "/sync",
				Handler: l.InterceptorAPI.Request(
					core.InterceptorItems{
						Handler: func(ctx core.AppContext) {
							common.Handler{}.Sync(ctx)
						},
					},
				),
			},
			{
				Method: http.MethodGet,
				Path:   "/sync/interservices",
				Handler: l.InterceptorAPI.Request(
					core.InterceptorItems{
						Guard: func(ctx core.AppContext) (interface{}, error) {
							return l.APIKey(ctx, true)
						},
						Handler: func(ctx core.AppContext) {
							common.Handler{}.Sync(ctx)
						},
					},
				),
			},
		},
	}
	AppRouting = append(AppRouting, group)
}
