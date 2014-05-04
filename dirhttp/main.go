package main

import (
	"log"
	"net/http"
	"os"
)

func usage() {
	println("dirhttp {dir} {addr}")
	os.Exit(1)
}
func main() {
	if len(os.Args) != 3 {
		usage()
	}
	log.Fatal(http.ListenAndServe(os.Args[2], http.FileServer(http.Dir(os.Args[1]))))
}
