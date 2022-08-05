package programas

import "fmt"

func slice() {
	var slice1 []int
	slice2 := make([]int, 3, 5)
	if slice1 == nil {
		fmt.Println("esta vacio")
	} else {
		fmt.Println(len(slice1))
	}
	if slice2 == nil {
		fmt.Println("esta vacio")
	} else {
		fmt.Println(cap(slice2))
		fmt.Println(len(slice2))
		fmt.Println(slice2)
		slice := append(slice2, 2)
		fmt.Println(slice)
	}

	//copy
	fmt.Println("---copy---")

	slice4 := []int{1, 2, 3, 4}
	copia := make([]int, len(slice4))

	copy(copia, slice4)

	fmt.Println(slice4)
	fmt.Println(copia)
}
