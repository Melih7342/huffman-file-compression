package worker

import (
	"fmt"
	"os"
	"time"

	"github.com/Melih7342/huffman-file-compression/internal/algorithm"
	"github.com/Melih7342/huffman-file-compression/internal/models"
)

func worker(jobs <-chan models.CompressionJob, results chan<- models.JobResult, mode string, cfg models.Config) {

	for job := range jobs {
		start := time.Now()
		var err error

		if mode == "c" {
			err = algorithm.CompressFile(job.SourcePath, job.TargetPath, cfg)
		} else if mode == "d" {
			err = algorithm.DecompressFile(job.SourcePath, job.TargetPath, cfg)
		}

		oldInfo, errOld := os.Stat(job.SourcePath)
		newInfo, errNew := os.Stat(job.TargetPath)

		if errOld != nil || errNew != nil {
			results <- models.JobResult{Path: job.SourcePath, Error: fmt.Errorf("stat error")}
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
