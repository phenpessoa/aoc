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

	doMagic(example) // Expects: 15 && 12
	fmt.Println()
	doMagic(input)
}

func doMagic(r io.Reader) {
	var score1, score2 int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		txt := scanner.Text()
		g := [2]byte{txt[0], txt[2]}
		switch g[0] {
		case 'A':
			switch g[1] {
			case 'X':
				score1++    // points for choosing rock
				score1 += 3 // points for draw

				// x means I need to loose
				score2 += 3 // points for choosing schissors
			case 'Y':
				score1 += 2 // points for choosing paper
				score1 += 6 // points for winning

				// y means I need to draw
				score2 += 3 // points for draw
				score2++    // points for rock
			case 'Z':
				score1 += 3 // points for choosing scissors

				// z means I need to win
				score2 += 6 // points for winning
				score2 += 2 // points for choosing paper
			}
		case 'B':
			switch g[1] {
			case 'X':
				score1++ // points for choosing rock

				// x means I need to loose
				score2++ // points for choosing rock
			case 'Y':
				score1 += 2 // points for choosing paper
				score1 += 3 // points for draw

				// y means I need to draw
				score2 += 3 // points for draw
				score2 += 2 // points for choosing paper
			case 'Z':
				score1 += 3 // points for choosing scissors
				score1 += 6 // points for winning

				// z means I need to win
				score2 += 6 // points for winning
				score2 += 3 // points for choosing scissors
			}
		case 'C':
			switch g[1] {
			case 'X':
				score1++    // points for choosing rock
				score1 += 6 // points for winning

				// x means I need to loose
				score2 += 2 // points for choosing paper
			case 'Y':
				score1 += 2 // points for choosing paper

				// y means I need to draw
				score2 += 3 // points for draw
				score2 += 3 // points for choosing schissors
			case 'Z':
				score1 += 3 // points for choosing scissors
				score1 += 3 // points for draw

				// z means I need to win
				score2 += 6 // points for winning
				score2++    // points for choosing rock
			}
		}
	}

	fmt.Println("First Answer:", score1)
	fmt.Println("Second Answer:", score2)
}
