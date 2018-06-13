package main

import (
	"container/list"
				"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"os"
)

type FN struct{
	FullPath string
	Info os.FileInfo
}

func insertSortedModTime(fileList *list.List, fileNode FN){
	if fileList.Len() == 0 {
		// as list is empty just return after insert
		fileList.PushFront(fileNode)
		return
	}
	for element := fileList.Front(); element != nil; element = element.Next(){
		if fileNode.Info.ModTime().Before(element.Value.(FN).Info.ModTime()) {
			fileList.InsertBefore(fileNode, element)
			return
		}
	}
	fileList.PushBack(fileNode)
}


func GetFilesInDirectoryBySize(fileList *list.List, path string){
	dirFiles, err := ioutil.ReadDir(path)
	if err != nil{
		log.Println("Error reading directory: "+ err.Error())
	}

	for _, dirFile := range dirFiles{
		fullPath := filepath.Join(path, dirFile.Name())
		if dirFile.IsDir(){
			GetFilesInDirectoryBySize(fileList, fullPath)
		}else if dirFile.Mode().IsRegular(){
			insertSortedModTime(fileList, FN{FullPath: fullPath, Info: dirFile})
		}
	}
}

func main(){
	fileList := list.New()
	GetFilesInDirectoryBySize(fileList, "/home/vagrant/go-projects/src/go-cookbook")
	for element := fileList.Front(); element != nil; element = element.Next(){
		fmt.Print(element.Value.(FN).Info.ModTime())
		fmt.Printf(" %s\n", element.Value.(FN).FullPath)
	}
}
