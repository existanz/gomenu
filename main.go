package main

import (
	"fmt"

	"log"

	"github.com/pkg/term"
)

const (
	up     byte = 65
	kUp    byte = 107 // for vimers
	wUp    byte = 119 // for wasd
	down   byte = 66
	jDown  byte = 106 // for vimers
	sDown  byte = 115 // for wasd
	escape byte = 27
	enter  byte = 13
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

type Menu struct {
	Prompt    string
	CursorPos int
	Items     []*MenuItem
}

type MenuItem struct {
	Label      string
	ID         string
	Selectable bool
	SubMenu    *Menu
}

func NewMenu(prompt string) *Menu {
	return &Menu{
		Prompt: prompt,
		Items:  make([]*MenuItem, 0),
	}
}

func (m *Menu) Render() {
	setColor(ColorCyan)
	setBold()
	fmt.Print(m.Prompt, ": \n")
	clearTextStyle()
	for i, menuItem := range m.Items {
		prefix := " "
		if i == m.CursorPos {
			prefix = ">"
			setColor(ColorYellow)
		}
		fmt.Println(prefix, menuItem.Label)
		setColor("")
	}
	moveCursorUp(len(m.Items) + 1)
}

func (m *Menu) Load() string {
	m.Render()
	cursorOff()
	defer cursorOn()
	var key byte = 0
	for {
		key = getInput()
		switch key {
		case enter:
			return m.Items[m.CursorPos].ID
		case escape:
			return ""
		case up, kUp, wUp:
			m.CursorPos--
			if m.CursorPos < 0 {
				m.CursorPos = 0
			}
		case down, jDown, sDown:
			m.CursorPos++
			if m.CursorPos >= len(m.Items) {
				m.CursorPos = len(m.Items) - 1
			}
		}
		m.Render()
	}
}

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

var menuItems []*MenuItem = []*MenuItem{
	{Label: "First", ID: "1st"},
	{Label: "Second", ID: "2nd"},
	{Label: "Third", ID: "3rd"},
	{Label: "Fours", ID: "4rs"},
	{Label: "Fifs", ID: "5fs"},
}

func main() {
	m := NewMenu("General")
	m.Items = menuItems
	sel := m.Load()
	if sel != "" {
		fmt.Println("Selected menu id =", sel)
	}
}
