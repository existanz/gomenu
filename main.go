package main

import (
	"fmt"

	"log"

	"github.com/pkg/term"
)

const (
	up     byte = 65
	down   byte = 66
	escape byte = 27
	enter  byte = 13
)

func getInput() byte {
	t, _ := term.Open("/dev/tty")

	err := term.RawMode(t)
	if err != nil {
		log.Fatal(err)
	}

	var read int
	readBytes := make([]byte, 3)
	read, _ = t.Read(readBytes)

	t.Restore()
	t.Close()
	if read == 3 {
		return readBytes[2]
	}
	return readBytes[0]
}
func main() {
	var ch byte = 0
	for ch != escape {
		ch = getInput()
		fmt.Println("Choice:", ch)
	}
}
