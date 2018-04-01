package main

import (
	"net/http"
	"fmt"
	"net/url"
	"strings"
	"io/ioutil"
)

type StringServer string

func (s StringServer) ServeHTTP(rw http.ResponseWriter, rq *http.Request){
	rq.ParseForm()
	fmt.Printf("Received form data: %v\n", rq.Form)
	fmt.Printf("Received Header %v\n", rq.Header)
	rw.Write([]byte(string(s)))
}

func createServer(addr string) http.Server{
	return http.Server{
		Addr: addr,
		Handler: StringServer("Hello Golang!"),
	}
}

const addr = "localhost:7070"

func main(){
	s := createServer(addr)
	go s.ListenAndServe()

	form := url.Values{}
	form.Set("name", "ritesh")
	form.Set("id", "5")

	req, err := http.NewRequest(http.MethodPost, "http://localhost:7070", strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil{
		panic(err)
	}
	data , err := ioutil.ReadAll(res.Body)
	if err != nil { panic(err)}
	res.Body.Close()

	fmt.Println("Response from the server: "+string(data))
}
