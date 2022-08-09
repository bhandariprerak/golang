package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to file handling in go")
	content := "hey there, this line will be in file :)"
	file, err := os.Create("mytestfile.txt")
	fmt.Println("file handle:", file)
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	fmt.Println("length of file is:", length)
	// fmt.Println("file contents are :", string())
	defer file.Close()
	ReadFile("mytestfile.txt")
}

func ReadFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Printf("the type of data is: %T\n", databyte)
	fmt.Println("the content of data is:", string(databyte))

}
