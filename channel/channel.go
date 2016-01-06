//
// Author: leafsoar
// Date: 2015-11-02 09:56:14
//

package channel

import (
	"fmt"
	"io/ioutil"

	"github.com/leafsoar/cocosupdate/manifest"
	"github.com/leafsoar/cocosupdate/util"
)

// Channel 渠道相关数据
type Channel struct {
	name     string    // 渠道名称
	path     string    // 路径
	pubpath  string    // 发布路径
	versions []Version // 每个渠道有多个版本的资源
}

// NewChannel 建立一个新渠道
func NewChannel(name string, path string, pubpath string) *Channel {
	return &Channel{
		name:    name,
		path:    path,
		pubpath: pubpath,
	}
}

// InitVersions 初始化版本 返回 引擎版本，如果有
func (c *Channel) InitVersions() string {
	paths := util.GetSubPaths(c.path)
	// 添加版本
	for _, name := range paths {
		c.versions = append(c.versions, NewVersion(name, c.path+"/"+name))
	}
	// fmt.Println(c.versions)
	if len(c.versions) > 0 {
		return c.versions[0].GetEngineVersion()
	}
	return ""
}

// Publish 发布资源
func (c *Channel) Publish(host, engine string) {
	if len(c.versions) <= 1 {
		fmt.Println("没有要发布的资源")
		return
	}
	channelpath := c.pubpath + "/" + c.name
	util.CheckOrCreateDir(channelpath)

	src := c.versions[0]

	mf := manifest.NewManifest()
	// 基本设置
	mf.SetURL(host + "/" + c.name)
	mf.SetVersion(src.name)
	mf.SetEngineVersion(engine)

	// 设置源版本
	vsrc := c.versions[0]
	// 变化从 1 索引开始，0 对比 1
	for i := 1; i < len(c.versions); i++ {
		// 目标版本
		vtar := c.versions[i]
		chg := NewChange(c.name, vsrc, vtar)
		zipfile := chg.ArchiveZip(c.pubpath)

		mf.AddGroupVersion(vtar.name)
		md5, _ := util.GetFileMD5(channelpath + "/" + zipfile)
		mf.AddAsset(zipfile, md5)

		vsrc = vtar
	}

	path := c.pubpath + "/" + c.name + "/"
	con, _ := mf.MarshalMini()
	ioutil.WriteFile(path+"version.manifest", con, 0644)
	// fmt.Println(string(con))
	con, _ = mf.Marshal()
	ioutil.WriteFile(path+"project.manifest", con, 0644)
}
