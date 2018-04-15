package main

import "fmt"

type Node struct{
	Value int
	Next *Node
}

type HashTable struct{
	Table map[int]*Node
	Size int
}

func hashFunc(i, size int) int{
	return i%size
}

func insert(hash *HashTable, value int) int{
	index := hashFunc(value, hash.Size)
	//fmt.Printf("index for value %d is: %d\n", value, index)
	//fmt.Printf("Next Node for value %d is:  %v\n", value, hash.Table[index])
	element := Node{Value: value,Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

func traverse(hash *HashTable){
	for k := range hash.Table{
		if hash.Table[k] != nil{
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				//fmt.Printf("Next Node for value %d is:  %v\n", t.Value, t.Next)
				t = t.Next
			}
		}
		fmt.Println()
	}
}

func main(){
	table := make(map[int]*Node, 10)
	hash := &HashTable{Table: table, Size: 10}
	fmt.Println("Number of spaces: ", hash.Size)
	for i := 0; i<15; i++{
		insert(hash, i)
	}
	traverse(hash)
}