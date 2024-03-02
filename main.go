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
	fmt.Println(m.Prompt, ": ")
	for i, menuItem := range m.Items {
		prefix := " "
		if i == m.CursorPos {
			prefix = ">"
		}
		fmt.Println(prefix, menuItem.Label)
	}
}

func moveCursorUp(n int) {
	fmt.Printf("\033[%dA", n)
}

func cursorOn() {
	fmt.Printf("\033[?25h")
}

func cursorOff() {
	fmt.Printf("\033[?25l")
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

func main() {
	m := NewMenu("General")
	m.Items = append(m.Items, &MenuItem{Label: "First", ID: "1st"})
	m.Items = append(m.Items, &MenuItem{Label: "Second", ID: "2nd"})
	m.Items = append(m.Items, &MenuItem{Label: "Third", ID: "3rd"})
	m.Items = append(m.Items, &MenuItem{Label: "Fours", ID: "4rs"})
	m.Items = append(m.Items, &MenuItem{Label: "Fifs", ID: "5fs"})
	m.Render()
	cursorOff()
	defer cursorOn()
	var ch byte = 0
	for ch != escape {
		ch = getInput()
		if ch == up || ch == kUp || ch == wUp {
			m.CursorPos--
			if m.CursorPos < 0 {
				m.CursorPos = 0
			}
		} else if ch == down || ch == jDown || ch == sDown {
			m.CursorPos++
			if m.CursorPos >= len(m.Items) {
				m.CursorPos = len(m.Items) - 1
			}
		}
		moveCursorUp(len(m.Items) + 1)
		m.Render()
	}
}
