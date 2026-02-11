package algorithm

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Melih7342/huffman-file-compression/internal/models"
)

func DecompressFile(path string) error {
	// Convert file to byte slice
	compressedFile, err := FileToBytes(path)
	if err != nil {
		return err
	}

	// Check if the file to unpack is a HUFF file
	marker := string(compressedFile[:4])

	if marker != "HUFF" {
		return fmt.Errorf("invalid compressed file marker")
	}

	// Convert the 4. - 8. bytes into a metadata struct instance
	metaDataLength := binary.LittleEndian.Uint32(compressedFile[4:8])
	var metaData models.HuffmanMetaData
	err = json.Unmarshal(compressedFile[8:8+metaDataLength], &metaData)
	if err != nil {
		return err
	}

	// Rebuild the tree
	root := BuildHuffmanTree(ConvertToNodeList(metaData.Frequencies))

	// Unpack into target file
	dataStart := 8 + int(metaDataLength)
	decodedData := make([]byte, 0, len(compressedFile)*2)
	wanderer := root
	// loop through the content and write them with their values into the file
	for i := dataStart; i < len(compressedFile); i++ {
		currentByte := compressedFile[i]

		limit := 8
		if i == len(compressedFile)-1 {
			limit = metaData.ValidBits
		}

		for bit := 7; bit >= (8 - limit); bit-- {
			isBitSet := (currentByte >> bit) & 1

			if isBitSet == 1 {
				wanderer = wanderer.Right
			} else {
				wanderer = wanderer.Left
			}

			if wanderer.Left == nil && wanderer.Right == nil {
				decodedData = append(decodedData, wanderer.Value)
				wanderer = root
			}
		}
	}
	newPath := strings.TrimSuffix(path, ".huff")
	err = os.WriteFile(newPath, decodedData, 0644)
	if err != nil {
		return fmt.Errorf("could not write unpacked file: %w", err)
	}
	return nil
}
