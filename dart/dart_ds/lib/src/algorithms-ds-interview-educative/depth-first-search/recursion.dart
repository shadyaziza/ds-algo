int factorial(int n) {
  if (n <= 1) return 1;
  return factorial(n - 1) * n;
}

int facotorialDetails(int n) {
  var stack = <int>[];
  //push each call to a stack
  //top of the stack is the base case
  while (n > 0) {
    stack.add(n);
    n -= 1;
  }
  var res = 1;
  //pop and use return value until stack is empty
  while (stack.isNotEmpty) {
    res *= stack.last;
    stack.removeLast();
  }
  return res;
}
