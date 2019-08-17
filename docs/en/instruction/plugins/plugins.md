# Plugins Usage

To use the plugin, just call the `AddPlugins` method of the engine.

Example:

```go
package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/chinlee1523/go-admin/adapter/gin" // adapter must be imported, if not - you have to implement it yourself
	"github.com/chinlee1523/go-admin/engine"
	"github.com/chinlee1523/go-admin/plugins/admin"
	"github.com/chinlee1523/go-admin/plugins/example"
	"github.com/chinlee1523/go-admin/modules/config"
	"github.com/chinlee1523/go-admin/examples/datamodel"
)

func main() {
	r := gin.Default()
	eng := engine.Default()
	cfg := config.Config{}

	adminPlugin := admin.NewAdmin(datamodel.Generators)
	examplePlugin := example.NewExample()

	eng.AddConfig(cfg).
		AddPlugins(adminPlugin, examplePlugin). // Load plugin
		Use(r)

	r.Run(":9033")
}
```

[Back to Contents](https://github.com/chinlee1523/go-admin/blob/master/docs/en/index.md)<br>
[Previous:go-admin introduction](https://github.com/chinlee1523/go-admin/blob/master/docs/en/instruction/instruction.md)<br>
[Next: Admin plugin](https://github.com/chinlee1523/go-admin/blob/master/docs/en/instruction/plugins/admin.md)
