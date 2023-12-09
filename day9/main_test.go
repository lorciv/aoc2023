package main

import "testing"

func TestPredict(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{input: []int{0, 0, 0, 0}, want: 0},
		{input: []int{3, 3, 3, 3, 3}, want: 3},
		{input: []int{0, 3, 6, 9, 12, 15}, want: 18},
		{input: []int{1, 3, 6, 10, 15, 21}, want: 28},
		{input: []int{10, 13, 16, 21, 30, 45}, want: 68},
	}

	for _, test := range tests {
		got := predict(test.input)
		if got != test.want {
			t.Errorf("predict(%d) = %d, want %d", test.input, got, test.want)
		}
	}
}

func TestPostdict(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{input: []int{0, 0, 0, 0}, want: 0},
		{input: []int{0, 3, 6, 9, 12, 15}, want: -3},
		{input: []int{1, 3, 6, 10, 15, 21}, want: 0},
		{input: []int{10, 13, 16, 21, 30, 45}, want: 5},
	}

	for _, test := range tests {
		got := postdict(test.input)
		if got != test.want {
			t.Errorf("postdict(%d) = %d, want %d", test.input, got, test.want)
		}
	}
}
