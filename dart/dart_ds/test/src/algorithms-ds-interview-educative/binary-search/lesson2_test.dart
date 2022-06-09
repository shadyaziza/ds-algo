import 'package:dart_ds/src/algorithms-ds-interview-educative/binary-search/lesson2.dart';
import 'package:test/test.dart';

void main() {
  group('lesson2 test cases --> ', () {
    test('case 1', () async {
      final tArr = [false, false, true, true, true];
      final result = lesson2(tArr);
      expect(result, 2);
    });
    test('case 2', () async {
      final tArr = [false, false, false, true, true, true];
      final result = lesson2(tArr);
      expect(result, 3);
    });
  });
}
