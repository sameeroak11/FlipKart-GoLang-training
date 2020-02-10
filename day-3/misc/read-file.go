package main

import (
	"fmt"
	"strings"
	"io"
	"io/ioutil"
	//"os"
)


func ReadFileContent(readFile string) (bool, []byte) {
	fmt.Printf("file: %s\n", readFile)
	byteData, readError := ioutil.ReadFile(readFile)
	if readError != nil {
		fmt.Printf("error: byteData read error from file: %s, error: %s\n",
			readFile, readError.Error())
		return false, []byte{}
	}

	return true, byteData
}


func main() {
	isOK, fileData := ReadFileContent("/home/sameeroak/ex.1.go")
	if !isOK {
		fmt.Printf("error: file read error.\n")
		return
	}

	reader := strings.NewReader(string(fileData))
	p := make([]byte, 30)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			fmt.Printf("EOF\n")
			break
		}
		fmt.Println(string(p[:n]))
	}

	fmt.Printf("back to join.\n")

	return
}
