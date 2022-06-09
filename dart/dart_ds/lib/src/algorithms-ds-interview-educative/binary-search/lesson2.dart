int lesson2(List<bool> arr) {
  var left = 0;
  var right = arr.length - 1;
  int boundaryIndex = -1;
  while (left <= right) {
    int mid = (left + right) ~/ 2;
    if (!arr[mid]) {
      left = mid + 1;
    } else {
      boundaryIndex = mid;
      right = mid - 1;
    }
  }
  return boundaryIndex;
}
