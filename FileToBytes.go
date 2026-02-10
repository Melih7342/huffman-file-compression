package main

import (
	"fmt"
	"os"
)

func FileToBytes(path string) ([]byte, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}

	defer file.Close()

	stats, err := file.Stat()

	if err != nil {
		return nil, err
	}

	if stats.IsDir() {
		return nil, fmt.Errorf("%s is a directory", path)
	}

	data := make([]byte, stats.Size())

	_, err = file.Read(data)

	if err != nil {
		return nil, fmt.Errorf("could not read file %v", path)
	}

	return data, nil
}
