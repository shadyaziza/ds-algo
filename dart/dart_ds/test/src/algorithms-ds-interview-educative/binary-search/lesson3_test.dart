import 'package:dart_ds/src/algorithms-ds-interview-educative/binary-search/lesson3.dart';
import 'package:test/test.dart';

void main() {
  group("lesson3 test cases --> ", () {
    test("case 1", () {
      final tArr = [1, 3, 3, 5, 8, 8, 10];
      final result = lesson3(tArr, 2);
      expect(result, 1);
    });
  });
}
