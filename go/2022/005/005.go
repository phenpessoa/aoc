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

	doMagic(example) // Expects: CMZ && MCD
	fmt.Println()
	doMagic(input)
}

func doMagic(r io.Reader) {
	scanner := bufio.NewScanner(r)
	var (
		capturingCrates  = true
		stacks1, stacks2 [][]string
		crates           = make(map[int][]string)
	)

	for scanner.Scan() {
		txt := scanner.Text()

		if capturingCrates && strings.Contains(txt, "1") {
			capturingCrates = false
			continue
		}

		if capturingCrates {
			var (
				skip       int
				lastIdx    int
				isFirstIdx = true
			)
			for {
				if skip > len(txt) {
					break
				}

				idx := strings.Index(txt[skip:], "[")
				if idx == -1 {
					break
				}

				var crateIdx int
				if !isFirstIdx {
					crateIdx = (idx / 4) + lastIdx + 1
				} else {
					crateIdx = (idx / 4)
					isFirstIdx = false
				}

				crates[crateIdx] = append(crates[crateIdx], txt[skip+idx+1:skip+idx+2])

				skip += idx + 4
				lastIdx = crateIdx
			}
		}

		if !capturingCrates {
			if txt == "" {
				stacks1 = make([][]string, len(crates))
				for idx, c := range crates {
					stacks1[idx] = c
				}

				stacks2 = make([][]string, len(crates))
				for idx, c := range crates {
					stacks2[idx] = c
				}
				continue
			}

			blocks := strings.Split(txt, " ")
			amount, _ := strconv.Atoi(blocks[1])
			from, _ := strconv.Atoi(blocks[3])
			to, _ := strconv.Atoi(blocks[5])

			newStack := make([]string, 0, len(stacks1[from-1][:amount])+len(stacks1[to-1]))
			for i := len(stacks1[from-1][:amount]) - 1; i > -1; i-- {
				newStack = append(newStack, stacks1[from-1][i])
			}
			newStack = append(newStack, stacks1[to-1]...)
			stacks1[to-1] = newStack
			stacks1[from-1] = stacks1[from-1][amount:]

			newStack2 := make([]string, 0, len(newStack))
			newStack2 = append(newStack2, stacks2[from-1][:amount]...)
			newStack2 = append(newStack2, stacks2[to-1]...)
			stacks2[to-1] = newStack2
			stacks2[from-1] = stacks2[from-1][amount:]
		}
	}

	buf := strings.Builder{}
	buf.Grow(len(stacks1))
	for _, s := range stacks1 {
		buf.WriteString(s[0])
	}
	fmt.Println("First answer:", buf.String())

	buf.Reset()
	buf.Grow(len(stacks2))
	for _, s := range stacks2 {
		buf.WriteString(s[0])
	}
	fmt.Println("Second answer:", buf.String())
}
