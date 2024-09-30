# Gomenu
Cli menu for go programs

### Usage:
import "github.com/existanz/gomenu"  

create new menu `m := gomenu.NewMenu("Chose your destiny")`
add menu items
```go
	m.AddItem("1", "Novice")
	m.AddItem("2", "Veteran")
	m.AddItem("3", "Evil")
```
if you want to add unpickable menu item for example delimiters use  
`m.AddUnpickableItem("", "---------------")`  

you can add submenu for menuitem
```go
  sm1 := gomenu.NewMenu("Novice")
  m.GetMenuItem("1").SubMenu = sm1
```
get result of choosin menu section
```go
res := m.Load()
if res != "" {
  fmt.Println("Selected item is:", res)
}
```

for more information see `example.go` file