package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type race struct {
	time, record int
}

func main() {
	var lines []string
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
	if len(lines) != 2 {
		log.Fatalf("got %d lines, 2 expected", len(lines))
	}

	times, distances := parse(lines[0]), parse(lines[1])
	if len(times) != len(distances) {
		log.Fatalf("len(times) = %d, len(distances) = %d, expected same len", len(times), len(distances))
	}

	var races []race
	for i := range times {
		races = append(races, race{time: times[i], record: distances[i]})
	}

	var counts []int
	for _, r := range races {
		count := 0
		fmt.Println(r)
		for i := 0; i <= r.time; i++ {
			speed := i
			dist := speed * (r.time - i)
			fmt.Println(speed, dist)
			if dist > r.record {
				count++
			}
		}
		counts = append(counts, count)
	}
	fmt.Println(counts)

	result := 1
	for _, c := range counts {
		result *= c
	}
	fmt.Println(result)
}

func parse(s string) []int {
	var nums []int

	_, after, _ := strings.Cut(s, ":")
	s = strings.TrimSpace(after)
	fields := strings.Fields(s)

	for _, f := range fields {
		n, _ := strconv.Atoi(f)
		nums = append(nums, n)
	}

	return nums
}
