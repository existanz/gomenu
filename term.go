package main

import (
	"fmt"

	"github.com/pkg/term"
)

const (
	CSI          = "\033["
	ColorBlack   = "30m"
	ColorRed     = "31m"
	ColorGreen   = "32m"
	ColorYellow  = "33m"
	ColorBlue    = "34m"
	ColorMagenta = "35m"
	ColorCyan    = "36m"
	ColorWhite   = "37m"
	ColorDefault = "39m"
)

func setColor(code string) {
	switch code {
	case ColorYellow, ColorGreen, ColorBlack, ColorBlue, ColorCyan, ColorMagenta, ColorRed:
		fmt.Printf("%v%v", CSI, code)
	default:
		fmt.Printf("%v%v", CSI, ColorDefault)
	}
}

func setBold() {
	fmt.Printf("%v1m", CSI)
}

func clearTextStyle() {
	setColor("")
	fmt.Printf("%v0m", CSI)
}

func moveCursorUp(n int) {
	fmt.Printf("%v%dA", CSI, n)
}

func cursorOn() {
	fmt.Printf("%v?25h", CSI)
}

func cursorOff() {
	fmt.Printf("%v?25l", CSI)
}
func getInput() byte {
	t, _ := term.Open("/dev/tty")

	err := term.RawMode(t)
	if err != nil {
		fmt.Println(err.Error())
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
