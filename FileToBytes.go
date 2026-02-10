package main

import (
	"bufio"
	"fmt"
	"os"
)

func FileToBytes(path string) ([]byte, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}

	defer file.Close()

}
