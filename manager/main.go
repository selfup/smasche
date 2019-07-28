package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/selfup/gdsm"
	"github.com/valyala/fasthttp"
)

func main() {
	daemon := gdsm.BuildDaemon()
	go gdsm.BootDaemon(daemon)

	var mutex sync.Mutex

	idx := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		dNode := balance(daemon, idx)
		mutex.Unlock()

		proxyURL := "http://" + dNode + ":9000/"
		_, body, err := fasthttp.Get(nil, proxyURL)
		if err != nil {
			log.Println("client get", err)
			w.Write([]byte("500"))
		}

		w.Write(body)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
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
