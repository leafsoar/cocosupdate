//u
// Author: leafsoar
// Date: 2015-10-27 16:19:13
//

package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
	fp "path/filepath"
	"strings"
)

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
		md5 := GetFileMD5(path)
		v.Items = append(v.Items, Item{
			Name: name,
			MD5:  md5,
		})
		return nil
	})
}

// GetFileMD5 获取一个文件的 MD5 值
func GetFileMD5(path string) string {
	file, err := os.Open(path)
	defer file.Close()
	checkErr(err)
	md5h := md5.New()
	io.Copy(md5h, file)
	md5v := hex.EncodeToString(md5h.Sum(nil))
	return md5v
}

// GetVersions 获取一个目录的子目录，第一级
func GetVersions(root string) []Version {
	var slice []Version
	fp.Walk(root, func(path string, f os.FileInfo, err error) error {
		val := strings.Replace(path, root+"/", "", 1)
		// 如果不是目录直接返回
		if !f.IsDir() ||
			strings.EqualFold(val, root) ||
			!strings.EqualFold(val, f.Name()) {
			return nil
		}
		// log.Println(path)
		version := Version{
			Name: f.Name(),
			Path: path,
		}
		version.initFiles()
		slice = append(slice, version)
		return nil
	})
	return slice
}

func main() {
	// log.Println("leafsoar v5 ~")

	// 获取所有的版本
	// list := GetVersions("assets")
	// log.Println(list)
	TestJSON()
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
