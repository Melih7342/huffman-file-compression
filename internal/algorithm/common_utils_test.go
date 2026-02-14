package algorithm

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileToBytes_Table(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(filePath, []byte("data"), 0644)

	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{"Valid file", filePath, false},
		{"Missing file", "missing.txt", true},
		{"Directory instead of file", tmpDir, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := FileToBytes(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileToBytes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
