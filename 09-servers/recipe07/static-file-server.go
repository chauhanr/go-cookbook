package main

import "net/http"

func main(){
	fileServ := http.FileServer(http.Dir("html"))
	fileServ = http.StripPrefix("/html", fileServ)

	http.HandleFunc("/welcome", serveWelcome)
	http.Handle("/html/", fileServ)
	http.ListenAndServe(":8080", nil)

}

func serveWelcome(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"welcome.txt")
}