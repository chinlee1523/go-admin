package main

import (
	_ "github.com/chinlee1523/go-admin/adapter/buffalo"
	"github.com/chinlee1523/go-admin/engine"
	"github.com/chinlee1523/go-admin/examples/datamodel"
	"github.com/chinlee1523/go-admin/modules/config"
	"github.com/chinlee1523/go-admin/modules/db"
	"github.com/chinlee1523/go-admin/modules/language"
	"github.com/chinlee1523/go-admin/plugins/admin"
	"github.com/chinlee1523/go-admin/plugins/example"
	"github.com/chinlee1523/go-admin/template/adminlte"
	"github.com/chinlee1523/go-admin/template/types"
	"github.com/gobuffalo/buffalo"
	"net/http"
)

func main() {
	bu := buffalo.New(buffalo.Options{
		Env:  "test",
		Addr: "127.0.0.1:9033",
	})

	eng := engine.Default()

	cfg := config.Config{
		DATABASE: []config.Database{
			{
				HOST:         "127.0.0.1",
				PORT:         "3306",
				USER:         "root",
				PWD:          "root",
				NAME:         "godmin",
				MAX_IDLE_CON: 50,
				MAX_OPEN_CON: 150,
				DRIVER:       db.DriverMysql,
			},
		},
		DOMAIN: "localhost",
		PREFIX: "admin",
		STORE: config.Store{
			PATH:   "./uploads",
			PREFIX: "uploads",
		},
		LANGUAGE:    language.EN,
		INDEX:       "/",
		DEBUG:       true,
		COLORSCHEME: adminlte.COLORSCHEME_SKIN_BLACK,
	}

	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// you can custom a plugin like:

	examplePlugin := example.NewExample()

	if err := eng.AddConfig(cfg).AddPlugins(adminPlugin, examplePlugin).Use(bu); err != nil {
		panic(err)
	}

	bu.ServeFiles("/uploads", http.Dir("./uploads"))

	// you can custom your pages like:

	bu.GET("/"+cfg.PREFIX+"/custom", func(ctx buffalo.Context) error {
		engine.Content(ctx, func() types.Panel {
			return datamodel.GetContent()
		})
		return nil
	})

	_ = bu.Serve()
}
