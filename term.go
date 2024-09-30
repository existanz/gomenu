package gomenu

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

// setColor sets the color of the terminal output based on the given code.
//
// code - a string representing the color code.
//
// No return values.
func setColor(code string) {
	switch code {
	case ColorYellow, ColorGreen, ColorBlack, ColorBlue, ColorCyan, ColorMagenta, ColorRed:
		fmt.Printf("%v%v", CSI, code)
	default:
		fmt.Printf("%v%v", CSI, ColorDefault)
	}
}

// setBold sets the terminal text to bold.
//
// No parameters.
// No return values.
func setBold() {
	fmt.Printf("%v1m", CSI)
}

// clearTextStyle clears the text style of the terminal output.
//
// No parameters.
// No return values.
func clearTextStyle() {
	setColor("")
	fmt.Printf("%v0m", CSI)
}

// clearLine clears the current line of the terminal output.
//
// No parameters.
// No return values.
func clearLine() {
	fmt.Printf("%v2K", CSI)
}

// clearScreen clears the entire terminal screen.
//
// No parameters.
// No return values.
func clearScreen() {
	fmt.Printf("%vJ", CSI)
}

// moveCursorUp moves the cursor up in the terminal by a given number of lines.
//
// n - the number of lines to move the cursor up.
//
// No return values.
func moveCursorUp(n int) {
	fmt.Printf("%v%dA", CSI, n)
}

// cursorOn enables the terminal cursor.
//
// No parameters.
// No return values.
func cursorOn() {
	fmt.Printf("%v?25h", CSI)
}

// cursorOff turns off the cursor in the terminal.
//
// No parameters.
// No return values.
func cursorOff() {
	fmt.Printf("%v?25l", CSI)
}

// getInput reads a single byte of input from the terminal.
//
// No parameters.
// Returns a byte representing the input read from the terminal.
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
