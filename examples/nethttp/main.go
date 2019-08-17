package main

import (
	_ "github.com/chinlee1523/go-admin/adapter/http"
	"github.com/chinlee1523/go-admin/engine"
	"github.com/chinlee1523/go-admin/examples/datamodel"
	"github.com/chinlee1523/go-admin/modules/config"
	"github.com/chinlee1523/go-admin/modules/db"
	"github.com/chinlee1523/go-admin/modules/language"
	"github.com/chinlee1523/go-admin/plugins/admin"
	"github.com/chinlee1523/go-admin/plugins/example"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

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
		DOMAIN:   "localhost",
		PREFIX:   "admin",
		INDEX:    "/",
		DEBUG:    true,
		LANGUAGE: language.CN,
	}

	examplePlugin := example.NewExample()

	if err := eng.AddConfig(cfg).AddPlugins(admin.NewAdmin(datamodel.Generators), examplePlugin).
		Use(mux); err != nil {
		panic(err)
	}

	_ = http.ListenAndServe(":9002", mux)
}
