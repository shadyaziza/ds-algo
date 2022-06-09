//Given an array of integers sorted in increasing order and a target,
// find the index of the first element in the array that is larger or equal to the target.
// Assume that it is guaranteed to find a satisfying number.

int lesson3(List<int> arr, int target) {
  var left = 0;
  var right = arr.length - 1;
  int boundaryIndex = -1;
  while (left <= right) {
    var mid = (left + right) ~/ 2;
    if (arr[mid] >= target) {
      boundaryIndex = mid;
      right = mid - 1;
    } else {
      left = mid + 1;
    }
  }
  return boundaryIndex;
}
