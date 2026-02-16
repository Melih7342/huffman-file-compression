package main

import (
	"github.com/Melih7342/huffman-file-compression/internal/algorithm"
	"github.com/Melih7342/huffman-file-compression/internal/models"
	"github.com/Melih7342/huffman-file-compression/internal/worker"
)

func main() {

	// Parsing the flags
	cfg := models.ParseConfig()

	files := algorithm.DetermineFiles(*cfg)

	finalPaths := make([]string, len(files))

	for i, file := range files {
		finalPaths[i], _ = algorithm.DetermineFinalPath(file, *cfg, len(files))
	}

	worker.Engine(files, finalPaths, cfg)

}
