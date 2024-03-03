## This is cli menu for go programs

### Usage:
import "github.com/existanz/gomenu"  
declare 
```Go
var mis []*gomenu.MenuItem = []*gomenu.MenuItem{
  {Label: "Novice",  ID: "1" },
  {Label: "Veteran", ID: "2" },
  {Label: "Evile",   ID: "3" },
}
```  
create new menu `m := gomenu.NewMenu("Chose your destiny")`  
set `m.Items = mis`  
get result of choosin menu section
```go
res := m.Load()
if res != "" {
  fmt.Println("Selected item is:", res)
}
```