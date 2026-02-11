package main

import (
	"flag"
	"fmt"

	"github.com/Melih7342/huffman-file-compression/internal/algorithm"
)

func main() {
	// Define flags
	compressMode := flag.Bool("c", false, "Compress files")
	decompressMode := flag.Bool("d", false, "DecompressFile files")

	// Read files from command line
	flag.Parse()
	files := flag.Args()

	if len(files) == 0 {
		fmt.Println("Usage: huffman -c/-d file1 file2 ...")
		return
	}

	for _, path := range files {
		if *compressMode {
			err := algorithm.CompressFile(path)
			if err != nil {
				return
			}
		} else if *decompressMode {
			err := algorithm.DecompressFile(path)
			if err != nil {
				return
			}
		}
	}
}
