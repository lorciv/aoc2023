package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func value(s string) int {
	var digits []rune
	for _, c := range s {
		if unicode.IsDigit(c) {
			digits = append(digits, c)
		}
	}
	first, _ := strconv.Atoi(string(digits[0]))
	last, _ := strconv.Atoi(string(digits[len(digits)-1]))
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
