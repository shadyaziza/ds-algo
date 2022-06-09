import 'package:dart_ds/src/algorithms-ds-interview-educative/binary-search/lesson1.dart';
import 'package:test/test.dart';

void main() {
  group('tests for lesson1 on binary search algo', () {
    test('lesson1 ...', () async {
      var arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11];
      final result = binarySearch(arr, 4);
      expect(result, equals(4));
    });
  });
}
