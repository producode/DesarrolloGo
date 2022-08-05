package programas

import (
	"fmt"
	"strconv"
)

func tipo() {
	edad := "22"

	edad_int, _ := strconv.Atoi(edad)

	fmt.Println(10 + edad_int)
}
