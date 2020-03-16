package main

import (
	"fmt"
	"io"
	"os"
)

func main () {
	if len(os.Args) <= 1 {
		fmt.Println("you should pass \"filename\" argument after running the main function.")
		os.Exit(1)
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)

	// go run main.go myfile.txt
	// go run main.go help/readme.md

}
