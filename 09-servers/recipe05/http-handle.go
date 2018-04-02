package main

import (
	"net/http"
	"fmt"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(rw http.ResponseWriter, rq *http.Request) {
		if rq.Method == http.MethodPost{
			fmt.Fprintln(rw, "User POST")
		}
		if rq.Method == http.MethodGet{
			fmt.Fprintln(rw, "User GET")
		}
	})
	// separate handlers
	itemMux := http.NewServeMux()
	itemMux.HandleFunc("/items/clothes", func(rw http.ResponseWriter, rq *http.Request) {
		fmt.Fprintln(rw, "Clothes")
	})
	mux.Handle("/items/", itemMux)

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/ports", func(rw http.ResponseWriter, rq *http.Request) {
		fmt.Fprintln(rw, "Ports")
	})

	mux.Handle("/admin/", http.StripPrefix("/admin", adminMux))
	http.ListenAndServe(":8080", mux)
}
