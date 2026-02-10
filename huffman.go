package main

import (
	"slices"

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

}
