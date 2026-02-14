package algorithm

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Melih7342/huffman-file-compression/internal/models"
)

func TestDetermineFiles(t *testing.T) {
	// GIVEN
	tmpDir := t.TempDir()

	file1 := filepath.Join(tmpDir, "file1.txt")
	subDir := filepath.Join(tmpDir, "sub")
	file2 := filepath.Join(subDir, "file2.txt")
	file3 := filepath.Join(subDir, "file3.huff")

	os.Mkdir(subDir, 0755)
	os.WriteFile(file1, []byte("a"), 0644)
	os.WriteFile(file2, []byte("b"), 0644)
	os.WriteFile(file3, []byte("c"), 0644)

	tests := []struct {
		name     string
		cfg      models.Config
		expected int
	}{
		{
			name: "Non-recursive, single file",
			cfg: models.Config{
				Recursive:  false,
				InputPaths: []string{file1},
			},
			expected: 1,
		},
		{
			name: "Recursive, should find all except .huff (in compress mode)",
			cfg: models.Config{
				Recursive:    true,
				CompressMode: true,
				InputPaths:   []string{tmpDir},
			},
			expected: 2,
		},
		{
			name: "Recursive, should find all files (in decompress mode)",
			cfg: models.Config{
				Recursive:      true,
				DecompressMode: true,
				InputPaths:     []string{tmpDir},
			},
			expected: 1,
		},
	}

	// WHEN & THEN
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			files := DetermineFiles(tt.cfg)
			if len(files) != tt.expected {
				t.Errorf("got %d files, want %d", len(files), tt.expected)
			}
		})
	}
}
