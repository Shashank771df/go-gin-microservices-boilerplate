package routes

import (
	"app/core"
	c "app/microservices/base_/config"
	"app/microservices/base_/controller/users"
	l "app/microservices/base_/lib"
	pipes "app/microservices/base_/pipes/users"
	"net/http"
)

func init() {
	const subject = "users"

	handler := users.Handler{}
	routes := Group{
		Prefix: "/" + c.AppInfo.Version + "/" + subject,
		Routes: []Route{
			{
				/************************************************************
				* * * * * * * * Admin - BackOffice  * * * * * * * *  * * *
				************************************************************/
				Method: http.MethodPost,
				Path:   "",
				Handler: l.InterceptorAPI.Request(core.InterceptorItems{
					Guard: func(ctx core.AppContext) (interface{}, error) {
						return l.CheckUserSession(ctx, false)
					},
					Pipe: func(ctx core.AppContext) interface{} {
						var pipe pipes.NewItem
						return ctx.BindPipe(&pipe)
					},
					Handler: func(ctx core.AppContext) {
						handler.NewItem(ctx)
					},
				}),
			},
			{
				Method: http.MethodGet,
				Path:   "/:id",
				Handler: l.InterceptorAPI.Request(core.InterceptorItems{
					Guard: func(ctx core.AppContext) (interface{}, error) {
						return l.CheckUserSession(ctx, false)
					},
					Pipe: func(ctx core.AppContext) interface{} {
						var pipe pipes.Item
						return ctx.BindPipe(&pipe)
					},
					Handler: func(ctx core.AppContext) {
						handler.Item(ctx)
					},
				}),
			},
			{
				Method: http.MethodPatch,
				Path:   "/:id",
				Handler: l.InterceptorAPI.Request(core.InterceptorItems{
					Guard: func(ctx core.AppContext) (interface{}, error) {
						return l.CheckUserSession(ctx, false)
					},
					Pipe: func(ctx core.AppContext) interface{} {
						var pipe pipes.UpdateItem
						return ctx.BindPipe(&pipe)
					},
					Handler: func(ctx core.AppContext) {
						handler.UpdateItem(ctx)
					},
				}),
			},
			{
				Method: http.MethodDelete,
				Path:   "/:id",
				Handler: l.InterceptorAPI.Request(core.InterceptorItems{
					Guard: func(ctx core.AppContext) (interface{}, error) {
						return l.CheckUserSession(ctx, false)
					},
					Pipe: func(ctx core.AppContext) interface{} {
						var pipe pipes.RemoveItem
						return ctx.BindPipe(&pipe)
					},
					Handler: func(ctx core.AppContext) {
						handler.RemoveItem(ctx)
					},
				}),
			},
			{
				Method: http.MethodGet,
				Path:   "",
				Handler: l.InterceptorAPI.Request(core.InterceptorItems{
					Guard: func(ctx core.AppContext) (interface{}, error) {
						return l.CheckUserSession(ctx, false)
					},
					Pipe: func(ctx core.AppContext) interface{} {
						var pipe pipes.Items
						return ctx.BindPipe(&pipe)
					},
					Handler: func(ctx core.AppContext) {
						handler.Items(ctx)
					},
				}),
			},
			{
				Method: http.MethodPatch,
				Path:   "/:id/activate",
				Handler: l.InterceptorAPI.Request(core.InterceptorItems{
					Guard: func(ctx core.AppContext) (interface{}, error) {
						return l.CheckUserSession(ctx, false)
					},
					Pipe: func(ctx core.AppContext) interface{} {
						var pipe pipes.Activate
						return ctx.BindPipe(&pipe)
					},
					Handler: func(ctx core.AppContext) {
						handler.ActivateItem(ctx)
					},
				}),
			},
			{
				Method: http.MethodPatch,
				Path:   "/:id/deactivate",
				Handler: l.InterceptorAPI.Request(core.InterceptorItems{
					Guard: func(ctx core.AppContext) (interface{}, error) {
						return l.CheckUserSession(ctx, false)
					},
					Pipe: func(ctx core.AppContext) interface{} {
						var pipe pipes.Deactivate
						return ctx.BindPipe(&pipe)
					},
					Handler: func(ctx core.AppContext) {
						handler.DeactivateItem(ctx)
					},
				}),
			},
			{
				Method: http.MethodGet,
				Path:   "/group-by-month",
				Handler: l.InterceptorAPI.Request(core.InterceptorItems{
					Guard: func(ctx core.AppContext) (interface{}, error) {
						return l.CheckUserSession(ctx, false)
					},
					Handler: func(ctx core.AppContext) {
						handler.UsersGroupByMonth(ctx)
					},
				}),
			},

			/************************************************************
			* * * * * * * * Users - FrontOffice * * * * * * * *  * * *
			************************************************************/

			/************************************************************
			* * * * * * * * Api - Interservices * * * * * * * * *  * * *
			************************************************************/
			{
				Method: http.MethodPost,
				Path:   "/search-item/interservices",
				Handler: l.InterceptorAPI.Request(
					core.InterceptorItems{
						Guard: func(ctx core.AppContext) (interface{}, error) {
							return l.APIKey(ctx, true)
						},
						Pipe: func(ctx core.AppContext) interface{} {
							var pipe pipes.SearchItemInterSVC
							return ctx.BindPipe(&pipe)
						},
						Handler: func(ctx core.AppContext) {
							handler.SearchItemInterSVC(ctx)
						},
					},
				),
			},
		},
	}

	AppRouting = append(AppRouting, routes)
}
