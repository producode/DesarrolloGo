package programas

import (
	"fmt"
	"io/ioutil"
)

func read() {
	file_data, err := ioutil.ReadFile("./read.txt")

	if err != nil {
		fmt.Println("Hubo un error")
	}
	fmt.Println(string(file_data))
}
