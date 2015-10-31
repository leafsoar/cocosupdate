//
// Author: leafsoar
// Date: 2015-10-27 17:49:41
//

// manifest 相关

package manifest

import (
	"encoding/json"
	"strings"
)

// Manifest 资源描述文件
type manifest struct {
	PackageURL        string            `json:"packageUrl"`
	RemoteManifestURL string            `json:"remoteManifestUrl"`
	RemoteVersionURL  string            `json:"remoteVersionUrl"`
	Version           string            `json:"version"`
	GroupVersions     map[string]string `json:"groupVersions"`
	EngineVersion     string            `json:"engineVersion"`
	Assets            map[string]Asset  `json:"assets"`
	SearchPaths       []string          `json:"searchPaths"`
}

// Asset 资源
type Asset struct {
	Path       string `json:"path"`
	MD5        string `json:"md5"`
	Compressed bool   `json:"compressed"`
	Group      string `json:"group"`
}

// Manifest 配置
type Manifest struct {
	mf       *manifest // 配置原文件
	curGroup string    // 当前资源组
}

// NewManifest 创建一个新的 Manifest
func NewManifest() *Manifest {
	mf := Manifest{}
	mf.mf = &manifest{}
	mf.mf.GroupVersions = map[string]string{}
	mf.mf.Assets = map[string]Asset{}
	mf.mf.SearchPaths = []string{}
	return &mf
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
func (m *Manifest) AddGroupVersion(num string, version string) {
	m.mf.GroupVersions[num] = version
	m.curGroup = num
}

// SetEngineVersion 设置引擎版本
func (m *Manifest) SetEngineVersion(version string) {
	m.mf.EngineVersion = version
}

// AddAsset 添加普通资源，默认添加到当前组
func (m *Manifest) AddAsset(name string, path string, md5 string) {
	// TODO: 根据后缀名称判断是否为压缩资源
	m.mf.Assets[name] = Asset{
		Path:       path,
		MD5:        md5,
		Compressed: strings.Index(path, ".zip") >= 0,
		Group:      m.curGroup,
	}
}

// AddSearchPath 添加搜索路径
func (m *Manifest) AddSearchPath(path string) {
	m.mf.SearchPaths = append(m.mf.SearchPaths, path)
}

// Marshal 返回字符串
func (m *Manifest) Marshal() ([]byte, error) {
	return json.MarshalIndent(m.mf, "", "  ")
}
