package models

type HuffmanMetaData struct {
	Frequencies map[byte]int `json:"f"`
	ValidBits   int          `json:"v"`
}
