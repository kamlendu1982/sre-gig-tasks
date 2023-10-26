package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

type page struct {
	heading string
	body    []byte
}

func request_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome home, try below\n")
	fmt.Fprint(w, "<link href=https://shenandoahvalley.org/scenic-drives-and-fall-foliage/>")

}
func main() {
	http.HandleFunc("/welcome", request_handler)
	listner, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	http.Serve(listner, nil)
}
