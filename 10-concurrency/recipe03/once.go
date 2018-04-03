package main

import (
	"sync"
	"fmt"
)

var names = []interface{}{"John", "Jai", "Jack", "Jo", "Steve", "Som"}

type Source struct{
	m *sync.Mutex
	o *sync.Once
	data []interface{}
}

func (s *Source) Pop() (interface{}, error){
	s.m.Lock()
	defer s.m.Unlock()
	s.o.Do(func(){
		s.data = names
		fmt.Println("Data has been loaded")
	})
	if len(s.data) > 0{
		res := s.data[0]
		s.data = s.data[1:]
		return res, nil
	}
	return nil, fmt.Errorf("No data available")
}

func main(){
	s := &Source{&sync.Mutex{}, &sync.Once{}, nil}
	wg := &sync.WaitGroup{}
	wg.Add(len(names))
	for i :=0; i<len(names); i++{
		go func(idx int){
			if val, err := s.Pop(); err ==nil{
				fmt.Printf("Pop %d returned: %s\n", idx, val)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
