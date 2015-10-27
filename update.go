//u
// Author: leafsoar
// Date: 2015-10-27 16:19:13
//

package main

import (
	"log"
	"os"
	fp "path/filepath"
	"strings"
)

// 计算每个文件的 MD5 值

func main() {
	log.Println("leafsoar v5 ~")

	// 首先获取 assets 目录下面的目录列表
	list := GetSubDir("assets")
	log.Println(list)
}

// GetSubDir 获取一个目录的子目录，第一级
func GetSubDir(root string) []string {
	var slice []string
	fp.Walk(root, func(path string, f os.FileInfo, err error) error {
		val := strings.Replace(path, root+"/", "", 1)
		// 如果不是目录直接返回
		if !f.IsDir() ||
			strings.EqualFold(val, root) ||
			!strings.EqualFold(val, f.Name()) {
			return nil
		}
		// log.Println(path)
		slice = append(slice, path)
		return nil
	})
	return slice
}
