package main

import (
	"fmt"
	"os"
)

func openFile(input string) (fileFromOs *os.File, openError error) {
	fileFromOs, err := os.Open(input)
	if err != nil {
		return fileFromOs, fmt.Errorf("openFile func: %w", err)
	}
	return fileFromOs, openError

}

func main() {

	fileFromOs, err := openFile("/home/jupyter/GolandProjects/rebrain_basic/02_08/in.txt")
	fmt.Println(fileFromOs, err)

}
