package algorithm

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Melih7342/huffman-file-compression/internal/models"
)

func DecompressFile(filePath string, targetPath string) error {
	// Convert file to byte slice
	compressedFile, err := FileToBytes(filePath)
	fmt.Printf("Converting the file %s to bytes\n", filePath)
	if err != nil {
		return fmt.Errorf("error converting the file %s to bytes", filePath)
	}

	// Check if the file to unpack is a HUFF file
	marker := string(compressedFile[:4])
	fmt.Println("checking the marker...")

	if marker != "HUFF" {
		return fmt.Errorf("invalid compressed file marker")
	}

	// Convert the 4. - 8. bytes into a metadata struct instance
	metaDataLength := binary.LittleEndian.Uint32(compressedFile[4:8])
	var metaData models.HuffmanMetaData
	err = json.Unmarshal(compressedFile[8:8+metaDataLength], &metaData)
	fmt.Println("checking the metadata...")
	if err != nil {
		return fmt.Errorf("error unmarshalling the metadata")
	}

	// Rebuild the tree
	root := BuildHuffmanTree(ConvertToNodeList(metaData.Frequencies))
	fmt.Println("building the huffman tree...")

	// Unpack into target file
	dataStart := 8 + int(metaDataLength)
	decodedData := make([]byte, 0, len(compressedFile)*2)
	wanderer := root
	// loop through the content and write them with their values into the file
	fmt.Println("decompressing the content...")
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
	fmt.Printf("Writing the decompressed content into %s...", targetPath)
	err = os.WriteFile(targetPath, decodedData, 0644)
	if err != nil {
		return fmt.Errorf("could not write unpacked file: %w", err)
	}
	return nil
}
