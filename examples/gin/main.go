package main

import (
	_ "github.com/chinlee1523/go-admin/adapter/gin"
	"github.com/chinlee1523/go-admin/engine"
	"github.com/chinlee1523/go-admin/examples/datamodel"
	"github.com/chinlee1523/go-admin/modules/config"
	"github.com/chinlee1523/go-admin/modules/db"
	"github.com/chinlee1523/go-admin/modules/language"
	"github.com/chinlee1523/go-admin/plugins/admin"
	"github.com/chinlee1523/go-admin/plugins/example"
	"github.com/chinlee1523/go-admin/template/adminlte"
	"github.com/chinlee1523/go-admin/template/types"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

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

				//DRIVER: db.DriverSqlite,
				//FILE:   "../datamodel/admin.db",
			},
		},
		DOMAIN: "localhost",
		PREFIX: "admin",
		STORE: config.Store{
			PATH:   "./uploads",
			PREFIX: "uploads",
		},
		LANGUAGE:    language.CN,
		INDEX:       "/",
		DEBUG:       true,
		COLORSCHEME: adminlte.COLORSCHEME_SKIN_BLACK,
	}

	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// you can custom a plugin like:

	examplePlugin := example.NewExample()

	if err := eng.AddConfig(cfg).AddPlugins(adminPlugin, examplePlugin).Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	// you can custom your pages like:

	r.GET("/"+cfg.PREFIX+"/custom", func(ctx *gin.Context) {
		engine.Content(ctx, func() types.Panel {
			return datamodel.GetContent()
		})
	})

	_ = r.Run(":9033")
}
