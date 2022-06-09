// Problem statement
// A mountain array is defined as an array that:
// Has at least 3 elements.
// Has an element with the largest value called the “peak”, at an index k.
// The array elements monotonically increase from the first element to A[k],
// and then monotonically decreases from A[k + 1] to the last element of the array.
// Thus creating a “mountain” of numbers.
// Find the index of the peak element. Assume there are no duplicates.
int lesson7(List<int> arr) {
  var left = 0;
  var right = arr.length - 1;
  var ans = -1;
  while (left <= right) {
    var mid = (left + right) ~/ 2;
    if (mid == arr.length - 1 || arr[mid] >= arr[mid + 1]) {
      ans = mid;
      right = mid - 1;
    } else {
      left = mid + 1;
    }
  }
  return ans;
}
