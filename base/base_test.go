//
// Author: leafsoar
// Date: 2015-11-02 10:15:40
//

package base

import (
	"fmt"
	"testing"
)

func TestBase(t *testing.T) {
	fmt.Println("leafsoar v5 ~")
	files := GetFiles("../assets").
		FilterRemove(".DS_Store")
	fmt.Println(files)
}
