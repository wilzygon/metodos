package metodos

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Clases  map[int]string
}

func (c Course) PrintClases() {
	text := "Las clases son: "
	for _, clases := range c.Clases {
		text += clases + ", "
	}
	fmt.Println(text[:len(text)-2])
}
