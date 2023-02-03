package main

import (
	"embed"
	"fmt"
	"io"
	"log"
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
	doMagic(example) // Expects: 7 && 19
}

func doMagic(r io.Reader) {
	data, err := io.ReadAll(r)
	if err != nil {
		log.Fatalf("failed to read all: %#v\n", err)
	}

	solve(data, 4)
	solve(data, 14)
}

func solve(data []byte, n int) {
outer:
	for i := 0; i < len(data); i++ {
		var bits uint64
		for j := 0; j < n; j++ {
			v := data[i+j] - 'a'

			if bits&(1<<v) != 0 {
				continue outer
			}

			bits |= 1 << v
		}

		// if we got here, all chars were different
		fmt.Println(i + n)
		return
	}
}
