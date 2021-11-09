//2.	Implementați un server HTTP folosind pachetul net/http și care să permită apeluri pentru rute predefinite (e.g., /hello)

package main

import (
	"fmt"
	"net/http"
)

func salut(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "salut\n")
}

func antete(w http.ResponseWriter, request *http.Request) {
	for nume, antete := range request.Header {
		for _, a := range antete {
			fmt.Fprintf(w, "%v -> %v\n", nume, a)
		}
	}
}

func main() {
	http.HandleFunc("/salut", salut)
	http.HandleFunc("/antete", antete)

	http.ListenAndServe(":8090", nil)
}
