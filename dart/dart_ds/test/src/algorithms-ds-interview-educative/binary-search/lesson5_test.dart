import 'package:dart_ds/src/algorithms-ds-interview-educative/binary-search/lesson5.dart';
import 'package:test/test.dart';

void main() {
  group('lesson5 test cases.. --->', () {
    test('case 1', () {
      final tNumber = 16;
      final result = lesson5(tNumber);
      expect(result, 4);
    });
    test('case 2', () {
      final tNumber = 40;
      final result = lesson5(tNumber);
      expect(result, 6);
    });
  });
}
