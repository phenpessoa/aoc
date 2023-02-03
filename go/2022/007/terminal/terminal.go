package terminal

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type TerminalType uint

const (
	TerminalTypeFile TerminalType = iota
	TerminalTypeDir
)

type Terminal struct {
	children map[string]*Terminal
	parent   *Terminal
	name     string
	typ      TerminalType
	size     int
}

func Parse(r io.Reader) (*Terminal, error) {
	scanner := bufio.NewScanner(r)
	var (
		base = &Terminal{
			children: make(map[string]*Terminal),
			typ:      TerminalTypeDir,
		}
		cur = base
	)
outer:
	for scanner.Scan() {
		txt := scanner.Text()

		switch {
		case strings.HasPrefix(txt, "$"):
			blocks := strings.Split(txt, " ")
			switch cmd := blocks[1]; cmd {
			case "ls":
				continue outer
			case "cd":
				switch arg := blocks[2]; arg {
				case "..":
					cur = cur.parent
				case "/":
					cur = base
				default:
					cur = cur.children[arg]
				}
			default:
				return nil, errors.New("unknown command: " + cmd)
			}
		case strings.HasPrefix(txt, "dir"):
			name := txt[len("dir "):]
			cur.children[name] = &Terminal{
				children: make(map[string]*Terminal),
				parent:   cur,
				typ:      TerminalTypeDir,
				name:     name,
			}
		case strings.ContainsAny(string(txt[0]), "123456789"):
			blocks := strings.Split(txt, " ")
			size, _ := strconv.Atoi(blocks[0])
			name := blocks[1]
			cur.children[name] = &Terminal{
				parent: cur,
				typ:    TerminalTypeFile,
				name:   name,
				size:   size,
			}
		default:
			return nil, errors.New("unknown terminal string format: " + txt)
		}
	}

	return base, nil
}

func (it *Terminal) Size() (int, error) {
	var sink int
	size, err := it.calcSumOfSizesBelowLimit(1<<63-1, &sink)
	if err != nil {
		return -1, fmt.Errorf("failed to calculate size: %w", err)
	}
	return size, nil
}

func (it *Terminal) SumOfSizesBelowLimit(limit int) (int, error) {
	var out int
	if _, err := it.calcSumOfSizesBelowLimit(limit, &out); err != nil {
		return -1, fmt.Errorf("failed to calculate sum: %w", err)
	}
	return out, nil
}

func (it *Terminal) calcSumOfSizesBelowLimit(limit int, out *int) (int, error) {
	if it == nil {
		return 0, nil
	}

	switch it.typ {
	case TerminalTypeFile:
		return it.size, nil
	case TerminalTypeDir:
		var total int
		for _, c := range it.children {
			val, err := c.calcSumOfSizesBelowLimit(limit, out)
			if err != nil {
				return -1, err
			}
			total += val
		}

		if total < limit {
			*out += total
		}

		return total, nil
	default:
		return -1, errors.New("unknown item type: " + strconv.Itoa(int(it.typ)))
	}
}

func (it *Terminal) SizeOfSmallestDirToFreeEnoughSpace(needed int) (int, error) {
	var out int
	if err := it.findSizeOfSmallestDirToFreeEnoughSpace(needed, &out); err != nil {
		return -1, fmt.Errorf("failed to calculate smallest dir: %w", err)
	}
	return out, nil
}

func (it *Terminal) findSizeOfSmallestDirToFreeEnoughSpace(needed int, out *int) error {
	if it == nil {
		return nil
	}

	switch it.typ {
	case TerminalTypeFile:
		return nil
	case TerminalTypeDir:
		s, err := it.Size()
		if err != nil {
			return err
		}
		if s < needed {
			return nil
		}

		if r := *out; r == 0 || s < r {
			*out = s
		}

		for _, c := range it.children {
			c.findSizeOfSmallestDirToFreeEnoughSpace(needed, out)
		}
	default:
		return errors.New("unknown item type: " + strconv.Itoa(int(it.typ)))
	}

	return nil
}
