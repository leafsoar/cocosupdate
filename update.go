//
// Author: leafsoar
// Date: 2015-10-27 16:19:13
//

package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/leafsoar/cocosupdate/channel"
)

// 根据源资源，生成发布资源
func publish(port string) {
	ch := channel.NewChannel("default", "assets", "publish")
	ch.InitVersions()
	// 发布
	ch.Publish("http://localhost:" + port)
}

func startServer(port string) {
	log.Println("启动 Http 服务 ...")

	http.Handle("/", http.FileServer(http.Dir("publish")))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("Error: " + err.Error())
	}
}

func handleSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}

func main() {
	log.Println("leafsoar v5 ~")

	port := "8000"

	// 发布资源
	publish(port)
	// 启动服务
	startServer(port)
	// 监听 kill
	handleSignals()
}
