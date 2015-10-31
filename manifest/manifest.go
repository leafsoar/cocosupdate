//
// Author: leafsoar
// Date: 2015-10-27 17:49:41
//

// manifest 相关

package manifest

import "log"

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

// Manifest 配置
type Manifest struct {
	mf *manifest
}

// SetURL 设置路径
func (m *Manifest) SetURL(url string) {
	m.mf.PackageURL = url
	m.mf.RemoteManifestURL = url + "/project.manifest"
	m.mf.RemoteVersionURL = url + "/version.manifest"
}

// Asset 资源
type Asset struct {
	Path       string `json:"path"`
	MD5        string `json:"md5"`
	Compressed bool   `json:"compressed"`
	Group      string `json:"group"`
}

// TestJSON 测试 json
func TestJSON() {
	log.Println("leafsoar")

	// mf := Manifest{}
	// mf.SetURL("http://192.168.1.50/res")
	// mf.Version = "1.0.0"
	// mf.GroupVersions = map[string]string{
	// 	"1": "1.0.1",
	// 	"2": "1.0.1",
	// }
	// mf.EngineVersion = "3.7.1"
	// mf.SearchPaths = []string{}

	// // assets
	// assets := map[string]Asset{}
	// assets["update1"] = Asset{
	// 	Path:       "src/app.zip",
	// 	MD5:        "lskjdklfjlsjdfl",
	// 	Compressed: true,
	// 	Group:      "1",
	// }

	// mf.Assets = assets

	// con, _ := json.MarshalIndent(mf, "", "  ")
	// log.Println(string(con))
}
