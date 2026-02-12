package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Melih7342/huffman-file-compression/internal/algorithm"
)

func main() {
	// Define flags
	compressMode := flag.Bool("c", false, "Compress file(s)")
	decompressMode := flag.Bool("d", false, "Decompress file(s)")
	verbosity := flag.Bool("v", false, "Verbose output")
	directory := flag.Bool("r", false, "Recursive directory content compression")
	help := flag.Bool("h", false, "Help")
	performance := flag.Bool("p", false, "Performance metrics")
	outputPath := flag.String("o", "", "Choose custom output location")

	// Read files from command line
	flag.Parse()

	if *help {
		flag.Usage()
	}

	initialPaths := flag.Args()

	// console output for -h or when no files are provided
	if len(initialPaths) == 0 {
		fmt.Println("\nExamples:")
		fmt.Println(" - Compression: ./huff -c test.txt")
		fmt.Println(" - Decompression: ./huff -d test.txt.huff")
		fmt.Println(" - Define output location: ./huff -c -o C:/Users/myUser test.txt")
		fmt.Println(" - Compress files in a directory: ./huff -c -r C:/Users/myUser")
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

				err := os.MkdirAll(*outputPath, 0755)
				if err != nil {
					fmt.Println("could not create output directory")
				}
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
			start := time.Now()

			err := algorithm.CompressFile(path, finalPath, *verbosity)
			if err != nil {
				if strings.Contains(err.Error(), "compression inefficient") {
					continue
				}
				fmt.Printf("Error compressing %s: %v\n", path, err)
				return
			}
			end := time.Now()

			if *performance {
				fmt.Println("Compressing took", end.Sub(start))
				err := algorithm.SizeReduction(path, finalPath)
				if err != nil {
					return
				}
			}

		} else if *decompressMode {
			start := time.Now()

			err := algorithm.DecompressFile(path, finalPath, *verbosity)
			if err != nil {
				return
			}
			end := time.Now()

			if *performance {
				fmt.Println("Decompressing took", end.Sub(start))
			}
		}
	}
}
