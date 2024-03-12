package main

import (
	"fmt"
	"gomenu"
)

var mis []*gomenu.MenuItem = []*gomenu.MenuItem{
	{Label: "Novice", ID: "1"},
	{Label: "Veteran", ID: "2"},
	{Label: "Evile", ID: "3"},
}

func main() {
	m := gomenu.NewMenu("Start Menu")
	m.Items = mis
	sel := m.Load()
	fmt.Println("Selected menu ID", sel)
}
