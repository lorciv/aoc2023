package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var grid [][]string

type cell struct {
	x, y int
}

func emptyRow(x int) bool {
	for j := 0; j < len(grid[x]); j++ {
		if grid[x][j] == "#" {
			return false
		}
	}
	return true
}

func emptyCol(y int) bool {
	for i := 0; i < len(grid); i++ {
		if grid[i][y] == "#" {
			return false
		}
	}
	return true
}

func galaxies() []cell {
	var g []cell
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "#" {
				g = append(g, cell{i, j})
			}
		}
	}
	return g
}

func dist(c, d cell, exp int) int {
	res := 0

	src, end := c.x, d.x
	if src > end {
		src, end = end, src
	}
	for i := src; i < end; i++ {
		if emptyRow(i) {
			res += exp
		} else {
			res++
		}
	}

	src, end = c.y, d.y
	if src > end {
		src, end = end, src
	}
	for j := src; j < end; j++ {
		if emptyCol(j) {
			res += exp
		} else {
			res++
		}
	}

	return res
}

var expansionFactor = flag.Int("e", 2, "expansion factor for empty rows/columns")

func main() {
	flag.Parse()

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		row := strings.Split(line, "")
		grid = append(grid, row)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	gal := galaxies()
	for i := 0; i < len(gal); i++ {
		for j := i + 1; j < len(gal); j++ {
			// expansion factor == 2 for part 1, 1000000 for part 2
			d := dist(gal[i], gal[j], *expansionFactor)
			sum += d
		}
	}
	fmt.Println(sum)
}
