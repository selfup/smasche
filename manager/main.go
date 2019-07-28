package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/selfup/gdsm"
)

func main() {
	daemon := gdsm.BuildDaemon()
	go gdsm.BootDaemon(daemon)

	var mutex sync.Mutex

	idx := 0

	tr := &http.Transport{
		MaxIdleConns:    1024,
		IdleConnTimeout: 3 * time.Second,
	}

	client := http.Client{Transport: tr}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		dNode := balance(daemon, idx)
		log.Println("DNODE", dNode)
		mutex.Unlock()

		res, err := client.Get("http://" + dNode + ":9000/")
		if err != nil {
			log.Println("client get", err)
			w.Write([]byte("500"))
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println("body read", err)
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
