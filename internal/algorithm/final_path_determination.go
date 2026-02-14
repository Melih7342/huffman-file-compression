package algorithm

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Melih7342/huffman-file-compression/internal/models"
)

func DetermineFinalPath(source string, cfg models.Config, fileCount int) (string, error) {
	sourceInfo, err := os.Stat(source)
	if err != nil {
		return "", fmt.Errorf("could not stat source: %w", err)
	}
	if sourceInfo.IsDir() {
		return "", fmt.Errorf("please provide a filepath")
	}

	var finalPath string

	if cfg.OutputPath != "" {
		outputInfo, err := os.Stat(cfg.OutputPath)
		outputIsDir := err == nil && outputInfo.IsDir()

		if fileCount > 1 || outputIsDir {
			fileName := filepath.Base(source)

			if cfg.CompressMode {
				if filepath.Ext(fileName) != ".huff" {
					fileName += ".huff"
				} else {
					return "", fmt.Errorf("the file %s has extension .huff and likely already compressed", fileName)
				}
			} else if cfg.DecompressMode {
				if filepath.Ext(fileName) == ".huff" {
					fileName = strings.TrimSuffix(fileName, ".huff")
				} else {
					return "", fmt.Errorf("the file %s doesn't have extension .huff and likely isn't compressed", fileName)
				}
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
			if filepath.Ext(source) != ".huff" {
				source += ".huff"
			} else {
				return "", fmt.Errorf("the file %s has extension .huff and likely already compressed", source)
			}
		} else if cfg.DecompressMode {
			if filepath.Ext(source) == ".huff" {
				source = strings.TrimSuffix(source, ".huff")
			} else {
				return "", fmt.Errorf("the file %s doesn't have extension .huff and likely isn't compressed", source)
			}
		}
	}
	return finalPath, nil
}
