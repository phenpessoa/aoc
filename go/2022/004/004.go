package main

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
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

	doMagic(example) // Expects: 2 && 4
	fmt.Println()
	doMagic(input)
}

func doMagic(r io.Reader) {
	scanner := bufio.NewScanner(r)
	var sum1, sum2 int

	for scanner.Scan() {
		txt := scanner.Text()
		pairs := strings.Split(txt, ",")

		firstSections := strings.Split(pairs[0], "-")
		secondSections := strings.Split(pairs[1], "-")

		fs1, _ := strconv.Atoi(firstSections[0])
		fs2, _ := strconv.Atoi(firstSections[1])
		ss1, _ := strconv.Atoi(secondSections[0])
		ss2, _ := strconv.Atoi(secondSections[1])

		if (fs1 <= ss1 && fs2 >= ss2) ||
			(ss1 <= fs1 && ss2 >= fs2) {
			sum1++
		}

		if !(ss1 > fs2 || ss2 < fs1) {
			sum2++
		}
	}

	fmt.Println("First answer:", sum1)
	fmt.Println("Second answer:", sum2)
}
