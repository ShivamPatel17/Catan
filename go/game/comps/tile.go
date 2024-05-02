package comps

import "fmt"

type Tile struct {
	number int
}

func (t *Tile) Print() {
	fmt.Println(t.number)
}
