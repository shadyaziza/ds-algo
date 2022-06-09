import 'package:dart_ds/src/algorithms-ds-interview-educative/binary-search/lesson4.dart';
import 'package:test/test.dart';

void main() {
  group('lesson4 ... --->', () {
    test('case 1', () {
      final tArr = [1, 3, 3, 3, 3, 6, 10, 10, 10, 100];
      final result = lesson4(tArr, 3);
      expect(result, 1);
    });
  });
}
