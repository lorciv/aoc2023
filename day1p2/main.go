package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var words = map[string]int{
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func digits(s string) []int {
	if len(s) == 0 {
		return nil
	}
	for word, val := range words {
		if strings.HasPrefix(s, word) {
			tail := digits(s[1:])
			return append([]int{val}, tail...)
		}
	}
	return digits(s[1:])
}

func value(s string) int {
	d := digits(s)
	first := d[0]
	last := d[len(d)-1]
	return 10*first + last
}

func main() {
	sum := 0

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		sum += value(line)
	}
	if err := scan.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(sum)
}
