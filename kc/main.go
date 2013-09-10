package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/petar/veranda/keychain"
)

var (
	flagServer   = flag.String("s", "", "Server to look for in the keychain")
	flagUser     = flag.String("u", "", "Preferred user to look for in the keychain or empty")
	flagPassword = flag.String("p", "", "Password, if adding")
)

const usefmt = `
%s add|find [-s=server] [-u=user] [-p=password]
`

func usage() {
	fmt.Errorf(usefmt, os.Args[0])
	os.Exit(1)
}

func main() {
	flag.Parse()
	switch flag.Arg(0) {
	case "add":
	case "find":
		usr, pwd, err := keychain.UserPasswd(*flagServer, *flagUser)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("Server=%s User=%s Password=%s\n", *flagServer, usr, pwd)
	default:
		usage()
	}
}
