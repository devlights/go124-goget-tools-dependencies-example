package main

import (
	"fmt"
	"slices"
)

//go:generate go tool stringer -type=Lang
type Lang int

const (
	Golang Lang = iota
	Java
	CSharp
	C
	CPP
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var (
		items = []Lang{Java, Golang, CSharp}
	)
	for item := range slices.Values(items) {
		fmt.Printf("%6[1]s(%[1]d)\n", item)
	}

	return nil
}
