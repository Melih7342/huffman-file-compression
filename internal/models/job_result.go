package models

import "time"

type JobResult struct {
	Path         string
	OriginalSize int64
	NewSize      int64
	Duration     time.Duration
	Error        error
}
