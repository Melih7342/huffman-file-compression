package algorithm

import (
	"fmt"
	"os"
)

func SizeReduction(file string, compressed string) (float64, error) {
	info, err := os.Stat(file)
	if err != nil {
		return 0, fmt.Errorf("error reading stats of file %s", file)
	}

	infoCompressed, err := os.Stat(compressed)
	if err != nil {
		return 0, fmt.Errorf("error reading stats of compressed file %s", compressed)
	}

	oldSize := info.Size()
	newSize := infoCompressed.Size()

	if oldSize == 0 {
		return 0, nil
	}

	diff := 100 - (float64(newSize) * 100 / float64(oldSize))

	return diff, nil
}
