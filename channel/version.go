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
	Name  string // 版本名称
	Path  string // 版本路径
	Items []Item // 资源文件
}

// Item 资源项
type Item struct {
	Name string // 资源路径
	MD5  string // md5 值
}

// 初始化资源列表
func (v *Version) initFiles() {
	fp.Walk(v.Path, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() || f.Name() == ".DS_Store" {
			return nil
		}
		name := strings.Replace(path, v.Path+"/", "", 1)
		md5, _ := base.GetFileMD5(path)
		v.Items = append(v.Items, Item{
			Name: name,
			MD5:  md5,
		})
		return nil
	})
}
