package programas

import "fmt"

type Humano struct {
	nombre string
}

func (this Humano) hablar() string {
	return "bla bla bla"
}

type Tutor struct {
	Humano
	Dummy
}

type Dummy struct {
	nombre string
}

func (this Tutor) hablar() string {
	return "la la la"
}

func campos_anonimos() {
	Tutor := Tutor{Humano{"Jaime"}, Dummy{"Pedro"}}

	fmt.Println(Tutor.Humano.nombre)
	fmt.Println(Tutor.hablar())
}
