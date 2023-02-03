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
	doMagic(example) // Expects: 13140 && __
}

func doMagic(r io.Reader) {
	scanner := bufio.NewScanner(r)
	var (
		sum int
		cpu = NewCPU()
		buf = &strings.Builder{}
	)
	for scanner.Scan() {
		txt := scanner.Text()
		blocks := strings.Split(txt, " ")
		cmd := blocks[0]
		switch cmd {
		case "noop":
			draw(cpu, buf)
			cpu.cycle++
			sum += verifyStrength(cpu)
		case "addx":
			val, _ := strconv.Atoi(blocks[1])
			for i := 0; i < 2; i++ {
				draw(cpu, buf)
				cpu.cycle++
				sum += verifyStrength(cpu)
			}
			cpu.x += val
		}
	}
	fmt.Println(sum)
	fmt.Println(buf.String())
}

func verifyStrength(cpu *CPU) int {
	switch cpu.cycle {
	case 20, 60, 100, 140, 180, 220:
		return cpu.x * int(cpu.cycle)
	default:
		return 0
	}
}

func draw(cpu *CPU, buf *strings.Builder) {
	re := int(cpu.cycle) % 40
	if re == 0 && cpu.cycle != 0 {
		buf.WriteString("\n")
	}
	switch cpu.x {
	case re, re + 1, re - 1:
		buf.WriteRune('#')
	default:
		buf.WriteRune('.')
	}
}

type CPU struct {
	x     int
	cycle uint
}

func NewCPU() *CPU {
	return &CPU{x: 1}
}
