package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

func usage() {
	println("dirhttp {dir} {port}")
	os.Exit(1)
}
func main() {
	if len(os.Args) != 3 {
		usage()
	}
	port, err := strconv.Atoi(os.Args[2])
	if err != nil {
		usage()
	}
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), http.FileServer(http.Dir(os.Args[1]))))
}
