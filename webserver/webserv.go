package main

import (
	"log"
	"net/http"

	"github.com/jfixby/gotest/lissajous/lissfunc"
)

func main() {

	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	liss.Lissajous(w)
}
