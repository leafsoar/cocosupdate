//
// Author: leafsoar
// Date: 2015-10-27 16:19:13
//

package main

import (
	"log"

	"github.com/leafsoar/cocosupdate/channel"
)

func main() {
	log.Println("leafsoar v5 ~")

	ch := channel.NewChannel("default", "assets", "publish")
	ch.InitVersions()
	ch.Publish()
}
