package course

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}


//	De la siguiente manera no se actualiza el metodo por que apunta hacia una copia
//  debo pasar un puntero para solucionar el siguiente error.
// func (c Course) ChangePrice(price float64) {
// 	c.Price = price
// }

// Asi se debe hacer.
func (c *Course) ChangePrice(price float64) {
	c.Price = price
}

func (c Course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}
