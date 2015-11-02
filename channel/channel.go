//
// Author: leafsoar
// Date: 2015-11-02 09:56:14
//

package channel

import (
	"fmt"

	"github.com/leafsoar/cocosupdate/base"
	"github.com/leafsoar/cocosupdate/manifest"
)

// Channel 渠道相关数据
type Channel struct {
	name     string    // 渠道名称
	path     string    // 路径
	versions []Version // 每个渠道有多个版本的资源
}

// NewChannel 建立一个新渠道
func NewChannel(name string, path string) *Channel {
	return &Channel{
		name: name,
		path: path,
	}
}

// InitVersions 初始化版本
func (c *Channel) InitVersions() {
	paths := base.GetSubPaths(c.path)
	// 添加版本
	for _, name := range paths {
		c.versions = append(c.versions, NewVersion(name, c.path+"/"+name))
	}
	// fmt.Println(c.versions)
}

// Publish 发布资源
func (c *Channel) Publish() {
	if len(c.versions) <= 1 {
		fmt.Println("没有要发布的资源")
		return
	}
	src := c.versions[0]

	mf := manifest.NewManifest()
	// 基本设置
	mf.SetURL("http://localhost:8080/default")
	mf.SetVersion(src.name)
	mf.SetEngineVersion("3.7.1")

	filter := src.items
	for i := 1; i < len(c.versions); i++ {
		tar := c.versions[i]
		// 对比两个版本，进行发布
		// fmt.Println(i)
		items := addGroupAssets(mf, tar, filter)
		filter = filter.Merge(items)
	}

	con, _ := mf.Marshal()
	fmt.Println(string(con))
}

// 添加组资源
func addGroupAssets(mf *manifest.Manifest, tar Version, filter Items) Items {
	mf.AddGroupVersion(tar.name)
	items := tar.items.Filter(filter)
	for _, item := range items {
		mf.AddAsset(item.Name, item.MD5)
	}
	return items
}
