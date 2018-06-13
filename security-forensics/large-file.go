package main

import (
	"os"
	"container/list"
	"io/ioutil"
	"log"
	"path/filepath"
	"fmt"
)

type FileNode struct{
	FullPath string
	Info os.FileInfo
}

func insertSorted(fileList *list.List, fileNode FileNode){
	if fileList.Len() == 0 {
		// as list is empty just return after insert
		fileList.PushFront(fileNode)
		return
	}
	for element := fileList.Front(); element != nil; element = element.Next(){
		if fileNode.Info.Size() < element.Value.(FileNode).Info.Size(){
			fileList.InsertBefore(fileNode, element)
			return
		}
	}
	fileList.PushBack(fileNode)
}

func getFilesInDirectoryBySize(fileList *list.List, path string){
	dirFiles, err := ioutil.ReadDir(path)
	if err != nil{
		log.Println("Error reading directory: "+ err.Error())
	}

	for _, dirFile := range dirFiles{
		fullPath := filepath.Join(path, dirFile.Name())
		if dirFile.IsDir(){
			getFilesInDirectoryBySize(fileList, fullPath)
		}else if dirFile.Mode().IsRegular(){
			insertSorted(fileList, FileNode{FullPath: fullPath, Info: dirFile})
		}
	}
}


func main(){
	fileList := list.New()
	getFilesInDirectoryBySize(fileList, "/home")
	for element := fileList.Front(); element != nil; element = element.Next(){
		fmt.Printf("%d ", element.Value.(FileNode).Info.Size())
		fmt.Printf("%s\n", element.Value.(FileNode).FullPath)
	}
}
