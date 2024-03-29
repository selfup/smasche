package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/selfup/gdsm"
)

func main() {
	daemon := gdsm.BuildDaemon()
	go gdsm.BootDaemon(daemon)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{}")
	})

	log.Fatal(http.ListenAndServe(":9000", nil))
}
