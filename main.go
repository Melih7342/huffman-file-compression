package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Melih7342/huffman-file-compression/internal/algorithm"
)

func main() {
	// Define flags
	compressMode := flag.Bool("c", false, "Compress files")
	decompressMode := flag.Bool("d", false, "DecompressFile files")
	// verbosity := flag.Bool("v", false, "Verbosity")
	// directory := flag.String("r", "", "Recursive directory content compression")
	// help := flag.Bool("h", false, "Help")
	outputPath := flag.String("o", "", "Output file path")

	// Read files from command line
	flag.Parse()
	files := flag.Args()

	if len(files) == 0 {
		fmt.Println("Usage: huffman -c/-d file1 file2 ...")
		return
	}

	for _, path := range files {
		var finalPath string
		if *outputPath != "" {
			info, err := os.Stat(*outputPath)
			isDir := err == nil && info.IsDir()
			if len(files) > 1 || isDir {
				fileName := filepath.Base(path)
				if *compressMode {
					fileName += ".huff"
				} else if *decompressMode {
					fileName = strings.TrimSuffix(fileName, ".huff")
				}
				finalPath = filepath.Join(*outputPath, fileName)
				os.MkdirAll(*outputPath, 0755)
			} else {
				finalPath = *outputPath
			}
		} else {
			if *compressMode {
				finalPath = path + ".huff"
			} else if *decompressMode {
				finalPath = strings.TrimSuffix(path, ".huff")
			}
		}

		if *compressMode {
			err := algorithm.CompressFile(path, finalPath)
			if err != nil {
				return
			}
		} else if *decompressMode {
			err := algorithm.DecompressFile(path, finalPath)
			if err != nil {
				return
			}
		}
	}
}
