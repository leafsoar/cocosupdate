//
// Author: leafsoar
// Date: 2015-10-31 15:01:12
//

// manifest 测试

package manifest

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	fmt.Println("leafsoar v 5 ~")

	mf := NewManifest()
	// 设置 url 以及版本信息
	mf.SetURL("http://192.168.1.50/res")
	mf.SetVersion("1.0.0")
	mf.SetEngineVersion("3.7.1")

	// 添加组以及组资源
	mf.AddGroupVersion("1", "1.0.1")
	mf.AddAsset("update1", "res/update.zip", "leafsoar")
	mf.AddAsset("test", "res/test.js", "leafsoar")

	mf.AddGroupVersion("2", "1.0.2")
	mf.AddAsset("test2", "res/test.js", "leafsoar")

	// 添加搜索路径
	mf.AddSearchPath("res")
	mf.AddSearchPath("src")

	// 返回简短描述
	con, _ := mf.MarshalMini()
	fmt.Println(string(con))

	// 返回详细描述
	con, _ = mf.Marshal()
	fmt.Println(string(con))

}
