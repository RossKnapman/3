package web

import (
	"code.google.com/p/mx3/engine"
	"code.google.com/p/mx3/gui"
	"code.google.com/p/mx3/util"
	"log"
	"net/http"
	"runtime"
)

var doc *gui.Doc

// Start web gui on given port, does not block.
func GoServe(port string) {

	doc = gui.NewDoc("/", templText)
	http.HandleFunc("/render/", render)

	size := engine.Mesh().Size()
	doc.SetValue("nx", size[2])
	doc.SetValue("ny", size[1])
	doc.SetValue("nz", size[0])

	log.Print(" =====\n open your browser and visit http://localhost", port, "\n =====\n")
	go func() {
		util.LogErr(http.ListenAndServe(port, nil))
	}()
	runtime.Gosched()
}
