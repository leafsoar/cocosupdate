//
// Author: leafsoar
// Date: 2015-10-27 17:49:41
//

// manifest 相关

package main

import (
	"log"

	"github.com/bitly/go-simplejson"
)

// TestJSON 测试 json
func TestJSON() {
	log.Println("leafsoar")

	js := simplejson.New()
	js.Set("packageUrl", "http://192.168.1.50/res")
	js.Set("remoteManifestUrl", "http://192.168.1.50/res")
	js.Set("remoteVersionUrl", "http://192.168.1.50/res")
	js.Set("version", "1.0.0")

	gv := simplejson.New()
	gv.Set("1", "1.0.1")
	gv.Set("2", "1.0.2")
	js.Set("groupVersions", gv)

	// at := simplejson.New()

	con, _ := js.EncodePretty()
	log.Println(string(con))
}
