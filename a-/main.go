package main

import (
	"bufio"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		line, err := r.ReadString('\n')
		if len(line) > 0 && line[0] == '\t' {
			os.Stdout.WriteString(line[1:])
		} else {
			os.Stdout.WriteString(line)
		}
		if err != nil {
			break
		}
	}
}
