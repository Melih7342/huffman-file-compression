package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Melih7342/huffman-file-compression/internal/algorithm"
	"github.com/Melih7342/huffman-file-compression/internal/models"
)

func main() {

	// Parsing the flags
	cfg := models.ParseConfig()

	var files []string

	if cfg.Recursive {
		for _, path := range cfg.InputPaths {
			err := filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
				if err != nil {
					return fmt.Errorf("error walking directory %s: %v", path, err)
				}
				if !d.IsDir() {
					if cfg.CompressMode && filepath.Ext(path) == ".huff" {
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
		files = cfg.InputPaths
	}

	for _, path := range files {
		var finalPath string

		if cfg.OutputPath != "" {
			info, err := os.Stat(cfg.OutputPath)
			isDir := err == nil && info.IsDir()

			if len(files) > 1 || isDir {
				fileName := filepath.Base(path)

				if cfg.CompressMode {
					fileName += ".huff"
				} else if cfg.DecompressMode {
					fileName = strings.TrimSuffix(fileName, ".huff")
				}

				finalPath = filepath.Join(cfg.OutputPath, fileName)

				err := os.MkdirAll(cfg.OutputPath, 0755)
				if err != nil {
					fmt.Println("could not create output directory")
				}
			} else {
				finalPath = cfg.OutputPath
			}
		} else {
			if cfg.CompressMode {
				finalPath = path + ".huff"
			} else if cfg.DecompressMode {
				finalPath = strings.TrimSuffix(path, ".huff")
			}
		}

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
