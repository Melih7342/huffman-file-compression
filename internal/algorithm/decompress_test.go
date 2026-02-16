package algorithm

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Melih7342/huffman-file-compression/internal/models"
)

func TestHuffman_Integration(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name      string
		content   string
		cfg       models.Config
		wantError bool
	}{
		{
			name:    "Simple",
			content: "hello world",
			cfg: models.Config{
				Force: true,
			},
		},
		{
			name:    "Repeated characters",
			content: "aaaaaaaaaaaaaaa",
			cfg: models.Config{
				Force: true,
			},
		},
		{
			name:    "All ASCII",
			content: "The quick brown fox jumps over the lazy dog",
			cfg: models.Config{
				Force: true,
			},
		},
		{
			name:    "Empty string",
			content: "",
			cfg: models.Config{
				Force: true,
			},
			wantError: true,
		},
		{
			name:    "Special characters",
			content: "Line1\nLine2\t@#$!%^&*()",
			cfg: models.Config{
				Force: true,
			},
		},
		{
			name:    "German Umlaute",
			content: "Österreichische Kompressionsstärke",
			cfg: models.Config{
				Force: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := filepath.Join(tmpDir, tt.name+".txt")
			comp := src + ".huff"
			restored := src + ".restored"

			// GIVEN
			os.WriteFile(src, []byte(tt.content), 0644)

			// WHEN
			err := CompressFile(src, comp, tt.cfg)
			if err != nil && !tt.wantError {
				t.Fatalf("Compression error: %v", err)
			} else if err == nil && tt.wantError {
				t.Fatalf("Expected error, but got none")
			}
			err = DecompressFile(comp, restored, tt.cfg)
			if err != nil && !tt.wantError {
				t.Fatalf("Decompression error: %v", err)
			} else if err == nil && tt.wantError {
				t.Fatalf("Expected error, but got none")
			}

			// THEN
			result, _ := os.ReadFile(restored)

			if string(result) != tt.content {
				t.Errorf("Mismatch!\nWant: %q\nGot:  %q", tt.content, string(result))
			}
		})
	}
}
