//
// Author: leafsoar
// Date: 2015-11-02 10:59:00
//

package channel

import (
	"os"
	fp "path/filepath"
	"strings"

	"github.com/leafsoar/cocosupdate/base"
)

// Author: leafsoar
// Date: 2015-11-02 10:58:15
//

// Version 资源版本
type Version struct {
	name  string // 版本名称
	path  string // 版本路径
	items Items  // 资源文件
}

// Item 资源项
type Item struct {
	Name string // 资源路径
	MD5  string // md5 值
}

// Items 项集合
type Items []Item

func (items Items) isContains(item Item) bool {
	for _, sitem := range items {
		if sitem.Name == item.Name &&
			sitem.MD5 == item.MD5 {
			return true
		}
	}
	return false
}

// Filter 过滤
func (items Items) Filter(filter Items) Items {
	var slice = make(Items, 0)
	for _, item := range items {
		if !filter.isContains(item) {
			slice = append(slice, item)
		}
	}
	return slice
}

// Merge 合并
func (items Items) Merge(app Items) Items {
	var slice = make(Items, len(items)+len(app))
	copy(slice, items)
	copy(slice[len(items):], app)
	return slice
}

// NewVersion 创建一个版本
func NewVersion(name string, path string) Version {
	v := Version{
		name: name,
		path: path,
	}
	v.initFiles()
	return v
}

// 初始化资源列表
func (v *Version) initFiles() {
	fp.Walk(v.path, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() || f.Name() == ".DS_Store" {
			return nil
		}
		name := strings.Replace(path, v.path+"/", "", 1)
		md5, _ := base.GetFileMD5(path)
		v.items = append(v.items, Item{
			Name: name,
			MD5:  md5,
		})
		return nil
	})
}
