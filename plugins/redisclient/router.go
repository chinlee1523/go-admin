package redisclient

import (
	"github.com/chinlee1523/go-admin/context"
	"github.com/chinlee1523/go-admin/modules/auth"
	"github.com/chinlee1523/go-admin/plugins/admin/controller"
)

func InitRouter(prefix string) *context.App {
	app := context.NewApp()

	authenticator := auth.SetPrefix(prefix).SetAuthFailCallback(func(ctx *context.Context) {
		ctx.Write(302, map[string]string{
			"Location": prefix + "/login",
		}, ``)
	}).SetPermissionDenyCallback(func(ctx *context.Context) {
		controller.ShowErrorPage(ctx, "permission denied")
	})

	app.GET(prefix+"/redisclient", authenticator.Middleware(Show))

	return app
}
