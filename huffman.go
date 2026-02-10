package main

import (
	"slices"
	"sort"

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

func Huffman(bytes []byte) []byte {
	// Check byte frequencies
	frequencies := FrequencyCounter(bytes)

	// Convert the Map of character-values and frequencies to a list of nodes
	ConvertToNodeList(frequencies)
}
