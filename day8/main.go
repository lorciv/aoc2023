package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	steps []string
	index int
)

func next() string {
	s := steps[index]
	index = (index + 1) % len(steps)
	return s
}

type node struct {
	left, right string
}

var nodes = make(map[string]node)

func main() {
	i := 0
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		i++

		// parse instructions
		if i == 1 {
			steps = strings.Split(line, "")
			continue
		}

		// skip white line
		if line == "" {
			continue
		}

		// parse network
		line = strings.ReplaceAll(line, "=", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ")", "")
		fields := strings.Fields(line)

		nodes[fields[0]] = node{left: fields[1], right: fields[2]}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	cur := "AAA"
	steps := 0

	for cur != "ZZZ" {
		n := nodes[cur]
		x := next()

		if x == "L" {
			cur = n.left
		}
		if x == "R" {
			cur = n.right
		}

		steps++
	}

	fmt.Println(steps)
}
