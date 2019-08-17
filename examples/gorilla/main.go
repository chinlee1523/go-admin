package main

import (
	"log"
	"net/http"

	_ "github.com/chinlee1523/go-admin/adapter/gorilla"

	"github.com/chinlee1523/go-admin/engine"
	"github.com/chinlee1523/go-admin/examples/datamodel"
	"github.com/chinlee1523/go-admin/modules/config"
	"github.com/chinlee1523/go-admin/modules/db"
	"github.com/chinlee1523/go-admin/modules/language"
	"github.com/chinlee1523/go-admin/plugins/admin"
	"github.com/chinlee1523/go-admin/plugins/example"
	"github.com/gorilla/mux"
)

func main() {
	app := mux.NewRouter()
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
		STORE: config.Store{
			PATH:   "./uploads",
			PREFIX: "uploads",
		},
		DOMAIN:   "localhost",
		PREFIX:   "admin",
		INDEX:    "/",
		DEBUG:    true,
		LANGUAGE: language.EN,
	}

	examplePlugin := example.NewExample()

	if err := eng.AddConfig(cfg).AddPlugins(admin.NewAdmin(datamodel.Generators), examplePlugin).
		Use(app); err != nil {
		panic(err)
	}

	log.Println("Listening 8080")
	log.Fatal(http.ListenAndServe(":8080", app))
}
