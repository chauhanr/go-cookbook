package main

import (
	"net/http"
	"fmt"
)

const addr = "localhost:7070"

type RedirectServer struct{
	redirectCount int
}

func (r *RedirectServer) ServeHTTP(rw http.ResponseWriter, rq *http.Request){

	r.redirectCount++
	fmt.Println("Received header: "+ rq.Header.Get("Known-redirects"))
	http.Redirect(rw, rq, fmt.Sprintf("/redirect%d", r.redirectCount),http.StatusTemporaryRedirect)
}

func main(){
	s := http.Server{
		Addr: addr,
		Handler: &RedirectServer{0},
	}

	go s.ListenAndServe()

	client := http.Client{}
	redirectCount := 0

	client.CheckRedirect = func(rq *http.Request, via []*http.Request) error{
		fmt.Println("Redirected")
		if redirectCount >2 {
			return fmt.Errorf("Too many redirects")

		}
		rq.Header.Set("Known-redirects", fmt.Sprintf("%d", redirectCount))
		redirectCount++
		for _, prReq := range via{
			fmt.Printf("Previous request : %v\n", prReq.URL)
		}
		return nil
	}

	_, err := client.Get("http://"+addr)
	if err != nil{
		panic(err)
	}
}
