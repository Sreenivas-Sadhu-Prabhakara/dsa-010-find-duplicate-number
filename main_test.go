package main

import "testing"

func TestFindDuplicate(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{1, 1}, 1},
		{[]int{2, 2, 2, 2, 2}, 2},
		{[]int{1, 4, 4, 2, 4}, 4},
	}
	for _, c := range cases {
		if got := FindDuplicate(c.in); got != c.want {
			t.Errorf("FindDuplicate(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
