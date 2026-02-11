package algorithm

import (
	"fmt"
	"os"
)

func FileToBytes(path string) ([]byte, error) {
	// Open the file provided in the path
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}

	defer file.Close()

	// Check the stats
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if stats.IsDir() {
		return nil, fmt.Errorf("%s is a directory", path)
	}

	// Create a slice for the bytes
	data := make([]byte, stats.Size())

	// Fill the byte slice with the provided file data
	_, err = file.Read(data)
	if err != nil {
		return nil, fmt.Errorf("could not read file %v", path)
	}

	return data, nil
}
