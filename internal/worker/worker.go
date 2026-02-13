package worker

import (
	"fmt"
	"os"
	"time"

	"github.com/Melih7342/huffman-file-compression/internal/algorithm"
	"github.com/Melih7342/huffman-file-compression/internal/models"
)

func worker(jobs <-chan models.CompressionJob, results chan<- models.JobResult, mode string, verbosity bool) {

	for job := range jobs {
		start := time.Now()
		var err error

		if mode == "c" {
			err = algorithm.CompressFile(job.SourcePath, job.TargetPath, verbosity)
		} else if mode == "d" {
			err = algorithm.DecompressFile(job.SourcePath, job.TargetPath, verbosity)
		}

		oldInfo, err := os.Stat(job.SourcePath)
		if err != nil {
			fmt.Printf("Error reading stats of source file %s", job.SourcePath)
			continue
		}
		newInfo, err := os.Stat(job.TargetPath)
		if err != nil {
			fmt.Printf("Error reading stats of target file %s", job.TargetPath)
			continue
		}

		sizeReduction, err := algorithm.SizeReduction(job.SourcePath, job.TargetPath)

		results <- models.JobResult{
			Path:          job.SourcePath,
			OriginalSize:  oldInfo.Size(),
			NewSize:       newInfo.Size(),
			Duration:      time.Since(start),
			SizeReduction: sizeReduction,
			Error:         err,
		}
	}
}
