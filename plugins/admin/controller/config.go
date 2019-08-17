package controller

import "github.com/chinlee1523/go-admin/modules/config"

var Config config.Config

func SetConfig(cfg config.Config) {
	Config = cfg
}
