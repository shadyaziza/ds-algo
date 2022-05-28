package leetcode

import (
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			6,
		},
		{
			"case 2",
			args{[]int{-2, 1}},
			1,
		},
		{
			"case 3",
			args{[]int{-2, -1}},
			-1,
		},
		{
			"case 4",
			args{[]int{1}},
			1,
		}, {
			"case 5",
			args{[]int{5, 4, -1, 7, 8}},
			23,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{1, 2},
			2,
		},
		{
			"case 2",
			args{-1, -2},
			-1,
		},
		{
			"case 3",
			args{-1, 1},
			1,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
