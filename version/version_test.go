//
// Author: leafsoar
// Date: 2016-01-06 11:05:37
//

package version

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	fmt.Println("leafsoar v5 ~")
	// 创建一个新版本
	ver := NewVersion("1.0.0", "../assets/1.0.1")
	engine := ver.GetEngineVersion()
	fmt.Println("engine:", engine)
}
