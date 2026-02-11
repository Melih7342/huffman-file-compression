package models

type Node struct {
	Value       byte
	Frequency   int
	Left, Right *Node
}
