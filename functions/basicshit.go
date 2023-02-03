package lib

import (
	"github.com/taubyte/go-sdk/event"
)

//go:wasm-module http/68747470733a2f2f6769746875622e636f6d2f746166736565722d6b68616e2f746573745f73747566662f7261772f6d61696e2f616c6d6f73746e6f7468696e672e7761736d
//export do_stuff
func do_stuff(x event.Event) uint32

//export doShit
func doShit(e event.Event) uint32 {
	do_stuff(e)
	// mv, err := i2mv.Open(mvId)
	// if err != nil {
	// 	return 1
	// }

	// data, err := ioutil.ReadAll(mv)
	// if err != nil {
	// 	return 1
	// }

	// http, err := e.HTTP()
	// if err != nil {
	// 	return 1
	// }

	// http.Write(data)
	return 0
}
