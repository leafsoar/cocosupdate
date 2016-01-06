//
// Author: leafsoar
// Date: 2015-10-27 17:49:41
//

// manifest 相关

package manifest

import (
	"encoding/json"
	"strconv"
	"strings"
)

// 更新基础配置
type base struct {
	PackageURL        string            `json:"packageUrl"`
	RemoteManifestURL string            `json:"remoteManifestUrl"`
	RemoteVersionURL  string            `json:"remoteVersionUrl"`
	Version           string            `json:"version"`
	GroupVersions     map[string]string `json:"groupVersions"`
	EngineVersion     string            `json:"engineVersion"`
}

// Manifest 资源详细描述文件
type manifest struct {
	base                         // 嵌入基本配置
	Assets      map[string]asset `json:"assets"`
	SearchPaths []string         `json:"searchPaths"`
}

// Asset 资源
type asset struct {
	Path       string `json:"path"`
	MD5        string `json:"md5"`
	Compressed bool   `json:"compressed"`
	Group      string `json:"group"`
}

// Manifest 配置
type Manifest struct {
	mf        *manifest // 配置原文件
	curGroup  string    // 当前资源组
	fileIndex int       // 当前文件编号
}

// NewManifest 创建一个新的 Manifest
func NewManifest() *Manifest {
	mf := Manifest{}
	mf.mf = &manifest{}
	mf.mf.GroupVersions = map[string]string{}
	mf.mf.Assets = map[string]asset{}
	mf.mf.SearchPaths = []string{}
	return &mf
}

// Unmarshal 从 JSON 数据解析
func (m *Manifest) Unmarshal(data []byte) {
	json.Unmarshal(data, m.mf)
}

// SetURL 设置路径
func (m *Manifest) SetURL(url string) {
	m.mf.PackageURL = url
	m.mf.RemoteManifestURL = url + "/project.manifest"
	m.mf.RemoteVersionURL = url + "/version.manifest"
}

// SetVersion 设置版本
func (m *Manifest) SetVersion(version string) {
	m.mf.Version = version
}

// AddGroupVersion 添加版本组
func (m *Manifest) AddGroupVersion(version string) {
	// 算出新的组号
	num := strconv.Itoa(len(m.mf.GroupVersions) + 1)
	m.mf.GroupVersions[num] = version
	m.curGroup = num
}

// SetEngineVersion 设置引擎版本
func (m *Manifest) SetEngineVersion(version string) {
	m.mf.EngineVersion = version
}

// GetEngineVersion 获取引擎版本
func (m *Manifest) GetEngineVersion() string {
	return m.mf.EngineVersion
}

// AddAsset 添加普通资源，默认添加到当前组
func (m *Manifest) AddAsset(path string, md5 string) {
	exis, key := m.FindAsset(path)
	// 如果有老版本的文件，则删除
	if exis {
		delete(m.mf.Assets, key)
	}
	// name 名字不能相同
	m.fileIndex = m.fileIndex + 1
	name := "asset" + strconv.Itoa(m.fileIndex)
	m.mf.Assets[name] = asset{
		Path:       path,
		MD5:        md5,
		Compressed: strings.Index(path, ".zip") >= 0,
		Group:      m.curGroup,
	}
}

// FindAsset 查找一个资源是否存在
func (m *Manifest) FindAsset(path string) (bool, string) {
	// map 遍历
	for k, v := range m.mf.Assets {
		if v.Path == path {
			return true, k
		}
	}
	return false, ""
}

// AddSearchPath 添加搜索路径
func (m *Manifest) AddSearchPath(path string) {
	m.mf.SearchPaths = append(m.mf.SearchPaths, path)
}

// Marshal 返回字符串
func (m *Manifest) Marshal() ([]byte, error) {
	return json.MarshalIndent(m.mf, "", "  ")
}

func (b *base) marshal() ([]byte, error) {
	return json.MarshalIndent(b, "", "  ")
}

// MarshalMini 返回简单描述字符串
func (m *Manifest) MarshalMini() ([]byte, error) {
	return m.mf.marshal()
}
