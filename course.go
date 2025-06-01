package course

import "fmt"

type course struct {
	name    string
	price   float64
	isFree  bool
	userIDs []uint
	classes map[uint]string
}

func (c *course) SetName(name string) { c.name = name} // Metodo set en GO
func (c *course) Name() string { return c.name } // Metodo get en GO

func (c *course) SetPrice(price float64) { c.price = price} // Metodo set en GO
func (c *course) Price() float64 { return c.price } // Metodo get en GO

func (c *course) SetIsFree(isFree bool) { c.isFree = isFree} // Metodo set en GO
func (c *course) IsFree() bool { return c.isFree } // Metodo get en GO

func (c *course) SetUserIDs(userIDs []uint) { c.userIDs = userIDs} // Metodo set en GO
func (c *course) UserIDs() []uint { return c.userIDs } // Metodo get en GO

func (c *course) SetClasses(classes map[uint]string) {
	c.classes = classes
}

// Funcion constructora
func New(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}

	return &course{
		name: name,
		price: price,
		isFree: isFree,
	}
}


//	De la siguiente manera no se actualiza el metodo por que apunta hacia una copia
//  debo pasar un puntero para solucionar el siguiente error.
// func (c Course) ChangePrice(price float64) {
// 	c.Price = price
// }

// Asi se debe hacer.
// Esta funcion ya no sera necesaria con los metodos "get" y "set".
// func (c *course) changePrice(price float64) {
// 	c.Price = price
// }



func (c course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.classes {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}
