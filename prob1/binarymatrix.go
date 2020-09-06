package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("ins.txt")
	if err != nil {
		f = os.Stdin
	}
	r := bufio.NewReader(f)

	N := NextInt(r)
	M := NextInt(r)

	var sum int

	count := make([]int, M)
	count[0] = N // Flip all left most rows

	for n := 0; n < N; n++ {
		flip := NextInt(r)
		for m := 1; m < M; m++ {
			i := NextInt(r)
			if i == flip {
				count[m] += 1
			}
		}
	}
	for m := 0; m < M; m++ {
		if count[m] <= N/2 {
			count[m] = N - count[m]
		}
		sum += (1 << (M - m - 1)) * count[m]
	}
	fmt.Println(sum)
}

func NextInt(r *bufio.Reader) (n int) {
	_, err := fmt.Fscan(r, &n)
	if err == nil || err == io.EOF {
		return n
	}
	panic(err)
}
