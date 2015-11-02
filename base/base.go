//
// Author: leafsoar
// Date: 2015-11-02 10:06:25
//

package base

import (
	"os"
	fp "path/filepath"
	"strings"
)

// ResItem 资源项
type ResItem struct {
	name string
	path string
}

// ResItems 资源集合
type ResItems []ResItem

// FilterRemove 移除过滤的文件
func (r ResItems) FilterRemove(name string) ResItems {
	ret := make(ResItems, 0)
	for _, item := range r {
		if !strings.Contains(item.path, name) {
			ret = append(ret, item)
		}
	}
	return ret
}

// GetFiles 获取一个目录下的所有文件
func GetFiles(root string) ResItems {
	ret := make(ResItems, 0)
	fp.Walk(root, func(path string, f os.FileInfo, err error) error {
		if f == nil || f.IsDir() {
			return nil
		}
		res := ResItem{
			name: f.Name(),
			path: path,
		}
		ret = append(ret, res)
		return nil
	})
	return ret
}
