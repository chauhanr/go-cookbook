package main

import (
	"sync"
	"fmt"
)

var names = []string{"John", "Jai", "Jack", "Jo", "Steve", "Som"}

type SyncList struct {
	slice []interface{}
	m sync.Mutex
}

func NewSyncList(cap int) *SyncList{
	return &SyncList{m: sync.Mutex{},slice: make([]interface{},cap),}
}

func (l *SyncList) Load(i int) interface{}{
	l.m.Lock()
	defer l.m.Unlock()
	return l.slice[i]
}

func (l *SyncList) Store(i int, val interface{}){
	l.m.Lock()
	defer l.m.Unlock()
	l.slice[i] = val
}

func (l *SyncList) Append(val interface{}){
	l.m.Lock()
	defer l.m.Unlock()
	l.slice = append(l.slice,val)
}

func main(){
	l := NewSyncList(0)
	wg := &sync.WaitGroup{}

	wg.Add(len(names))
	for i:=0; i<len(names); i++{
		go func(idx int){
			l.Append(names[idx])
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i:=0; i <len(names); i++{
		fmt.Printf("Val: %v stored at idx: %d\n", l.Load(i), i)
	}
}

