package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode/utf8"
)

func main() {
	log.SetFlags(log.Lshortfile)
	//
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		u, err := strconv.ParseUint(scanner.Text(), 16, 32)
		if err != nil {
			log.Printf("parse rune hex rep'n")
			continue
		}
		p := make([]byte, 4)
		n := utf8.EncodeRune(p, rune(u))
		fmt.Println(string(p[:n]))
	}
}
