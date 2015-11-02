//
// Author: leafsoar
// Date: 2015-11-02 09:56:14
//

package channel

import (
	"fmt"

	"github.com/leafsoar/cocosupdate/base"
)

// Channel 渠道相关数据
type Channel struct {
	name string // 渠道名称
	path string // 路径
}

// NewChannel 建立一个新渠道
func NewChannel(name string, path string) *Channel {
	return &Channel{
		name: name,
		path: path,
	}
}

// InitVersions 初始化版本
func (c *Channel) InitVersions() {
	paths := base.GetSubPaths(c.path)
	fmt.Println(paths)
}
