package main

import (
	"fmt"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	var str string = "<h1 style=\"text-align: center; background-color: #444444; color: white;\">Olá Mundo!</h1>"

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {

		n, err := fmt.Fprint(w, str)
		check(err)
		fmt.Println("Bytes number: ", n)
	})

	var err error = http.ListenAndServe(":8080", nil)
	check(err)
}
