//
// Author: leafsoar
// Date: 2016-01-05 09:27:52
//

// 每个版本的变化记录

package channel

import (
	"os"

	"cocosupdate/util"
	"cocosupdate/version"
)

// Change 版本变化
type Change struct {
	name   string          // 渠道名称
	chname string          // 变化名称
	vsrc   version.Version // 源版本
	vtar   version.Version // 目标版本
}

// NewChange 创建一个版本变化
func NewChange(name string, vsrc, vtar version.Version) Change {
	return Change{
		name:   name,
		vsrc:   vsrc,
		vtar:   vtar,
		chname: vsrc.Name + "-" + vtar.Name,
	}
}

// ArchiveZip 打包资源
func (c *Change) ArchiveZip(pubpath string) string {
	util.CheckOrCreateDir(pubpath + "/" + c.name)
	temppath, delpath := c.moveToTemp(pubpath)
	zippath := pubpath + "/" + c.name + "/" + c.chname + ".zip"
	util.ArchiveZip(zippath, temppath)
	// 删除临时目录
	_ = os.RemoveAll(delpath)
	return c.chname + ".zip"
}

// MoveToTemp 移动到临时目录
func (c *Change) moveToTemp(pubpath string) (string, string) {
	temppath := pubpath + "/" + c.name + "_temp/" + c.chname
	util.CheckOrCreateDir(temppath)

	// 获取变化的文件
	files := c.vtar.CompareFilter(&c.vsrc)
	for _, name := range files {
		src := c.vtar.Path + "/" + name

		dst := temppath + "/" + name
		util.CopyFile(src, dst)
	}
	delpath := pubpath + "/" + c.name + "_temp"
	return temppath, delpath
}
