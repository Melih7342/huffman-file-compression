package test

import (
	"os"
	"strings"
)

func gen() {
	content := "Dies ist ein Test-Satz, der oft wiederholt wird.\n" +
		"Huffman sollte das sehr gut komprimieren können, " +
		"da die Buchstabenhäufigkeit extrem ungleich verteilt ist.\n"

	finalText := strings.Repeat(content, 5000)

	err := os.WriteFile("large_test.txt", []byte(finalText), 0644)
	if err != nil {
		panic(err)
	}
}
