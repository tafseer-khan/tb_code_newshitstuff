package lib

import (
	"io/ioutil"

	"github.com/taubyte/go-sdk/event"
	"github.com/taubyte/go-sdk/i2mv"
)

//go:wasm-module http/68747470733a2f2f6769746875622e636f6d2f746166736565722d6b68616e2f746573745f73747566662f626c6f622f6d61696e2f68656c6c6f5f7761736d2e776173692e7761736d3f7261773d74727565
//export do_stuff
func do_stuff(x event.Event) uint32

//export doShit
func doShit(e event.Event) uint32 {
	mvId := do_stuff(e)
	mv, err := i2mv.Open(mvId)
	if err != nil {
		return 1
	}

	data, err := ioutil.ReadAll(mv)
	if err != nil {
		return 1
	}

	http, err := e.HTTP()
	if err != nil {
		return 1
	}

	http.Write(data)
	return 0
}
