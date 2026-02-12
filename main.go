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
	verbosity := flag.Bool("v", false, "Verbosity")
	directory := flag.Bool("r", false, "Recursive directory content compression")
	// help := flag.Bool("h", false, "Help")
	outputPath := flag.String("o", "", "Output file path")

	// Read files from command line
	flag.Parse()
	initialPaths := flag.Args()

	if len(initialPaths) == 0 {
		fmt.Println("Usage: huffman -c/-d file1 file2 ...")
		return
	}

	var files []string

	if *directory {
		for _, path := range initialPaths {
			err := filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
				if err != nil {
					return fmt.Errorf("error walking directory %s: %v", path, err)
				}
				if !d.IsDir() {
					if *compressMode && filepath.Ext(path) == ".huff" {
						return nil
					}
					files = append(files, path)
				}
				return nil
			})
			if err != nil {
				fmt.Printf("error reading directory %s: %v", path, err)
			}
		}
	} else {
		files = initialPaths
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
			err := algorithm.CompressFile(path, finalPath, *verbosity)
			if err != nil {
				return
			}
		} else if *decompressMode {
			err := algorithm.DecompressFile(path, finalPath, *verbosity)
			if err != nil {
				return
			}
		}
	}
}
