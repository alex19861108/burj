package main

import (
	"flag"
	"io/ioutil"
	"os"
	"os/signal"

	"github.com/alex19861108/burj/burj_center/iris/common/zk_client"
	"github.com/alex19861108/burj/burj_center/iris/config"
	"github.com/alex19861108/burj/burj_client/server"
	"gopkg.in/yaml.v2"
)

var (
	pConfigFile = flag.String("c", "config.yaml", "yaml config path")
)

func init() {
	flag.Parse()

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

	go zk_client.InitZkClient(cfg)
	go server.InitAdbServer(cfg)
	go server.InitRpcServer(cfg)

	server.InitHttpServer(cfg)
}

func main() {
}
