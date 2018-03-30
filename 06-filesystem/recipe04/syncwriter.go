package main

import (
	"sync"
	"io"
	"os"
	"fmt"
)

type SyncWriter struct{
	 m sync.Mutex
	 Writer io.Writer
}

func (w *SyncWriter) Write(b []byte) (n int, err error){
	w.m.Lock()
	defer w.m.Unlock()
	return w.Writer.Write(b)
}

var data = []string{
	"Hello!", "Ola!", "Ahoj!",
}

func main(){
	f, err := os.Create("sample.file")
	if err != nil{
		panic(err)
	}
	wr := &SyncWriter{sync.Mutex{}, f}
	wg := sync.WaitGroup{}
	for _, val := range data{
		wg.Add(1)
		go func(greeting string){
			fmt.Fprintln(wr, greeting)
			wg.Done()
		}(val)
	}
	wg.Wait()
}
