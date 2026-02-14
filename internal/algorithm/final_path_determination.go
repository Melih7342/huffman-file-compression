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

	target := source

	if cfg.OutputPath != "" {
		outputInfo, err := os.Stat(cfg.OutputPath)
		outputIsDir := err == nil && outputInfo.IsDir()

		if fileCount > 1 || outputIsDir {
			fileName := filepath.Base(source)
			if err := extensionHandling(cfg, &fileName); err != nil {
				return "", err
			}

			if err := os.MkdirAll(cfg.OutputPath, 0755); err != nil {
				return "", fmt.Errorf("could not create directory: %w", err)
			}
			return filepath.Join(cfg.OutputPath, fileName), nil
		}
		return cfg.OutputPath, nil
	}

	if err := extensionHandling(cfg, &target); err != nil {
		return "", err
	}

	return target, nil
}

func extensionHandling(cfg models.Config, filename *string) error {
	if cfg.CompressMode {
		if filepath.Ext(*filename) == ".huff" {
			return fmt.Errorf("the file %s already has .huff extension", *filename)
		}
		*filename += ".huff"
	} else if cfg.DecompressMode {
		if filepath.Ext(*filename) != ".huff" {
			return fmt.Errorf("the file %s does not have .huff extension", *filename)
		}
		*filename = strings.TrimSuffix(*filename, ".huff")
	}
	return nil
}
