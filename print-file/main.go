package main

import (
	"fmt"
	"io"
	"os"
)

// func main() {
// 	content, err := os.ReadFile(os.Args[1])
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(1)
// 	}

// 	os.Stdout.Write(content)
// }

func main() {
	filePtr, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, filePtr)
}
