//
// Author: leafsoar
// Date: 2015-10-27 16:19:13
//

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"cocosupdate/channel"

	"github.com/codegangsta/cli"
)

// 根据源资源，生成发布资源
func publish(address, assets, publish, name, engine string) {
	fmt.Println("开始生成发布资源 ...")
	fmt.Println("资源目录: ", assets)
	fmt.Println("发布目录: ", publish)
	fmt.Println("发布地址: ", address)
	fmt.Println("发布名称: ", name)
	ch := channel.NewChannel(name, assets, publish)
	pubengine := ch.InitVersions()
	// 如果传入引擎版本不等于 3.7.1，使用强制赋值
	if !strings.EqualFold(engine, "3.7.1") {
		pubengine = engine
	}
	// 发布
	ch.Publish("http://"+address, pubengine)
	fmt.Println("引擎版本: ", pubengine)
	fmt.Printf("发布完成:  http://%s/%s\n", address, name)
}

func startServer(port string, publish string) {
	fmt.Printf("启动 Http 服务 path: %s port: %s...", publish, port)

	http.Handle("/", http.FileServer(http.Dir(publish)))
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
	app := cli.NewApp()
	app.Name = "Cocos Update"
	app.Usage = "Cocos 资源热更新部署工具 (请配合 AssetsManagerEx 使用)"
	app.Version = "0.3.3"
	app.Author = "leafsoar"
	app.Email = "kltwjt@gmail.com"

	app.Commands = []cli.Command{
		{
			Name:        "build",
			Usage:       "构建热更新资源文件",
			Description: "默认发布地址: http://localhost:8000/res",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "address, addr",
					Value: "localhost:8000",
					Usage: "发布资源地址 (ip:port)",
				},
				cli.StringFlag{
					Name:  "assets, a",
					Value: "assets",
					Usage: "资源目录 (包含各个版本号的资源)",
				},
				cli.StringFlag{
					Name:  "publish, p",
					Value: "publish",
					Usage: "发布目录 (发布资源根目录)",
				},
				cli.StringFlag{
					Name:  "name, n",
					Value: "res",
					Usage: "发布名称 (区分渠道或者地址)",
				},
				cli.StringFlag{
					Name:  "engine, e",
					Value: "3.7.1",
					Usage: "cocos 引擎版本 (不指定时自动读取 project.manifest 值 || 3.7.1)",
				},
			},
			Action: func(c *cli.Context) {
				publish(
					c.String("address"),
					c.String("assets"),
					c.String("publish"),
					c.String("name"),
					c.String("engine"),
				)
			},
		},
		{
			Name:  "start",
			Usage: "启动 HTTP 资源服务器",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "port, p",
					Value: "8000",
					Usage: "热更新端口号",
				},
				cli.StringFlag{
					Name:  "publish, pub",
					Value: "publish",
					Usage: "热更新资源目录",
				},
			},
			Action: func(c *cli.Context) {
				startServer(c.String("port"), c.String("publish"))
				handleSignals()
			},
		},
	}
	app.Run(os.Args)
}
