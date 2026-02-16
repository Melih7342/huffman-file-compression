package worker

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/Melih7342/huffman-file-compression/internal/models"
)

func Engine(sourcePaths []string, finalPaths []string, cfg *models.Config) {
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
			worker(jobs, results, cfg.Mode, *cfg)
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

	var totalOriginal int64
	var totalNew int64
	processedCount := 0

	fmt.Printf("Processing %d files using %d workers...\n\n", numJobs, numWorkers)

	for result := range results {
		processedCount++
		if result.Error != nil {
			fmt.Printf("[%d/%d] Error at %s: %v\n", processedCount, numJobs, result.Path, result.Error)
			continue
		}
		if cfg.Performance {
			fmt.Printf("[%d/%d] ✅ %-25s | %10v | Saved: %6.2f%%\n",
				processedCount,
				numJobs,
				filepath.Base(result.Path),
				result.Duration.Truncate(time.Millisecond),
				result.SizeReduction,
			)
			totalOriginal += result.OriginalSize
			totalNew += result.NewSize
		} else {
			fmt.Printf("[%d/%d] ✅ Finished: %s\n", processedCount, numJobs, filepath.Base(result.Path))
		}
	}
	if cfg.Performance && totalOriginal > 0 && cfg.Mode == "c" {
		totalDiff := 100 - (float64(totalNew) * 100 / float64(totalOriginal))
		fmt.Println("\n" + strings.Repeat("=", 65))
		fmt.Printf("FINAL STATISTICS\n")
		fmt.Printf("Total Files:     %d\n", numJobs)
		fmt.Printf("Total Original:  %.2f MB\n", float64(totalOriginal)/(1024*1024))
		fmt.Printf("Total Packed:    %.2f MB\n", float64(totalNew)/(1024*1024))
		fmt.Printf("Overall Savings: %.2f%%\n", totalDiff)
		fmt.Println(strings.Repeat("=", 65))
	}
}
