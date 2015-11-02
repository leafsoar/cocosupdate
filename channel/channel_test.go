//
// Author: leafsoar
// Date: 2015-11-02 10:01:52
//

package channel

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	fmt.Println("leafsoar v5 ~")
	// 建立一个渠道更新 (渠道名称，渠道资源目录)
	channel := NewChannel("default", "../assets")
	channel.InitVersions()
	channel.Publish()
}
