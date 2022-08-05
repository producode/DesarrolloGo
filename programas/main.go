package programas

import (
	"bufio"
	"fmt"
	"os"
)

func algo() {
	var edad int
	fmt.Println("Ingresa tu edad: ")
	fmt.Scanf("%d\n", &edad)
	fmt.Printf("Mi edad es %d\n", edad)
	if edad > 18 {
		fmt.Println("Guau!! eres mayor de edad")
	}

	fmt.Println("Ingresa tu nombre: ")
	reader := bufio.NewReader(os.Stdin)
	nombre, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Hola %v¡Es un gusto conocerte!\n", nombre)
	}

	for i := 1; ; i = i + 1 {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 6 {
			break
		}
	}
}
