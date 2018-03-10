package main

import (
	"flag"
	"io/ioutil"

	"github.com/alex19861108/burj/burj_center/iris/server"

	"github.com/alex19861108/burj/burj_center/iris/common/zk_server"

	"github.com/alex19861108/burj/burj_center/iris/common/database"

	"github.com/alex19861108/burj/burj_center/iris/config"
	"github.com/alex19861108/burj/burj_center/iris/repositories"
	"github.com/alex19861108/burj/burj_center/iris/services"
	"github.com/alex19861108/burj/burj_center/iris/web/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"gopkg.in/yaml.v2"
)

var (
	pConfigFile = flag.String("c", "config.yaml", "yaml config path")
	cfg         = config.Config{}
)

func init() {
	flag.Parse()

	content, _ := ioutil.ReadFile(*pConfigFile)
	yaml.Unmarshal(content, &cfg)
}

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	// Load the template files.
	app.RegisterView(iris.HTML("./web/views", ".html"))
	//app.StaticWeb("/public", "../public")

	// init db
	database.InitMgoDB(cfg)
	repo := repositories.NewNodeRepository()
	nodeService := services.NewNodeService(repo)
	// init zk
	go zk_server.InitZkServer(cfg.BurjCenterConfig, nodeService)
	// init rpc server
	go server.InitRpcServer(cfg, nodeService)

	// Serve our controllers.
	mvc.New(app.Party("/")).Handle(new(controllers.HomeController))
	mvc.Configure(app.Party("/node"), nodes)
	mvc.Configure(app.Party("/device"), devices)

	app.Run(
		// Start the web server at localhost:8080
		iris.Addr(":"+cfg.ServerConfig.Port.Http),
		// disables updates
		iris.WithoutVersionChecker,
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}

// note the mvc.Application, it's not iris.Application.
func nodes(app *mvc.Application) {
	// Create our movie repository with some (memory) data from the database.
	repo := repositories.NewNodeRepository()
	// Create our movie service, we will bind it to the movie app's dependencies.
	nodeService := services.NewNodeService(repo)
	app.Register(cfg)
	app.Register(nodeService)

	// serve our movies controller.
	// Note that you can serve more than one controller
	// you can also create child mvc apps using the `movies.Party(relativePath)` or `movies.Clone(app.Party(...))`
	// if you want.
	app.Handle(new(controllers.NodeController))
}

func devices(app *mvc.Application) {
	repo := repositories.NewDeviceRepository()
	deviceService := services.NewDeviceService(repo)
	app.Register(deviceService)
	app.Handle(new(controllers.DeviceController))
}
