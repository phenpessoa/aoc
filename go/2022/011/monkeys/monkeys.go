package monkeys

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"
)

func ParseMonkeys(r io.Reader) Monkeys {
	scanner := bufio.NewScanner(r)
	var (
		monkeys = make(Monkeys, 0)
		mID, i  int
	)
	for scanner.Scan() {
		txt := scanner.Text()
		switch i {
		case 0:
			const (
				text   = "Monkey "
				offset = len(text)
			)
			mID, _ = strconv.Atoi(txt[offset : offset+1])
		case 1:
			const (
				text   = "  Starting items:"
				offset = len(text)
			)
			items := strings.Split(txt[offset:], ",")
			its := make([]int, len(items))
			for i, it := range items {
				it = strings.TrimSpace(it)
				its[i], _ = strconv.Atoi(it)
			}
			monkeys = append(monkeys, Monkey{
				id:    mID,
				items: its,
			})
		case 2:
			const (
				text   = "  Operation: new = "
				offset = len(text)
			)
			data := strings.Split(txt[offset:], " ")
			op := Operation{signal: data[1]}
			if v, err := strconv.Atoi(data[2]); err == nil {
				op.xInt = v
			} else {
				op.x = data[2]
			}
			monkeys[mID].op = op
		case 3:
			const (
				text   = "  Test: divisible by "
				offset = len(text)
			)
			v, _ := strconv.Atoi(txt[offset:])
			monkeys[mID].div = v
		case 4:
			const (
				text   = "    If true: throw to monkey "
				offset = len(text)
			)
			v, _ := strconv.Atoi(txt[offset:])
			monkeys[mID].t = v
		case 5:
			const (
				text   = "    If false: throw to monkey "
				offset = len(text)
			)
			v, _ := strconv.Atoi(txt[offset:])
			monkeys[mID].f = v
		case 6:
			i = -1
		}
		i++
	}
	return monkeys
}

type Monkey struct {
	op      Operation
	items   []int
	id      int
	counter int
	div     int
	t, f    int
}

type Operation struct {
	x      string
	signal string
	xInt   int
}

type Monkeys []Monkey

func (ms Monkeys) copy() Monkeys {
	out := make(Monkeys, len(ms))
	for i, m := range ms {
		out[i] = Monkey{
			op:      m.op,
			items:   append([]int{}, m.items...),
			id:      m.id,
			counter: m.counter,
			div:     m.div,
			t:       m.t,
			f:       m.f,
		}
	}
	return out
}

func (ms Monkeys) LevelOfBusiness(rounds, div int) int {
	ms = ms.copy()

	globalDiv := 1
	for _, m := range ms {
		globalDiv *= m.div
	}
	for i := 0; i < rounds; i++ {
		for mID, m := range ms {
			for _, it := range m.items {
				ms[mID].counter++
				var v int
				switch m.op.signal {
				case "*":
					if m.op.x != "" {
						v = it * it
					} else {
						v = it * m.op.xInt
					}
				case "+":
					if m.op.x != "" {
						v = it + it
					} else {
						v = it + m.op.xInt
					}
				}
				if div > 0 {
					v /= div
				} else {
					v %= globalDiv
				}
				if v%m.div == 0 {
					ms[m.t].items = append(ms[m.t].items, v)
				} else {
					ms[m.f].items = append(ms[m.f].items, v)
				}
			}
			ms[mID].items = ms[mID].items[:0]
		}
	}
	slices.SortFunc(ms, func(a, b Monkey) int {
		if a.counter == b.counter {
			return 0
		}

		if a.counter < b.counter {
			return 1
		}

		return -1
	})
	return ms[0].counter * ms[1].counter
}
