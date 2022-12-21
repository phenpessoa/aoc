package main

import (
	"bufio"
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

	doMagic(example) // Expects: 157 && 70
	fmt.Println()
	doMagic(input)
}

func doMagic(r io.Reader) {
	scanner := bufio.NewScanner(r)
	var (
		sum1, sum2 int
		lc         int
		lines      = make([]string, 3)
	)
	for scanner.Scan() {
		txt := scanner.Text()
		c1 := txt[:len(txt)/2]
		c2 := txt[len(c1):]

		var bits uint64

		for _, r := range c1 {
			bits |= 1 << getRuneValue(r)
		}

		var same rune
		for _, r := range c2 {
			if bits&(1<<getRuneValue(r)) != 0 {
				same = r
				break
			}
		}

		sum1 += getRuneValue(same)

		if lc == 3 {
			lc = 0
		}

		lines[lc] = txt

		if lc == 2 {
			var bits1, bits2 uint64

			for _, r := range lines[0] {
				bits1 |= 1 << getRuneValue(r)
			}

			for _, r := range lines[1] {
				bits2 |= 1 << getRuneValue(r)
			}

			var same rune

			for _, r := range lines[2] {
				if bits1&(1<<getRuneValue(r)) != 0 &&
					bits2&(1<<getRuneValue(r)) != 0 {
					same = r
					break
				}
			}

			sum2 += getRuneValue(same)
		}

		lc++
	}
	fmt.Println("First answer:", sum1)
	fmt.Println("Second answer:", sum2)
}

func getRuneValue(r rune) int {
	if r <= 'Z' {
		return int(r - 'A' + 27)
	}
	return int(r - 'a' + 1)
}
