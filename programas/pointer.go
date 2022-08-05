package programas

import "fmt"

func pointer() {
	var x, y *int
	entero := 5

	x = &entero
	y = &entero

	fmt.Println(*x, *y)

	*x = 2

	fmt.Println(*x, *y)
}
