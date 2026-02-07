using tree;




tree.BST<int, int> bst = new BST<int, int>(6, 111);
bst.Insert(4, 33);
bst.Insert(2, 33);
bst.Insert(1, 33);
bst.Insert(3, 33);

var visit = (int key, int data) => Console.WriteLine(key.ToString(), data.ToString());

bst.PreOrderRecursive(bst, visit);

Console.Write("\n");
bst = bst.RemoveNode(bst, 1)!;
bst.PreOrderRecursive(bst, visit);

Console.Write("\n");
bst = bst.RemoveNode(bst, 2)!;
bst.PreOrderRecursive(bst, visit);

Console.Write("\n");
bst = bst.RemoveNode(bst, 6)!;
bst.PreOrderRecursive(bst, visit);

Console.Write("\n");
bst = bst.RemoveNode(bst, 3)!;
bst.PreOrderRecursive(bst, visit);
