package algorithm

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Melih7342/huffman-file-compression/internal/models"
)

func DetermineFinalPath(source string, cfg models.Config, fileCount int) string {
	var finalPath string

	if cfg.OutputPath != "" {
		info, err := os.Stat(cfg.OutputPath)
		isDir := err == nil && info.IsDir()

		if fileCount > 1 || isDir {
			fileName := filepath.Base(source)

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
			finalPath = source + ".huff"
		} else if cfg.DecompressMode {
			finalPath = strings.TrimSuffix(source, ".huff")
		}
	}
	return finalPath
}
