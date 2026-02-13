package worker

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/Melih7342/huffman-file-compression/internal/models"
)

func Engine(sourcePaths []string, finalPaths []string, mode string, verbosity bool) {
	jobs := make(chan models.CompressionJob, len(sourcePaths))
	results := make(chan models.JobResult, len(sourcePaths))
	numJobs := len(sourcePaths)

	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()
	if numWorkers > numJobs {
		numWorkers = numJobs
	}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(jobs, results, mode, verbosity)
		}()
	}

	for i := 0; i < numJobs; i++ {
		jobs <- models.CompressionJob{
			SourcePath: sourcePaths[i],
			TargetPath: finalPaths[i],
		}
	}

	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.Error != nil {
			fmt.Printf("Error at %s: %v\n", result.Path, result.Error)
		} else {
			fmt.Printf("finished: %s in %v\n", result.Path, result.Duration)
		}
	}
}
