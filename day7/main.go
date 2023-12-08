package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type typ int

const (
	five typ = iota
	four
	full
	three
	twoPair
	onePair
	high
)

func (t typ) String() string {
	return []string{"five", "four", "full", "three", "two pair", "one pair", "high"}[t]
}

func handTyp(hand string) typ {
	cards := make(map[string]int)
	for i := range hand {
		c := hand[i : i+1]
		cards[c]++
	}

	if len(cards) == 1 {
		return five
	}
	if len(cards) == 2 {
		for _, n := range cards {
			if n == 1 || n == 4 {
				return four
			}
			return full
		}
	}
	if len(cards) == 3 {
		for _, n := range cards {
			if n == 3 {
				return three
			}
			if n == 2 {
				return twoPair
			}
		}
	}
	if len(cards) == 4 {
		return onePair
	}
	if len(cards) == 5 {
		return high
	}
	panic("unreached")
}

func handStrength(c string) int {
	const order = "AKQJT98765432"
	s := strings.Index(order, c)
	if s == -1 {
		panic("unreached")
	}
	return s
}

type item struct {
	hand string
	bid  int
}

func parse(r io.Reader) ([]item, error) {
	var items []item

	scan := bufio.NewScanner(r)
	for scan.Scan() {
		line := scan.Text()
		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, fmt.Errorf("got %d fields, expected 2", len(fields))
		}

		hand := fields[0]
		bid, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, fmt.Errorf("could not parse bid: %v", err)
		}

		items = append(items, item{hand: hand, bid: bid})
	}
	if err := scan.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func main() {
	items, err := parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(items, func(i, j int) bool {
		typ1, typ2 := handTyp(items[i].hand), handTyp(items[j].hand)
		if typ1 < typ2 {
			return true
		}
		if typ1 > typ2 {
			return false
		}
		for k := 0; k < 5; k++ {
			s1, s2 := handStrength(items[i].hand[k:k+1]), handStrength(items[j].hand[k:k+1])
			if s1 < s2 {
				return true
			}
			if s1 > s2 {
				return false
			}
		}
		panic("unreached")
	})

	total := 0
	for i, t := range items {
		rank := len(items) - i
		win := t.bid * rank
		total += win
	}
	fmt.Println("total:", total)
}
