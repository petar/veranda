package main

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		println("mkpass [user] [password]")
		os.Exit(1)
	}
	rand.Seed(int64(time.Now().UnixNano()))
	seed := make([]byte, 5)
	for i, _ := range seed {
		seed[i] = alpha[rand.Intn(len(alpha))]
	}
	fmt.Println(os.Args[1], string(seed), hash(string(seed) + os.Args[2]))
}

const alpha = "1234567890qwertyopasdfghjkzxcbnm"

func hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

