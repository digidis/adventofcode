package main

import (
	"fmt"
)

const (
	players = 428
	last    = 7082500
)

type marble struct {
	v    int
	next *marble
	prev *marble
}

func main() {
	var (
		current = &marble{0, nil, nil}
		scores  = make([]int, players+1)
		c       = 1
	)
	current.next = current
	current.prev = current
	for {
		for p := 1; p <= players; p++ {
			if c%23 == 0 {
				scores[p] += c
				current = current.prev.prev.prev.prev.prev.prev
				scores[p] += current.prev.v
				remove(current.prev)
			} else {
				current = insertAfter(current.next, c)
			}
			c++
			if c > last {
				m := 0
				for _, s := range scores {
					if s > m {
						m = s
					}
				}
				fmt.Printf("High score: %d\n", m)
				return
			}
		}
	}
}

func insertAfter(p *marble, v int) *marble {
	m := &marble{v: v, prev: p, next: p.next}
	p.next = m
	m.next.prev = m
	return m
}
func remove(m *marble) {
	m.prev.next = m.next
	m.next.prev = m.prev
}
