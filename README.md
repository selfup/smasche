# Smasche

Example of how one can use [gdsm](https://github.com/selfup/gdsm)

Code snippet is `main.go`

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/selfup/gdsm/gdsm"
)

func main() {
	daemon := gdsm.BuildDaemon()
	go gdsm.BootDaemon(daemon)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "workers, %q", daemon.Nodes())
	})

	log.Fatal(http.ListenAndServe(":9000", nil))
}
```

Load balanced smasche http servers via haproxy.

Each smasche server uses a gdsm worker that connects to the gdsm manager.

A simple `curl` will return all nodes in the cluster:

```
selfup@win42 MINGW64 ~/go/src/github.com/selfup/smasche (master)
$ curl localhost:9000
workers, ["172.24.0.3"]
selfup@win42 MINGW64 ~/go/src/github.com/selfup/smasche (master)
$ docker-compose scale smasche=8
Starting smasche_smasche_1 ... done
Creating smasche_smasche_2 ... done
Creating smasche_smasche_3 ... done
Creating smasche_smasche_4 ... done
Creating smasche_smasche_5 ... done
Creating smasche_smasche_6 ... done
Creating smasche_smasche_7 ... done
Creating smasche_smasche_8 ... done

selfup@win42 MINGW64 ~/go/src/github.com/selfup/smasche (master)
$ curl localhost:9000
workers, ["172.24.0.11" "172.24.0.3" "172.24.0.5" "172.24.0.8" "172.24.0.6" "172.24.0.7" "172.24.0.9" "172.24.0.10"]
```
