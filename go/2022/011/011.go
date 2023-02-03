package main

import (
	"embed"
	"fmt"
	"io"
	"log"

	"github.com/Pedro-Pessoa/aoc/go/2022/011/monkeys"
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
	doMagic(example) // Expects: 10605 && 2713310158
}

func doMagic(r io.Reader) {
	ms := monkeys.ParseMonkeys(r)
	fmt.Println(ms.LevelOfBusiness(20, 3))
	fmt.Println(ms.LevelOfBusiness(10000, 0))
}
