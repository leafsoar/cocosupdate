//
// Author: leafsoar
// Date: 2015-11-02 10:06:25
//

package base

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	fp "path/filepath"
	"strings"
)

// ResItem 资源项
type ResItem struct {
	Name string
	Path string
}

// ResItems 资源集合
type ResItems []ResItem

// FilterRemove 移除过滤的文件
func (r ResItems) FilterRemove(name string) ResItems {
	ret := make(ResItems, 0)
	for _, item := range r {
		if !strings.Contains(item.Path, name) {
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
			Name: f.Name(),
			Path: path,
		}
		ret = append(ret, res)
		return nil
	})
	return ret
}

// GetFileMD5 获取一个文件的 MD5 值
func GetFileMD5(path string) (string, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return "", err
	}
	md5h := md5.New()
	io.Copy(md5h, file)
	md5v := hex.EncodeToString(md5h.Sum(nil))
	return md5v, nil
}

// GetSubPaths 获取一个目录下的目录，不包括子目录
func GetSubPaths(root string) []string {
	var slice []string
	list, err := ioutil.ReadDir(root)
	if err != nil {
		return slice
	}
	for _, item := range list {
		if item.IsDir() {
			slice = append(slice, item.Name())
		}
	}
	return slice
}
