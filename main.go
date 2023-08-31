package main

func main() {
	Go := metodos.Course{
		Name:    "Go desde cero",
		Price:   12.24,
		IsFree:  false,
		UserIDs: []uint{12, 56, 89},
		Clases: map[int]string{
			1: "Introducción",
			2: "Estructuras",
			3: "metodos",
		},
	}
	//PrintClases(Go)
	Go.PrintClases()

}

/* func PrintClases(c Course) {
	text := "Las clases son: "
	for _, clases := range c.Clases {
		text += clases + ", "
	}
	fmt.Println(text[:len(text)-2])
}
*/
