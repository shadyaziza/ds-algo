//Problem statement
//A sorted array was rotated at an unknown pivot. For example,
//For example, [10, 20, 30, 40, 50] becomes [30, 40, 50, 10, 20].
//Find the index of the minimum element in this array.
int lesson6(List<int> arr) {
  var left = 0;
  var right = arr.length - 1;
  var ans = -1;
  var last = arr[right];
  while (left <= right) {
    var mid = (left + right) ~/ 2;
    if (arr[mid] <= last) {
      ans = mid;
      right = mid - 1;
    } else {
      left = mid + 1;
    }
  }
  return ans;
}
