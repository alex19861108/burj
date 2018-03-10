package server

import (
	"log"
	"reflect"
	"sort"
	"time"

	"github.com/alex19861108/burj/burj_center/iris/common/zk_client"
	"github.com/alex19861108/burj/burj_center/iris/config"
	"github.com/alex19861108/burj/burj_client/service"
)

type AdbServer struct {
	adb     service.ADB
	serials []string
	cfg     config.Config
}

func (d *AdbServer) serve() {
	for {
		serials := d.adb.Serials()
		sort.Strings(serials)
		sort.Strings(d.serials)
		if !reflect.DeepEqual(serials, d.serials) {
			log.Println("device changed.")
			devices, err := d.adb.Devices()
			if err == nil {
				node, err := zk_client.NodeGet(d.cfg)
				if err != nil {
					panic(err)
				}
				node.Devices = devices
				if err = zk_client.NodeSet(d.cfg, node); err != nil {
					panic(err)
				}
				d.serials = serials
			}
		}
		time.Sleep(time.Second)
	}
}

func InitAdbServer(cfg config.Config) {
	a := service.ADB{}
	m := AdbServer{adb: a, cfg: cfg}
	m.serve()
}
