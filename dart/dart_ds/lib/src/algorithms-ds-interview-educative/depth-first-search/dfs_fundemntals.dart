class Node {
  Node? left;
  Node? right;
  final int value;

  Node({required this.value, this.left, this.right});

  static Node? buildTree(List<int> data) {
    if (data.isEmpty) {
      return null;
    }
    final root = Node(value: data[0]);
    for (int i = 1; i < data.length; i++) {
      root.insert(data[i]);
    }
    return root;
  }

  void insert(int data) {
    if (data < value) {
      if (left == null) {
        left = Node(value: data);
      } else {
        left!.insert(data);
      }
    } else {
      if (right == null) {
        right = Node(value: data);
      } else {
        right!.insert(data);
      }
    }
  }

  bool search(int data) {
    if (data == value) {
      return true;
    }
    if (data < value) {
      if (left == null) {
        return false;
      }
      return left!.search(data);
    }
    if (right == null) {
      return false;
    }
    return right!.search(data);
  }
}
