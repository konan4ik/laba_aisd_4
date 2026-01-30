package str_search

import (
	"container/list"
	"fmt"
	p "laba4/parsing"
	r "laba4/recs"
	"sync"
	"time"
)

type node struct {
	root bool

	b []byte

	output bool

	index int

	counter uint64

	child [256]*node

	fails [256]*node

	suffix *node

	fail *node
}

type Matcher struct {
	counter uint64
	trie    []node
	extent  int
	root    *node

	heap sync.Pool
}

func (m *Matcher) findBlice(b []byte) *node {
	n := &m.trie[0]

	for n != nil && len(b) > 0 {
		n = n.child[int(b[0])]
		b = b[1:]
	}

	return n
}

func (m *Matcher) getFreeNode() *node {
	m.extent += 1

	if m.extent == 1 {
		m.root = &m.trie[0]
		m.root.root = true
	}

	return &m.trie[m.extent-1]
}

func (m *Matcher) buildTrie(dictionary [][]byte) {

	max := 1
	for _, blice := range dictionary {
		max += len(blice)
	}
	m.trie = make([]node, max)

	m.getFreeNode()

	for i, blice := range dictionary {
		n := m.root
		var path []byte
		for _, b := range blice {
			path = append(path, b)

			c := n.child[int(b)]

			if c == nil {
				c = m.getFreeNode()
				n.child[int(b)] = c
				c.b = make([]byte, len(path))
				copy(c.b, path)

				if len(path) == 1 {
					c.fail = m.root
				}

				c.suffix = m.root
			}

			n = c
		}

		n.output = true
		n.index = i
	}

	l := new(list.List)
	l.PushBack(m.root)

	for l.Len() > 0 {
		n := l.Remove(l.Front()).(*node)

		for i := 0; i < 256; i++ {
			c := n.child[i]
			if c != nil {
				l.PushBack(c)

				for j := 1; j < len(c.b); j++ {
					c.fail = m.findBlice(c.b[j:])
					if c.fail != nil {
						break
					}
				}

				if c.fail == nil {
					c.fail = m.root
				}

				for j := 1; j < len(c.b); j++ {
					s := m.findBlice(c.b[j:])
					if s != nil && s.output {
						c.suffix = s
						break
					}
				}
			}
		}
	}

	for i := 0; i < m.extent; i++ {
		for c := 0; c < 256; c++ {
			n := &m.trie[i]
			for n.child[c] == nil && !n.root {
				n = n.fail
			}

			m.trie[i].fails[c] = n
		}
	}

	m.trie = m.trie[:m.extent]
}

func NewMatcher(dictionary [][]byte) *Matcher {
	m := new(Matcher)

	m.buildTrie(dictionary)

	return m
}

func NewStringMatcher(dictionary []string) *Matcher {
	m := new(Matcher)

	var d [][]byte
	for _, s := range dictionary {
		d = append(d, []byte(s))
	}

	m.buildTrie(d)

	return m
}

func (m *Matcher) Contains(in []byte) bool {
	n := m.root
	for _, b := range in {
		c := int(b)
		if !n.root {
			n = n.fails[c]
		}

		if n.child[c] != nil {
			f := n.child[c]
			n = f

			if f.output {
				return true
			}

			for !f.suffix.root {
				return true
			}
		}
	}
	return false
}

func CorasickTimed(fDict [][]byte, dDict [][]byte, arr []r.Record, n int) {
	start := time.Now()

	var output []string
	output = append(output, "Строка   Data\n")

	f := NewMatcher(fDict)
	d := NewMatcher(dDict)
	for i := 0; i < n; i++ {
		rec := arr[i]

		fullname := rec.FullName.Name + " " + rec.FullName.SurName + " " + rec.FullName.Otchestvo

		fullMatched := f.Contains([]byte(fullname))
		descrMatched := d.Contains([]byte(rec.Descrp))

		if fullMatched && descrMatched {
			result := fmt.Sprintf(
				"%-8d %-30s   %-55s \n", rec.Number, fullname, rec.Descrp,
			)
			output = append(output, result)
		}
	}

	elapsed := time.Since(start)

	output = append(output, fmt.Sprintf("Времени затрачено: %s", elapsed.String()))

	p.FillFile(output, "output/Corasick.txt", len(output))
}
