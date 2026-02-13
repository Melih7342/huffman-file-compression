package worker

import (
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

		results <- models.JobResult{
			Path:     job.SourcePath,
			Duration: time.Since(start),
			Error:    err,
		}
	}
}
