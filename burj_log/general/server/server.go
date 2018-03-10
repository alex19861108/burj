package server

import (
	"net/http"
	"strings"

	"github.com/alex19861108/burj/burj_center/general/common/net"
	"github.com/alex19861108/burj/burj_center/iris/config"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome To Burj Log!"))
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	trimedPath := strings.Trim(r.URL.Path, "/")
	paths := strings.Split(trimedPath, "/")
	if paths[0] != "log" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	r.ParseForm()
	net.WriteJSON(w, http.StatusOK, r.Form)
}

func InitHttpServer(cfg config.ServerConfig) {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/log", logHandler)

	err := http.ListenAndServe(":"+cfg.Port.Http, nil)
	if err != nil {
		panic(err)
	}
}
