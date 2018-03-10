package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/alex19861108/burj/burj_center/general/service"

	"github.com/alex19861108/burj/burj_center/general/common/net"
	"github.com/alex19861108/burj/burj_center/general/config"
	"github.com/alex19861108/burj/burj_center/general/model"
	"github.com/samuel/go-zookeeper/zk"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome To Burj Server!"))
}

type NodeHandler struct {
	Cfg config.BurjCenterConfig
}

const (
	timeout = 10
)

func (h *NodeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	conn, _, err := zk.Connect([]string{h.Cfg.HOST}, time.Duration(timeout)*time.Second)
	if err != nil {
		panic(err)
		return
	}
	defer conn.Close()
	children, _, err := conn.Children(h.Cfg.ROOT)
	net.WriteJSON(w, http.StatusOK, children)
}

func registerNodeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	node := &model.Node{}
	err := json.Unmarshal([]byte(r.Form.Get("node")), &node)
	if err != nil {
		net.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	err = service.RegisterNode(node)
	if err != nil {
		net.WriteJSON(w, http.StatusBadRequest, err)
	} else {
		net.WriteJSON(w, http.StatusOK, node)
	}
	return
}

func listNodeHandler(w http.ResponseWriter, r *http.Request) {
	net.WriteJSON(w, http.StatusOK, r)
}

func InitHttpServer(cfg config.Config) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)
	mux.Handle("/node", &NodeHandler{Cfg: cfg.BurjCenterConfig})
	mux.HandleFunc("/node/list", listNodeHandler)
	mux.HandleFunc("/node/list", listNodeHandler)
	mux.HandleFunc("/node/register", registerNodeHandler)

	err := http.ListenAndServe(":"+cfg.ServerConfig.Port.HTTP, mux)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
