package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type dictionary []mapping

func (d dictionary) apply(n int) int {
	for _, m := range d {
		if n >= m.source && n < m.source+m.len {
			diff := n - m.source
			return m.dest + diff
		}
	}
	return n
}

type mapping struct {
	dest   int
	source int
	len    int
}

func parse(s string) mapping {
	var m mapping
	fmt.Sscanf(s, "%d %d %d", &m.dest, &m.source, &m.len)
	return m
}

func main() {
	scan := bufio.NewScanner(os.Stdin)

	var (
		seeds []int
		dict  dictionary
	)

	i := 0
	for scan.Scan() {
		line := scan.Text()
		i++

		// seeds
		if i == 1 {
			line = strings.TrimPrefix(line, "seeds: ")
			fields := strings.Fields(line)
			for _, f := range fields {
				n, _ := strconv.Atoi(f)
				seeds = append(seeds, n)
			}
			continue
		}

		// new dictionary
		if strings.Contains(line, ":") {
			dict = make(dictionary, 0)
			continue
		}

		// end of dictionary: apply to seeds
		if line == "" {
			for s := range seeds {
				seeds[s] = dict.apply(seeds[s])
			}
			continue
		}

		// mapping
		m := parse(line)
		dict = append(dict, m)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	// end of file: apply last set of mapping
	for s := range seeds {
		seeds[s] = dict.apply(seeds[s])
	}

	fmt.Println(seeds)

	min := math.MaxInt
	for _, s := range seeds {
		if s < min {
			min = s
		}
	}
	fmt.Println(min)
}
