package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type game struct {
	id    int
	hands []hand
}

type hand struct {
	red, green, blue int
}

func possible(g game, h hand) bool {
	m := hand{red: 0, green: 0, blue: 0}
	for _, h := range g.hands {
		if h.red > m.red {
			m.red = h.red
		}
		if h.green > m.green {
			m.green = h.green
		}
		if h.blue > m.blue {
			m.blue = h.blue
		}
	}
	return (h.red >= m.red) && (h.green >= m.green) && (h.blue >= m.blue)
}

func parseGame(s string) game {
	before, after, _ := strings.Cut(s, ":")
	var g game
	fmt.Sscanf(before, "Game %d", &g.id)
	split := strings.Split(after, ";")
	for _, p := range split {
		h := parseHand(p)
		g.hands = append(g.hands, h)
	}
	return g
}

func parseHand(s string) hand {
	var h hand

	split := strings.Split(s, ",")
	for _, s := range split {
		s = strings.TrimSpace(s)
		var (
			qty int
			typ string
		)
		fmt.Sscanf(s, "%d %s", &qty, &typ)
		switch typ {
		case "red":
			h.red = qty
		case "green":
			h.green = qty
		case "blue":
			h.blue = qty
		}
	}

	return h
}

func main() {
	var h = hand{red: 12, green: 13, blue: 14}

	sum := 0

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		g := parseGame(line)
		if possible(g, h) {
			sum += g.id
		}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
