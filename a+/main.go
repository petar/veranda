package main

import (
	"bufio"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		line, err := r.ReadString('\n')
		os.Stdout.Write([]byte{'\t'})
		os.Stdout.WriteString(line)
		if err != nil {
			break
		}
	}
}
