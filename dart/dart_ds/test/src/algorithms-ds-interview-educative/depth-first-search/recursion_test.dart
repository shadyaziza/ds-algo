import 'package:dart_ds/src/algorithms-ds-interview-educative/depth-first-search/recursion.dart';
import 'package:test/test.dart';

void main() {
  group('recursion test cases -->', () {
    test('case1', () {
      final tN = 4; // factorial of 4 is 24
      var result1 = factorial(4);
      var result2 = facotorialDetails(4);
      expect(result1, result2);
    });
  });
}
