package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

//The code in this example lists the paths and sizes of all files and
//directories in the file tree rooted at the current directory.

func main() {
	//My directory root:
	root := "C:/Users/SERCAN/Desktop/visibility"
	err := filepath.WalkDir(root,
		//Closure Function
		func(path string, di fs.DirEntry, err error) error {
			fmt.Printf("->: %s\n", path)
			return nil
		})
	fmt.Printf("Error: %v\n", err)
}

/*
Output:
->: C:/Users/SERCAN/Desktop/visibility
->: C:\Users\SERCAN\Desktop\visibility\converter
->: C:\Users\SERCAN\Desktop\visibility\converter\go.mod
->: C:\Users\SERCAN\Desktop\visibility\converter\temperatureConverter.go
->: C:\Users\SERCAN\Desktop\visibility\go.mod
->: C:\Users\SERCAN\Desktop\visibility\temperatureConverterTest.go
*/
