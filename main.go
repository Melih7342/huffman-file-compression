package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	compressMode := flag.Bool("c", false, "Compress files")
	decompressMode := flag.Bool("d", false, "Decompress files")

	// Read files from command line
	files := flag.Args()

	if len(files) == 0 {
		fmt.Println("Usage: huffman -c/-d file1 file2 ...")
		return
	}

	for _, path := range files {
		if *compressMode {

		}
	}
}
