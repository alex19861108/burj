package main

import (
	"flag"
	"io/ioutil"

	"github.com/alex19861108/burj/burj_center/iris/common/database"
	"github.com/alex19861108/burj/burj_center/iris/config"
	"github.com/alex19861108/burj/burj_server/iris/repositories"
	"github.com/alex19861108/burj/burj_server/iris/services"
	"github.com/alex19861108/burj/burj_server/iris/web/controllers"
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

	// init db
	database.InitMgoDB(cfg)
	database.InitMysqlDB(cfg)

	// Serve our controllers.
	mvc.New(app.Party("/")).Handle(new(controllers.HomeController))
	mvc.Configure(app.Party("/job"), jobs)

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

func jobs(app *mvc.Application) {
	jobRepo := repositories.NewJobRepository()
	jobService := services.NewJobService(jobRepo)
	app.Register(jobService)

	imageRepo := repositories.NewImageRepository()
	imageService := services.NewImageService(imageRepo)
	app.Register(imageService)

	tagRepo := repositories.NewTagRepository()
	tagService := services.NewTagService(tagRepo)
	app.Register(tagService)

	app.Register(cfg)

	app.Handle(new(controllers.JobController))
}
