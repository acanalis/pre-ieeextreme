package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type node struct {
	id  int
	nei []*node
	lat []int64
}

func solve(s *node, t int) string {
	visited := make(map[int]bool)
	dist := make(map[int]int64)
	var curr *node
	var queue []*node
	if s != nil {
		queue = append(queue, s)
		dist[s.id] = 0
	}
	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		if visited[curr.id] {
			continue
		}
		visited[curr.id] = true
		for i, n := range curr.nei {
			if !visited[n.id] {
				queue = append(queue, n)
			}
			d := dist[curr.id] + curr.lat[i]
			_, init := dist[n.id]
			if d < dist[n.id] || !init {
				dist[n.id] = d
			}
		}
	}
	d, ok := dist[t]
	if ok {
		return fmt.Sprint(d)
	}
	return "unreachable"
}

func main() {
	f, err := os.Open("ins.txt")
	if err != nil {
		f = os.Stdin
	}
	r := bufio.NewReader(f)

	tcase := NextInt(r)
	for t := 0; t < tcase; t++ {
		NextInt(r)
		m := NextInt(r)
		S := NextInt(r)
		T := NextInt(r)

		nmap := make(map[int]*node)
		for c := 0; c < m; c++ {
			ai := NextInt(r)
			a, ok := nmap[ai]
			if !ok {
				a = &node{id: ai}
			}
			bi := NextInt(r)
			b, ok := nmap[bi]
			if !ok {
				b = &node{id: bi}
			}
			a.nei = append(a.nei, b)
			b.nei = append(b.nei, a)
			latency := int64(NextInt(r))
			a.lat = append(a.lat, latency)
			b.lat = append(b.lat, latency)
			nmap[ai] = a
			nmap[bi] = b
		}
		fmt.Printf("Case #%d: %s\n", t+1, solve(nmap[S], T))
	}
}

func NextInt(r *bufio.Reader) (n int) {
	_, err := fmt.Fscan(r, &n)
	if err == nil || err == io.EOF {
		return n
	}
	panic(err)
}
