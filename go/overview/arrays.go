package overview

// Searching

//LinearSearch Linear Search
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

//BinarySearch Binary Search
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

//BubbleSort Bubble Sort
//In bubble sort, the largest element bubbles to the top of the array.
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

//InsertionSort Insertion Sort
//In insertion sort, the smallest element is inserted into the sorted array.
//Time: Worst: O(n^2), Best: O(n)
//Space: O(1)
func InsertionSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}

	return arr
}

//SelectionSort Selection Sort
//In selection sort, the smallest element is selected and inserted into the
//sorted array.
//Time: Worst: O(n^2), Best: O(n^2)
//Space: O(1)
func SelectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}

	return arr
}

//QuickSort Quick Sort
//In quick sort, the pivot element is selected and the array is partitioned
//around it.
//Time: Worst: O(n^2), Best: O(n log n)
//Space: O(log n)
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[0]
	var left []int
	var right []int
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}

//MergeSort Merge Sort
//In merge sort, the array is split into two parts and the two parts are
//merged.
//Time: Worst: O(n log n), Best: O(n log n)
//Space: O(n)
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(sort []int, sort2 []int) []int {
	var result []int
	for len(sort) > 0 && len(sort2) > 0 {
		if sort[0] < sort2[0] {
			result = append(result, sort[0])
			sort = sort[1:]
		} else {
			result = append(result, sort2[0])
			sort2 = sort2[1:]
		}
	}
	if len(sort) > 0 {
		result = append(result, sort...)
	}
	if len(sort2) > 0 {
		result = append(result, sort2...)
	}
	return result
}
