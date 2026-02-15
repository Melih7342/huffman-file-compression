package algorithm

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSizeReduction(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name         string
		originalSize int
		packedSize   int
		expectedDiff float64
	}{
		{
			name:         "50 percent reduction",
			originalSize: 100,
			packedSize:   50,
			expectedDiff: 50.0,
		},
		{
			name:         "No reduction",
			originalSize: 100,
			packedSize:   100,
			expectedDiff: 0.0,
		},
		{
			name:         "Inefficient compression (file grew)",
			originalSize: 100,
			packedSize:   120,
			expectedDiff: -20.0,
		},
		{
			name:         "empty old size (division by zero)",
			originalSize: 0,
			packedSize:   100,
			expectedDiff: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			origPath := filepath.Join(tmpDir, tt.name+"_orig")
			packPath := filepath.Join(tmpDir, tt.name+"_pack")

			os.WriteFile(origPath, make([]byte, tt.originalSize), 0644)
			os.WriteFile(packPath, make([]byte, tt.packedSize), 0644)

			result, err := SizeReduction(origPath, packPath)

			// Assertions
			if err != nil {
				t.Fatalf("SizeReduction failed: %v", err)
			}

			if fmt.Sprintf("%.2f", result) != fmt.Sprintf("%.2f", tt.expectedDiff) {
				t.Errorf("expected diff %.2f, got %.2f", tt.expectedDiff, result)
			}
		})
	}
}
