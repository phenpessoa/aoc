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

	doMagic(input)
	doMagic(example) // Expects: 13 && 1
}

func doMagic(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var (
		// Part 1
		visited = map[Pos]token{{0, 0}: {}}
		rope    = NewRope(1)

		// Part 2
		visited2 = map[Pos]token{{0, 0}: {}}
		rope2    = NewRope(9)
	)

	for scanner.Scan() {
		txt := scanner.Text()
		dir := txt[0]
		s := strings.Split(txt, " ")
		n, _ := strconv.Atoi(s[1])

		for i := 1; i < n+1; i++ {
			// part 1
			visited[rope.Move(dir)] = token{}

			// part 2
			visited2[rope2.Move(dir)] = token{}
		}
	}

	fmt.Println(len(visited))
	fmt.Println(len(visited2))
}

type token struct{}

type Pos struct{ x, y int }

type Rope struct {
	tail  *Rope
	pos   Pos
	knots int
}

func NewRope(knots int) *Rope {
	r := &Rope{knots: knots}
	cur := r
	for i := 0; i < knots; i++ {
		cur.tail = &Rope{knots: knots}
		cur = cur.tail
	}
	return r
}

func (r *Rope) Move(dir byte) Pos {
	switch dir {
	case 'R':
		r.pos.x++
	case 'L':
		r.pos.x--
	case 'U':
		r.pos.y++
	case 'D':
		r.pos.y--
	}

	cur := r
	for i := 0; i < r.knots; i++ {
		switch {
		case cur.tail.pos.x-cur.pos.x < -1:
			cur.tail.pos.x++
			switch {
			case cur.pos.y == cur.tail.pos.y:
			case cur.pos.y > cur.tail.pos.y:
				cur.tail.pos.y++
			case cur.pos.y < cur.tail.pos.y:
				cur.tail.pos.y--
			}
		case cur.tail.pos.x-cur.pos.x > 1:
			cur.tail.pos.x--
			switch {
			case cur.pos.y == cur.tail.pos.y:
			case cur.pos.y > cur.tail.pos.y:
				cur.tail.pos.y++
			case cur.pos.y < cur.tail.pos.y:
				cur.tail.pos.y--
			}
		}

		switch {
		case cur.tail.pos.y-cur.pos.y < -1:
			cur.tail.pos.y++
			switch {
			case cur.pos.x == cur.tail.pos.x:
			case cur.pos.x > cur.tail.pos.x:
				cur.tail.pos.x++
			case cur.pos.x < cur.tail.pos.x:
				cur.tail.pos.x--
			}
		case cur.tail.pos.y-cur.pos.y > 1:
			cur.tail.pos.y--
			switch {
			case cur.pos.x == cur.tail.pos.x:
			case cur.pos.x > cur.tail.pos.x:
				cur.tail.pos.x++
			case cur.pos.x < cur.tail.pos.x:
				cur.tail.pos.x--
			}
		}
		cur = cur.tail
	}
	return cur.pos
}
