import 'package:dart_ds/src/algorithms-ds-interview-educative/binary-search/lesson6.dart';
import 'package:test/test.dart';

void main() {
  group('lesson 6 ... --->', () {
    test('case1', () {
      var tArr = [30, 40, 50, 10, 20];
      final result = lesson6(tArr);
      expect(result, 3);
    });
  });
}
