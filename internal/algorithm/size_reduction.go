package algorithm

import (
	"fmt"
	"os"
)

func SizeReduction(file string, compressed string) error {
	info, err := os.Stat(file)
	if err != nil {
		return fmt.Errorf("error reading stats of file %s", file)
	}

	infoCompressed, err := os.Stat(compressed)
	if err != nil {
		return fmt.Errorf("error reading stats of compressed file %s", compressed)
	}

	oldSize := info.Size()
	newSize := infoCompressed.Size()

	diff := 100 - (float64(newSize) * 100 / float64(oldSize))

	fmt.Printf("Size reduced by %.2f%% (%v -> %v bytes)\n", diff, oldSize, newSize)

	return nil
}
