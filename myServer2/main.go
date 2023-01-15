package main

import (
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./pages")))
	err := http.ListenAndServe(":8080", nil)
	check(err)

}
