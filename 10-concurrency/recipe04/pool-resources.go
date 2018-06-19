package main

import (
	"sync"
	"fmt"
	"time"
)

type Worker struct{
	id string
}

func (w *Worker) String() string{
	return w.id
}

var globalCount = 0

var pool = sync.Pool{
	New: func() interface{}{
		res := &Worker{fmt.Sprintf("%d", globalCount)}
		globalCount++
		return res
	},
}

func main(){
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i:=0; i<10; i++{
		go func(idx int){
			w := pool.Get().(*Worker)
			fmt.Println("Go Worker id: "+w.String())
			time.Sleep(time.Second)
			pool.Put(w)
			wg.Done()
		}(i)
	}
	wg.Wait()
}