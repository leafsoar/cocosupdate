//
// Author: leafsoar
// Date: 2015-11-02 10:06:25
//

package base

import (
	"archive/zip"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
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

// CopyFile 复制文件
func CopyFile(srcName, dstName string) {
	dstpath := path.Dir(dstName)
	_, err := os.Stat(dstpath)
	if os.IsNotExist(err) {
		os.MkdirAll(dstpath, os.ModePerm)
	}

	src, err := os.Open(srcName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dst.Close()
	io.Copy(dst, src)
}

// CheckOrCreateDir 检测目录是否存在，不存在则创建
func CheckOrCreateDir(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
}

// ArchiveZip 压缩文件
func ArchiveZip(name, path string) {
	File, _ := os.Create(name)
	PS := strings.Split(path, "\\")
	PathName := strings.Join(PS[:len(PS)-1], "\\")
	os.Chdir(PathName)
	path = PS[len(PS)-1]
	defer File.Close()
	Zip := zip.NewWriter(File)
	defer Zip.Close()
	walk := func(Path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		inpath := strings.Replace(Path, path, "", 1)
		if strings.EqualFold(Path, path) {
			return nil
		}
		if info.IsDir() {
			// 如果是目录也需要写入
			h := &zip.FileHeader{Name: inpath, Method: zip.Deflate, Flags: 0x800}
			h.SetMode(0755 | os.ModeDir)
			Zip.CreateHeader(h)

			return nil
		}
		Src, _ := os.Open(Path)
		defer Src.Close()

		h := &zip.FileHeader{Name: inpath, Method: zip.Deflate, Flags: 0x800}
		FileName, _ := Zip.CreateHeader(h)
		io.Copy(FileName, Src)
		Zip.Flush()
		return nil
	}
	if err := fp.Walk(path, walk); err != nil {
		fmt.Println(err)
	}
}
