package main

import (
	"flag"
	"io/ioutil"
	"os"
	"os/signal"

	"github.com/alex19861108/burj/burj_center/general/config"
	"github.com/alex19861108/burj/burj_center/general/repository"
	"github.com/alex19861108/burj/burj_center/general/server"
	"gopkg.in/yaml.v2"
)

var (
	pConfigFile = flag.String("c", "config.yaml", "yaml config path")
)

func init() {
	flag.Parse()

	// init config
	content, _ := ioutil.ReadFile(*pConfigFile)
	cfg := config.Config{}
	yaml.Unmarshal(content, &cfg)

	// signal interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		os.Exit(1)
	}()

	//zk_server.InitZkServer(cfg)

	repository.InitDB(cfg)
	server.InitHttpServer(cfg)
}

func main() {

}
