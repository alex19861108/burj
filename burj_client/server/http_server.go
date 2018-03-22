package server

import (
	"log"
	"net/http"

	"github.com/alex19861108/burj/burj_center/general/common/net"
	"github.com/alex19861108/burj/burj_center/iris/config"
	"github.com/alex19861108/burj/burj_client/service"
)

const (
	TEST_PIC_PKG_PATH = "/Users/megvii/Downloads/vivo-83.zip"
	TEST_APK_PATH     = "/Users/megvii/Downloads/app-debug.apk"
	TEST_APK_PKG_NAME = "com.megvii.qa.quicktestproject12x"
	TEST_APK_CLS_NAME = "MainActivity"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome To Burj Client!"))
}

func runHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		r.FormValue("serial")
	}

	adb := service.ADB{}
	serials := adb.Serials()
	adb.Run(serials[0], TEST_PIC_PKG_PATH, TEST_APK_PATH, TEST_APK_PKG_NAME, TEST_APK_CLS_NAME, "jid", "0")
	net.WriteJSON(w, http.StatusOK, "success")
}

func InitHttpServer(cfg config.Config) {
	// listen and serve
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/run", runHandler)
	err := http.ListenAndServe(":"+cfg.ServerConfig.Port.Http, nil)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
