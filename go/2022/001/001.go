package main

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"log"
	"strconv"

	"golang.org/x/exp/slices"
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

	doMagic(example) // Expects: 24000 && 45000
	fmt.Println()
	doMagic(input)
}

func doMagic(r io.Reader) {
	vals := make([]int, 0, 255)
	var cur int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			vals = append(vals, cur)
			cur = 0
			continue
		}
		cur += val
	}

	if cur != 0 {
		vals = append(vals, cur)
	}

	slices.Sort(vals)
	fmt.Println("First answer:", vals[len(vals)-1])
	fmt.Println("Second answer:", vals[len(vals)-1]+vals[len(vals)-2]+vals[len(vals)-3])
}
