package main

import (
	g "gtracker"
	"net/http"
)

func main() {

	// Set routing rules
	http.HandleFunc("/", g.MainHandler)

	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	//Use the default DefaultServeMux.
	http.ListenAndServe("localhost:8080", nil)

}
