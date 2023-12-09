package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func predict(nums []int) int {
	zeros := true
	for _, n := range nums {
		if n != 0 {
			zeros = false
			break
		}
	}
	if zeros {
		return 0
	}

	diff := make([]int, len(nums)-1)
	for i := 0; i < len(diff); i++ {
		diff[i] = nums[i+1] - nums[i]
	}
	return nums[len(nums)-1] + predict(diff)
}

func parse(s string) ([]int, error) {
	var nums []int
	fields := strings.Fields(s)
	for _, f := range fields {
		n, err := strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func main() {
	sum := 0

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()

		nums, err := parse(line)
		if err != nil {
			log.Fatal(err)
		}
		sum += predict(nums)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
