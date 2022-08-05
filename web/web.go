package web

import (
	"net/http"
)

func web() {
	/*http.HandleFunc("/holamundo", func(response http.ResponseWriter, request *http.Request) {
		io.WriteString(response, "Hola mundo")
	})*/
	fileServer := http.FileServer(http.Dir("public"))

	http.Handle("/", http.StripPrefix("/", fileServer))
	http.ListenAndServe(":8000", nil)
}

/*
	func handler(response http.ResponseWriter, request *http.Request) {
		http.ServeFile(response, request, "index.html")
	}

func handler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, request.URL.Path[1:])
}
*/
