package main

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"log"
	"strconv"
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
	doMagic(example) // Expects: 21 && 8
}

func doMagic(r io.Reader) {
	scanner := bufio.NewScanner(r)
	trees := make([][]int, 0)
	for scanner.Scan() {
		txt := scanner.Text()
		l := make([]int, len(txt))
		for i, c := range txt {
			tree, _ := strconv.Atoi(string(c))
			l[i] = tree
		}
		trees = append(trees, l)
	}

	rows := len(trees)
	cols := len(trees[0])
	visible := make(map[treePos]token, 0)

	for i := 0; i < rows; i++ {
		highest := -1
		for j := 0; j < cols; j++ {
			if tree := trees[i][j]; tree > highest {
				highest = tree
				visible[treePos{i, j}] = token{}
			}
		}

		highest = -1
		for j := cols - 1; j > -1; j-- {
			if tree := trees[i][j]; tree > highest {
				highest = tree
				visible[treePos{i, j}] = token{}
			}
		}
	}

	for i := 0; i < cols; i++ {
		highest := -1
		for j := 0; j < rows; j++ {
			if tree := trees[j][i]; tree > highest {
				highest = tree
				visible[treePos{j, i}] = token{}
			}
		}

		highest = -1
		for j := rows - 1; j > -1; j-- {
			if tree := trees[j][i]; tree > highest {
				highest = tree
				visible[treePos{j, i}] = token{}
			}
		}
	}

	fmt.Println(len(visible))

	// part 2
	var highest int
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			s := new(Score)
			cur := trees[i][j]
			for k := i + 1; k < rows; k++ {
				s.a++
				if tree := trees[k][j]; tree >= cur {
					break
				}
			}

			for k := i - 1; k > -1; k-- {
				s.b++
				if tree := trees[k][j]; tree >= cur {
					break
				}
			}

			for k := j + 1; k < cols; k++ {
				s.c++
				if tree := trees[i][k]; tree >= cur {
					break
				}
			}

			for k := j - 1; k > -1; k-- {
				s.d++
				if tree := trees[i][k]; tree >= cur {
					break
				}
			}

			if total := s.Total(); total > highest {
				highest = total
			}
		}
	}

	fmt.Println(highest)
}

type token struct{}

type treePos struct{ x, y int }

type Score struct {
	a, b, c, d int
}

func (s *Score) Total() int {
	return s.a * s.b * s.c * s.d
}
