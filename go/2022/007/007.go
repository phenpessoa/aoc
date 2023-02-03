package main

import (
	"embed"
	"fmt"
	"io"
	"log"

	"github.com/Pedro-Pessoa/aoc/go/2022/007/terminal"
)

//go:embed input.txt
var Input embed.FS

//go:embed example.txt
var Example embed.FS

func main() {
	input, err := Input.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %s\n", err)
	}
	defer input.Close()

	example, err := Example.Open("example.txt")
	if err != nil {
		log.Fatalf("failed to open example file: %s\n", err)
	}
	defer example.Close()

	doMagic(input)
	doMagic(example) // Expects: 95437 && 24933642
}

func doMagic(r io.Reader) {
	term, err := terminal.Parse(r)
	if err != nil {
		log.Fatalf("failed to parse terminal: %#v", err)
	}

	const limit = 100000
	sum, err := term.SumOfSizesBelowLimit(limit)
	if err != nil {
		log.Fatalf("failed to calculate sum: %#v\n", err)
	}
	fmt.Println(sum)

	const (
		total     = 70000000
		needTotal = 30000000
	)

	size, _ := term.Size()
	needed := needTotal - (total - size)
	if needed <= 0 {
		fmt.Println("there is no need to delete any directory")
		return
	}

	smallest, err := term.SizeOfSmallestDirToFreeEnoughSpace(needed)
	if err != nil {
		log.Fatalf("failed to calculate smallest: %#v\n", err)
	}
	fmt.Println(smallest)
}
