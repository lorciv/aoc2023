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
	count       int
	id          int
	actual, win []int
}

func matches(c card) int {
	m := 0
	for _, a := range c.actual {
		for _, w := range c.win {
			if a == w {
				m++
			}
		}
	}
	return m
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
	c := card{
		count: 1,
	}

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
	var cards []card

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		c := parse(line)
		cards = append(cards, c)
	}

	// Part 1
	totPoints := 0
	for _, c := range cards {
		totPoints += points(c)
	}
	fmt.Println("points", totPoints)

	// Part 2
	for i, c := range cards {
		m := matches(c)
		for j := 0; j < m; j++ {
			cards[i+j+1].count += c.count
		}
	}
	totCards := 0
	for _, c := range cards {
		totCards += c.count
	}
	fmt.Println("cards", totCards)
}
