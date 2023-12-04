package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func digit(s string) bool {
	return strings.Contains("0123456789", s)
}

type grid [][]string

func (g grid) cell(x, y int) string {
	if x < 0 || x >= len(g) || y < 0 || y >= len(g[0]) {
		return "."
	}
	return g[x][y]
}

func (g grid) adj(x, y int) bool {
	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			if dx == 0 && dy == 0 {
				continue
			}
			s := g.cell(x+dx, y+dy)
			if s != "." && !digit(s) {
				return true
			}
		}
	}
	return false
}

func parse(r io.Reader) grid {
	var g grid
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		split := strings.Split(line, "")
		g = append(g, split)
	}
	return g
}

func partnums(g grid) []int {
	var (
		nums  []int
		cur   string
		valid bool
	)
	for i := range g {
		for j := range g[i] {
			if digit(g[i][j]) {
				cur += g[i][j]
				if g.adj(i, j) {
					valid = true
				}
				continue
			}

			if len(cur) > 0 && valid {
				n, err := strconv.Atoi(cur)
				if err != nil {
					panic(err)
				}
				nums = append(nums, n)
			}
			cur = ""
			valid = false
			continue
		}
		if valid && len(cur) > 0 {
			n, err := strconv.Atoi(cur)
			if err != nil {
				panic(err)
			}
			nums = append(nums, n)
		}
		cur = ""
		valid = false
	}
	return nums
}

func main() {
	g := parse(os.Stdin)

	sum := 0
	for _, n := range partnums(g) {
		sum += n
	}
	fmt.Println(sum)
}
