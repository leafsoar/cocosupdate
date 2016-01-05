//
// Author: leafsoar
// Date: 2015-10-27 16:19:13
//

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/leafsoar/cocosupdate/channel"

	"github.com/codegangsta/cli"
)

// 根据源资源，生成发布资源
func publish(url string, assets string, publish string, name string) {
	fmt.Println("发布地址: ", url)
	fmt.Println("资源目录: ", assets)
	fmt.Println("发布目录: ", publish)
	fmt.Println("发布名称: ", name)
	ch := channel.NewChannel(name, assets, publish)
	ch.InitVersions()
	// 发布
	ch.Publish(url)
	fmt.Printf("发布完成: http://%s/%s/%s", url, publish, name)
}

func startServer(port string) {
	log.Printf("启动 Http 服务 %s ...", port)

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
	// fmt.Println("Cocos 热更新")

	app := cli.NewApp()
	app.Name = "Cocos Update"
	app.Usage = "Cocos 资源热更新部署工具 (配合 AssetsManagerEx 使用)"
	app.Version = "0.3.3"
	app.Author = "leafsoar"
	app.Email = "kltwjt@gmail.com"

	app.Commands = []cli.Command{
		{
			Name:        "build",
			Usage:       "构建更新资源文件",
			Description: "默认发布地址: http://localhost:8000/res",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "address, addr",
					Value: "localhost:8000",
					Usage: "发布资源地址 (host:port)",
				},
				cli.StringFlag{
					Name:  "assets",
					Value: "assets",
					Usage: "资源目录 (包含各个版本号的资源)",
				},
				cli.StringFlag{
					Name:  "publish",
					Value: "publish",
					Usage: "发布目录 (发布资源根目录)",
				},
				cli.StringFlag{
					Name:  "name",
					Value: "res",
					Usage: "发布名称 (区分渠道或者地址)",
				},
			},
			Action: func(c *cli.Context) {
				fmt.Println("开始生成发布资源 ...")
				publish(
					c.String("address"),
					c.String("assets"),
					c.String("publish"),
					c.String("name"),
				)

			},
		},
		{
			Name:  "start",
			Usage: "启动 HTTP 资源服务器",
			Action: func(c *cli.Context) {
				fmt.Println("HTTP 服务器已经开启 ...")
			},
		},
	}

	app.Run(os.Args)

	// port := "8001"
	// url := "http://192.168.1.51:" + port
	// // 发布资源
	// publish(url)
	// // 启动服务
	// startServer(port)
	// // 监听 kill
	// handleSignals()
}
