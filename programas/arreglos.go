package programas

import "fmt"

func arreglos() {
	arreglo := [3]int{1, 2}

	for i := 0; i < len(arreglo); i++ {
		fmt.Println(arreglo[i])
	}
}
