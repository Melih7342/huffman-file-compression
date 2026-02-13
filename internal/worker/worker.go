package worker

import (
	"time"

	"github.com/Melih7342/huffman-file-compression/internal/algorithm"
	"github.com/Melih7342/huffman-file-compression/internal/models"
)

func worker(jobs <-chan models.CompressionJob, results chan<- models.JobResult, verbosity bool) {
	for job := range jobs {
		start := time.Now()
		err := algorithm.CompressFile(job.SourcePath, job.TargetPath, verbosity)

		results <- models.JobResult{
			Path:     job.SourcePath,
			Duration: time.Since(start),
			Error:    err,
		}
	}
}
