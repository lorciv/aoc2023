package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type card struct {
	id          int
	actual, win []int
}

func points(c card) int {
	p := 0
	for _, a := range c.actual {
		for _, w := range c.win {
			if a == w && p == 0 {
				p = 1
			} else if a == w && p > 0 {
				p *= 2
			}
		}
	}
	return p
}

func parse(s string) card {
	var c card
	before, after, _ := strings.Cut(s, ":")
	fmt.Sscanf(before, "Card %d", &c.id)

	actual, win, _ := strings.Cut(after, "|")
	c.actual = parseNumbers(actual)
	c.win = parseNumbers(win)

	return c
}

func parseNumbers(s string) []int {
	var nums []int

	s = strings.TrimSpace(s)
	fields := strings.Fields(s)
	for _, f := range fields {
		n, _ := strconv.Atoi(f)
		nums = append(nums, n)
	}

	sort.Ints(nums)

	return nums
}

func main() {
	tot := 0

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		c := parse(line)
		tot += points(c)
	}

	fmt.Println(tot)
}
