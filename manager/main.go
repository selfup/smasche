package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/selfup/gdsm"
	"github.com/valyala/fasthttp"
)

func main() {
	daemon := gdsm.BuildDaemon()
	go gdsm.BootDaemon(daemon)

	var mutex sync.Mutex

	idx := 0

	h := func(ctx *fasthttp.RequestCtx) {
		mutex.Lock()
		dNode := balance(daemon, idx)
		mutex.Unlock()

		proxyURL := "http://" + dNode + ":9000/"
		_, body, err := fasthttp.Get(nil, proxyURL)
		if err != nil {
			log.Println("client get", err)
			fmt.Fprintf(ctx, "500")
		}

		fmt.Fprintf(ctx, string(body)+"\n")
	}

	if err := fasthttp.ListenAndServe(":8080", h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func balance(daemon *gdsm.Operator, idx int) string {
	dNodes := daemon.Nodes()

	var dNode string

	if idx <= len(dNodes) {
		dNode = dNodes[idx]
	} else {
		idx = 0
		dNode = dNodes[idx]
	}

	idx++

	return dNode
}
