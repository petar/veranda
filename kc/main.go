package main

import (
	"flag"
	"fmt"
	"os"

	"code.google.com/p/rsc/keychain"
)

var (
	flagServer = flag.String("s", "", "Server to look for in the keychain")
	flagUser   = flag.String("u", "", "Preferred user to look for in the keychain or empty")
)

func main() {
	flag.Parse()
	usr, pwd, err := keychain.UserPasswd(*flagServer, *flagUser)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Server=%s User=%s Password=%s\n", *flagServer, usr, pwd)
}
