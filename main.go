package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Melih7342/huffman-file-compression/internal/algorithm"
	"github.com/Melih7342/huffman-file-compression/internal/models"
)

func main() {

	// Parsing the flags
	cfg := models.ParseConfig()

	files := algorithm.DetermineFiles(*cfg)

	for _, path := range files {

		finalPath := algorithm.DetermineFinalPath(path, *cfg, len(files))

		if cfg.CompressMode {
			start := time.Now()

			err := algorithm.CompressFile(path, finalPath, cfg.Verbosity)
			if err != nil {
				if strings.Contains(err.Error(), "compression inefficient") {
					continue
				}
				fmt.Printf("Error compressing %s: %v\n", path, err)
				return
			}
			end := time.Now()

			if cfg.Performance {
				fmt.Println("Compressing took", end.Sub(start))
				err := algorithm.SizeReduction(path, finalPath)
				if err != nil {
					return
				}
			}

		} else if cfg.DecompressMode {
			start := time.Now()

			err := algorithm.DecompressFile(path, finalPath, cfg.Verbosity)
			if err != nil {
				return
			}
			end := time.Now()

			if cfg.Performance {
				fmt.Println("Decompressing took", end.Sub(start))
			}
		}
	}
}
