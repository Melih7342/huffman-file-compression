package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/Melih7342/huffman-file-compression/models"
)

func FrequencyCounter(bytes []byte) map[byte]int {
	frequencies := make(map[byte]int)
	for i := 0; i < len(bytes); i++ {
		if _, ok := frequencies[bytes[i]]; ok {
			frequencies[bytes[i]]++
		} else {
			frequencies[bytes[i]] = 1
		}
	}
	return frequencies
}

func ConvertToNodeList(m map[byte]int) []*models.Node {
	nodes := make([]*models.Node, 0, len(m))
	for value, frequency := range m {
		newNode := &models.Node{
			Value:     value,
			Frequency: frequency}
		nodes = append(nodes, newNode)
	}
	slices.SortFunc(nodes, func(a, b *models.Node) int {
		return a.Frequency - b.Frequency
	})
	return nodes
}

func BuildHuffmanTree(nodes []*models.Node) *models.Node {
	for len(nodes) > 1 {
		left := nodes[0]
		right := nodes[1]

		parent := &models.Node{
			Frequency: left.Frequency + right.Frequency,
			Left:      left,
			Right:     right,
		}
		nodes = nodes[2:]
		nodes = append(nodes, parent)

		slices.SortFunc(nodes, func(a, b *models.Node) int {
			return a.Frequency - b.Frequency
		})
	}
	return nodes[0]
}

func generateCodes(node *models.Node, currentPath string, codes map[byte]string) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		codes[node.Value] = currentPath
		return
	}
	generateCodes(node.Left, currentPath+"0", codes)
	generateCodes(node.Right, currentPath+"1", codes)
}

func ByteSliceToString(bytes []byte, codes map[byte]string) string {
	sb := &strings.Builder{}
	for _, b := range bytes {
		code := codes[b]
		sb.WriteString(code)
	}
	return sb.String()
}

func PackBits(bitString string) ([]byte, int) {
	// Determine the number of needed bytes
	output := make([]byte, 0, (len(bitString)+7)/8)

	var currentByte byte
	bitCount := 0

	for _, char := range bitString {
		// Bit shifting to make room
		currentByte <<= 1

		if char == '1' {
			currentByte |= 1
		}

		bitCount++

		// If the byte is full, append it to the byte slice
		if bitCount == 8 {
			output = append(output, currentByte)
			currentByte = 0
			bitCount = 0
		}
	}

	validCount := 8

	// If there is a rest, do padding and set the validCount
	// to the number of rest bits
	if bitCount > 0 {
		currentByte <<= 8 - bitCount
		output = append(output, currentByte)
		validCount = bitCount
	}
	return output, validCount
}

func SaveToFile(path string, metadata models.HuffmanMetaData, compressedData []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}

	defer file.Close()

	_, err = file.WriteString("HUFF")
	if err != nil {
		return err
	}

	metaBytes, err := json.Marshal(metadata)
	if err != nil {
		return fmt.Errorf("could not marshal metadata: %w", err)
	}

	metaLength := uint32(len(metaBytes))
	err = binary.Write(file, binary.LittleEndian, metaLength)
	if err != nil {
		return fmt.Errorf("could not write meta length: %w", err)
	}

	_, err = file.Write(metaBytes)
	if err != nil {
		return fmt.Errorf("could not write metadata: %w", err)
	}

	_, err = file.Write(compressedData)
	if err != nil {
		return fmt.Errorf("could not write compressed data: %w", err)
	}

	return nil
}

func Huffman(bytes []byte) []byte {
	// Check byte frequencies
	frequencies := FrequencyCounter(bytes)

	// Convert the Map of character-values and frequencies to a list of nodes
	nodeList := ConvertToNodeList(frequencies)

	// Build a huffman tree and save the root node
	rootNode := BuildHuffmanTree(nodeList)

	// Create a map for the paths of the node values
	codes := make(map[byte]string)
	generateCodes(rootNode, "", codes)

	// Generate a string, which is composed of the new bit representation of
	// the huffman coding map
	compressionString := ByteSliceToString(bytes, codes)

}
