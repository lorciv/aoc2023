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
	if zero(nums) {
		return 0
	}

	d := diff(nums)
	return nums[len(nums)-1] + predict(d)
}

func postdict(nums []int) int {
	if zero(nums) {
		return 0
	}

	d := diff(nums)
	return nums[0] - postdict(d)
}

func diff(nums []int) []int {
	d := make([]int, len(nums)-1)
	for i := 0; i < len(d); i++ {
		d[i] = nums[i+1] - nums[i]
	}
	return d
}

func zero(nums []int) bool {
	z := true
	for _, n := range nums {
		if n != 0 {
			z = false
			break
		}
	}
	return z
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
	presum, postsum := 0, 0

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()

		nums, err := parse(line)
		if err != nil {
			log.Fatal(err)
		}
		presum += predict(nums)
		postsum += postdict(nums)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(presum, postsum)
}
