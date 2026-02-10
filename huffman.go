package main

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

func Huffman(bytes []byte) []byte {
	// Check byte frequencies
	frequencies := FrequencyCounter(bytes)

}
