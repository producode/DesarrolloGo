package programas

import (
	"encoding/json"
	"net/http"
)

type Curso struct {
	Titulo         string
	NumeroDeVideos int
}

type Cursos []Curso

func jsons() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cursos := Cursos{
			Curso{"curso de go", 27},
			Curso{"curso de c", 27},
			Curso{"curso de java", 27},
		}
		json.NewEncoder(w).Encode(cursos)
	})

	http.ListenAndServe(":8001", nil)
}
