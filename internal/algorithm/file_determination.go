package algorithm

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Melih7342/huffman-file-compression/internal/models"
)

func DetermineFiles(cfg models.Config) []string {
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

	return files
}
