package overview

import (
	"reflect"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	type args struct {
		arr    []any
		target any
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{[]any{1, 2, 3, 4, 5}, 1},
			0,
		},
		{
			"case 2",
			//add a new test case here with string
			args{[]any{"a", "b", "c", "d", "e"}, "a"},
			0,
		},
		{
			"case 3",
			args{[]any{1, 2, 3, 4, 5}, 6},
			-1,
		},
		{
			"case 4",
			//add failure case with string
			args{[]any{"a", "b", "c", "d", "e"}, "f"},
			-1,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := LinearSearch(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("LinearSearch() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, -3},
			2,
		}, {
			"case 2",
			args{[]int{-2, 1}, 1},
			1,
		}, {
			"case 3",
			args{[]int{-2, -1}, -1},
			1,
		},
		// failure case
		{
			"case 4",
			args{[]int{1}, 2},
			-1,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := BinarySearch(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case 1",
			args{[]int{3, 2, 5, 4, 1}},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"case 1",
			args{[]int{90, 56, 102, 300, 70}},
			[]int{56, 70, 90, 102, 300},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := BubbleSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case 1",
			args{[]int{3, 2, 5, 4, 1}},
			[]int{1, 2, 3, 4, 5},
		}, {
			"case 2",
			args{[]int{90, 56, 102, 300, 70}},
			[]int{56, 70, 90, 102, 300},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := InsertionSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
