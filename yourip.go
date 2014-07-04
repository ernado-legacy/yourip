package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"flag"
)

var (
	port = flag.Int("port", 80, "port")
	host = flag.String("host", "", "host")
	prefix = flag.String("prefix", "/", "uri prefix")
)

func PrintIp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, r.RemoteAddr)
}


func main() {
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Println("listening on", addr)
	router := httprouter.New()
	httprouter.New().GET(*prefix, PrintIp)
	log.Fatal(http.ListenAndServe(addr, router))
}
