package main

import (
	"bufio"
	"fmt"
	"os"
)

func Reverse(s string) string {
	rune := []rune(s)
	n := len(rune)
	for i := 0; i < n/2; i++ { 
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i] 
	} 
	return string(rune)
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println(Reverse(s.Text()))
	}
}
