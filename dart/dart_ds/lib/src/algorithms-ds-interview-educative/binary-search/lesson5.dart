//Problem Statement
//Given an integer, find its square root without using the built-in square root function
//. Only return the integer part (truncate the decimals).
int lesson5(int number) {
  if (number == 0) return number;
  var left = 1;
  var right = number;
  var res = -1;
  while (left <= right) {
    var mid = (left + right) ~/ 2;
    if (mid * mid <= number) {
      res = mid;
      left = mid + 1;
    } else {
      right = mid - 1;
    }
  }
  return res;
}
