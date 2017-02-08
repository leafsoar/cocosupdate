//
// Author: leafsoar
// Date: 2015-11-02 10:15:40
//

package util

import (
	"fmt"
	"sort"
	"testing"
)

func TestBase(t *testing.T) {
	fmt.Println("leafsoar v5 ~")
	// files := GetFiles("../assets").
	// 	FilterRemove(".DS_Store")

	// for _, item := range files {
	// 	fmt.Println(item.Path)
	// }
	// fmt.Println(files)

	// path := "../assets/1.0.0/data.bin"
	// temp := strings.Replace(path, "../", "", 1)
	// fmt.Println(temp)

	testVersions(t)
}

func testVersions(t *testing.T) {
	vs := []string{
		"1.0.13",
		"1.0.3",
		"1.2.1",
		"1.3.8",
	}
	sort.Sort(VersionSlice(vs))
	fmt.Println(vs)
}
