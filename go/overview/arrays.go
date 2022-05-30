package overview

// Searching

//LinearSearch:
//Time: Worst: O(n), Best: O(1)
//Space: O(1)

func LinearSearch(arr []any, target any) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1

}

//BinarySearch
//A binary search or half-interval search is a search algorithm that works by
// comparing the target value to the middle element of a sorted array.
//Time: Worst: O(log n), Best: O(1)
//Space: O(1)

func BinarySearch(arr []int, target int) int {
	// TODO: Test with negative numbers
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// Sorting

// Bubble Sort
// In bubble sort, the largest element bubbles to the top of the array.
//Time: Worst: O(n^2), Best: O(n)
//Space: O(1)

func BubbleSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}
