import 'package:dart_ds/src/algorithms-ds-interview-educative/binary-search/lesson7.dart';
import 'package:test/test.dart';

void main() {
  group('lesson7...--->', () {
    test('case1', () {
      final tArr = [0, 1, 2, 3, 2, 1, 0];
      final result = lesson7(tArr);
      expect(result, 3);
    });
  });
}
