package algorithm

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Melih7342/huffman-file-compression/internal/models"
)

func TestDetermineFinalPath(t *testing.T) {
	tempDir := t.TempDir()

	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(tempDir, "file2.txt.huff")

	os.WriteFile(file1, []byte("file1"), 0644)
	os.WriteFile(file2, []byte("file2"), 0644)

	fileCount := 1

	test := []struct {
		name           string
		input          string
		cfg            models.Config
		expectedOutput string
		expectError    bool
	}{
		{
			name:  "Compress mode with standard .txt extension",
			input: file1,
			cfg: models.Config{
				CompressMode: true,
				OutputPath:   tempDir,
			},
			expectedOutput: "file1.txt.huff",
		},
		{
			name:  "Compress mode with .huff extension (already compressed)",
			input: file2,
			cfg: models.Config{
				CompressMode: true,
				OutputPath:   tempDir,
			},
			expectedOutput: "",
			expectError:    true,
		},
		{
			name:  "Decompress mode with .huff extension",
			input: file2,
			cfg: models.Config{
				DecompressMode: true,
				OutputPath:     tempDir,
			},
			expectedOutput: "file2.txt",
		},
		{
			name:        "Provide a directory as source path",
			input:       tempDir,
			cfg:         models.Config{},
			expectError: true,
		},
		{
			name:  "Provide a target file name",
			input: file1,
			cfg: models.Config{
				CompressMode: true,
				OutputPath:   "customfilename.huff",
			},
			expectedOutput: "customfilename.huff",
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result, err := DetermineFinalPath(tt.input, tt.cfg, fileCount)

			if err != nil && !tt.expectError {
				t.Fatalf("unexpected error for %s: %v", tt.name, err)
			}

			if err == nil && tt.expectError {
				t.Fatalf("expected error for %s, but got none", tt.name)
			}

			if tt.expectError {
				return
			}

			resultFileName := filepath.Base(result)
			if resultFileName != tt.expectedOutput {
				t.Errorf("%s: expecting filename %s, got %s", tt.name, tt.expectedOutput, resultFileName)
			}
		})
	}
}
