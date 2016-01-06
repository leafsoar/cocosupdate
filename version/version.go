//
// Author: leafsoar
// Date: 2015-11-02 10:59:00
//

package version

import (
	"io/ioutil"
	"os"
	fp "path/filepath"
	"strings"

	"github.com/leafsoar/cocosupdate/manifest"
	"github.com/leafsoar/cocosupdate/util"
)

// Author: leafsoar
// Date: 2015-11-02 10:58:15
//

// Version 资源版本
type Version struct {
	Name  string // 版本名称
	Path  string // 版本路径
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
func (items Items) filter(filter Items) Items {
	var slice = make(Items, 0)
	for _, item := range items {
		if !filter.isContains(item) {
			slice = append(slice, item)
		}
	}
	return slice
}

// NewVersion 创建一个版本
func NewVersion(name string, path string) Version {
	v := Version{
		Name: name,
		Path: path,
	}
	v.initFiles()
	return v
}

// 初始化资源列表
func (v *Version) initFiles() {
	fp.Walk(v.Path, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() || f.Name() == ".DS_Store" {
			return nil
		}
		name := strings.Replace(path, v.Path+"/", "", 1)
		md5, _ := util.GetFileMD5(path)
		v.items = append(v.items, Item{
			Name: name,
			MD5:  md5,
		})
		return nil
	})
}

// GetEngineVersion 获取引擎版本
func (v *Version) GetEngineVersion() string {
	// 遍历所有文件，查看是否有 project.manifest 文件
	for _, item := range v.items {
		// 如果等于指定文件名
		if strings.EqualFold("project.manifest", fp.Base(item.Name)) {
			filename := v.Path + "/" + item.Name
			f, _ := ioutil.ReadFile(filename)
			mf := manifest.NewManifest()
			mf.Unmarshal(f)
			return mf.GetEngineVersion()
		}
	}
	return ""
}

// CompareFilter 对比忽略文件
func (v *Version) CompareFilter(srcv *Version) Items {
	return v.items.filter(srcv.items)
}
