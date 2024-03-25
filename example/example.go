package main

import (
	"fmt"
	"github/existanz/gomenu"
)

var sub1mis []*gomenu.MenuItem = []*gomenu.MenuItem{
	{Label: "Easy", ID: "1_1"},
	{Label: "Rookie", ID: "1_2"},
	{Label: "Beginner", ID: "1_3"},
}

var sub2mis []*gomenu.MenuItem = []*gomenu.MenuItem{
	{Label: "Medium", ID: "2_1"},
	{Label: "Normal", ID: "2_2"},
	{Label: "Average", ID: "2_3"},
}

var sub3mis []*gomenu.MenuItem = []*gomenu.MenuItem{
	{Label: "Hard", ID: "3_1"},
	{Label: "Expert", ID: "3_2"},
	{Label: "Difficult", ID: "3_3"},
}

var mis []*gomenu.MenuItem = []*gomenu.MenuItem{
	{Label: "Novice", ID: "1", SubMenu: &gomenu.Menu{Prompt: "Novice", Items: sub1mis}},
	{Label: "Veteran", ID: "2", SubMenu: &gomenu.Menu{Prompt: "Veteran", Items: sub2mis}},
	{Label: "Evil", ID: "3", SubMenu: &gomenu.Menu{Prompt: "Evil", Items: sub3mis}},
}

func main() {
	m := gomenu.NewMenu("Start Menu")
	m.Items = mis
	sel := m.Load()
	fmt.Println("Selected menu ID", sel)
}
