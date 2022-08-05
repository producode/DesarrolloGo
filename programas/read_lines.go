package programas

import (
	"bufio"
	"fmt"
	"os"
)

func read_lines() {
	ejecucion := readFile()
	fmt.Println(ejecucion)
}

func readFile() bool {
	archivo, err := os.Open("./reada.txt")
	defer func() {
		archivo.Close()
		fmt.Println("defer")

		recover()
	}()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println(linea, " Linea")
	}

	archivo.Close()

	if true {
		return true
	}

	fmt.Println("Nunca me ejecuto")
	return true
}
