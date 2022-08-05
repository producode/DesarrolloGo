package programas

import "fmt"

type User struct {
	edad             int
	nombre, apellido string
}

func (usuario User) nombre_completo() string {
	return usuario.nombre + " " + usuario.apellido
}

func (this *User) set_name(n string) {
	this.nombre = n
}

func structs() {
	usuario := User{nombre: "Jose", apellido: "Herrero", edad: 24}

	fmt.Println(usuario)
	usuario.set_name("Manuel")
	fmt.Printf("el usuario %v, tiene %d años", usuario.nombre_completo(), usuario.edad)
}
