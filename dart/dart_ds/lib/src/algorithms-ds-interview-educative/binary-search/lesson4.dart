//Problem Statement
//Given a sorted array of integers and a target integer, find the first occurrence of the target and return its index.
// Return -1 if the target is not in the array.

int lesson4(List<int> arr, int target) {
  var left = 0;
  var right = arr.length - 1;
  var ans = -1;
  while (left <= right) {
    var mid = (left + right) ~/ 2;
    if (target == arr[mid]) {
      ans = mid;
      right = mid - 1;
    } else if (target < arr[mid]) {
      right = mid - 1;
    } else {
      left = mid + 1;
    }
  }
  return ans;
}
