import importlib.util
import unittest

file = importlib.util.spec_from_file_location("binary_tree","./binary_tree/binary_tree.py")
module = importlib.util.module_from_spec(file)
file.loader.exec_module(module)
BinaryTree = module.BinaryTree


                
# An array of values for filling objects of the BinaryTree class. 
# The dimension must be equal to ARRAY_TREES.


ARRAY_VALUES = [
    [10, 7, 12, 6, 8, 11, 13], #1
    [10, 7, 12, 6, 8, 11], #2
    [10, 7, 12, 6, 11], #3
    [10, 7, 12, 8, 13], #4
    [10, 7, 12, 6, 11, 13], #5
    [10, 7, 12, 8, 11, 13], #6
    [10, 7, 12, 6, 8, 13], #7
    [10, 7, 6, 8], #8
    [10, 12, 11, 13], #9
    [10, 7, 12], #10
]




# Required function for filling ARRAY_TREES objects with data.
def initBinaryTrees() -> list[BinaryTree]:
    # An array of objects of the BinaryTree class.
    # The dimension must be equal to ARRAY_VALUES.
    ARRAY_TREES = [BinaryTree(array[0]) for array in ARRAY_VALUES]
    if len(ARRAY_TREES) == len(ARRAY_VALUES):
        for i in range(len(ARRAY_TREES)):
            for value in ARRAY_VALUES[i]:
                ARRAY_TREES[i].insert(value)
    return ARRAY_TREES




#CASE 1
OUTPUT_DATA_CASE1 = [
    [10, 7, 6, 8, 12, 11, 13], #1
    [10, 7, 6, 8, 12, 11], #2
    [10, 7, 6, 12, 11], #3
    [10, 7, 8, 12, 13], #4
    [10, 7, 6, 12, 11, 13], #5
    [10, 7, 8, 12, 11, 13], #6
    [10, 7, 6, 8, 12, 13], #7
    [10, 7, 6, 8], #8
    [10, 12, 11, 13], #9
    [10, 7, 12], #10
]

#CASE 2
OUTPUT_DATA_CASE2 = [
    [10, 7, 6, 8, 12, 11, 13], #1
    [10, 7, 6, 8, 12, 11], #2
    [10, 7, 6, 12, 11], #3
    [10, 7, 8, 12, 13], #4
    [10, 7, 6, 12, 11, 13], #5
    [10, 7, 8, 12, 11, 13], #6
    [10, 7, 6, 8, 12, 13], #7
    [10, 7, 6, 8], #8
    [10, 12, 11, 13], #9
    [10, 7, 12], #10
]

#CASE 3
OUTPUT_DATA_CASE3 = [
    [6, 8, 7, 11, 13, 12, 10], #1
    [6, 8, 7, 11, 12, 10], #2
    [6, 7, 11, 12, 10], #3
    [8, 7, 13, 12, 10], #4
    [6, 7, 11, 13, 12, 10], #5
    [8, 7, 11, 13, 12, 10], #6
    [6, 8, 7, 13, 12, 10], #7
    [6, 8, 7, 10], #8
    [11, 13, 12, 10], #9
    [7, 12, 10], #10
]

#CASE 4
OUTPUT_DATA_CASE4 = [
    [6, 7, 8, 10, 11, 12, 13], #1
    [6, 7, 8, 10, 11, 12], #2
    [6, 7, 10, 11, 12], #3
    [7, 8, 10, 12, 13], #4
    [6, 7, 10, 11, 12, 13], #5
    [7, 8, 10, 11, 12, 13], #6
    [6, 7, 8, 10, 12, 13], #7
    [6, 7, 8, 10], #8
    [10, 11, 12, 13], #9
    [7, 10, 12], #10
]

#CASE 5
OUTPUT_DATA_CASE5 = [
    [10, 7, 12, 6, 8, 11, 13], #1
    [10, 7, 12, 6, 8, 11], #2
    [10, 7, 12, 6, 11], #3
    [10, 7, 12, 8, 13], #4
    [10, 7, 12, 6, 11, 13], #5
    [10, 7, 12, 8, 11, 13], #6
    [10, 7, 12, 6, 8, 13], #7
    [10, 7, 6, 8], #8
    [10, 12, 11, 13], #9
    [10, 7, 12], #10
]


#CASE 6
INPUT_DATA_CASE6 = [
    [10, 7, 12, 6, 8, 11, 13], #1
    [10, 7, 12, 6, 8, 11], #2
    [10, 7, 12, 6, 11], #3
    [10, 7, 12, 8, 13], #4
    [10, 7, 12, 6, 11, 13], #5
    [10, 7, 12, 8, 11, 13], #6
    [10, 7, 12, 6, 8, 13], #7
    [10, 7, 6, 8], #8
    [10, 12, 11, 13], #9
    [10, 7, 12], #10
]

OUTPUT_DATA_CASE6 = [
    [10, 7, 12, 6, 8, 11, 13], #1
    [10, 7, 12, 6, 8, 11], #2
    [10, 7, 12, 6, 11], #3
    [10, 7, 12, 8, 13], #4
    [10, 7, 12, 6, 11, 13], #5
    [10, 7, 12, 8, 11, 13], #6
    [10, 7, 12, 6, 8, 13], #7
    [10, 7, 6, 8], #8
    [10, 12, 11, 13], #9
    [10, 7, 12], #10
]


#CASE 7
INPUT_DATA_CASE7 = [
    [10], #1
    [8], #2
    [12, 11], #3
    [13], #4
    [7], #5
    [12], #6
    [7], #7
    [8], #8
    [11], #9
    [10, 7, 12], #10
]

OUTPUT_DATA_CASE7 = [
    [11, 7, 12, 6, 8, 13], #1
    [10, 7, 12, 6, 11], #2
    [10, 7, 6], #3
    [10, 7, 12, 8], #4
    [10, 6, 12, 11, 13], #5
    [10, 7, 13, 8, 11], #6
    [10, 6, 12, 8, 13], #7
    [10, 7, 6], #8
    [10, 12, 13], #9
    [None], #10
]

class TestBinaryTree(unittest.TestCase):
    #CASE 1
    #Testing the direct bypass function using stack.
    def test_traversal_preorder_stack(self):
        ARRAY_TREES = initBinaryTrees()
        for i, tree in enumerate(ARRAY_TREES):
            if len(OUTPUT_DATA_CASE1) != len(ARRAY_TREES):
                break
            real_data = tree.traversal_preorder_stack()
            test_date = OUTPUT_DATA_CASE1[i]
            self.assertListEqual(real_data, test_date)

    #CASE 2
    #Testing the direct traversal function using recursion.
    def test_traversal_preorder_recursive(self):
        ARRAY_TREES = initBinaryTrees()
        for i, tree in enumerate(ARRAY_TREES):
            if len(OUTPUT_DATA_CASE2) != len(ARRAY_TREES):
                break
            real_data = [[] for k in range(len(OUTPUT_DATA_CASE2))]
            real_data[i] = tree.traversal_preorder_recursive(tree, real_data[i])
            test_date = OUTPUT_DATA_CASE2[i]
            self.assertListEqual(real_data[i], test_date)

    #CASE 3
    # Testing the postorder traversal function using recursion
    def test_traversal_postorder_recursive(self):
        ARRAY_TREES = initBinaryTrees()
        for i, tree in enumerate(ARRAY_TREES):
            if len(OUTPUT_DATA_CASE3) != len(ARRAY_TREES):
                break
            real_data = [[] for k in range(len(OUTPUT_DATA_CASE3))]
            real_data[i] = tree.traversal_postorder_recursive(tree, real_data[i])
            test_date = OUTPUT_DATA_CASE3[i]
            self.assertListEqual(real_data[i], test_date)

    #CASE 4
    # Testing the inorder traversal function using recursive
    def test_traversal_inorder_recursive(self):
        ARRAY_TREES = initBinaryTrees()
        for i, tree in enumerate(ARRAY_TREES):
            if len(OUTPUT_DATA_CASE4) != len(ARRAY_TREES):
                break
            real_data = [[] for k in range(len(OUTPUT_DATA_CASE4))]
            real_data[i] = tree.traversal_inorder_recursive(tree, real_data[i])
            test_date = OUTPUT_DATA_CASE4[i]
            self.assertListEqual(real_data[i], test_date)

    #CASE 5
    # Testing the beadth traversal function using a queue (Level-order)
    def test_traversal_lvl_order_queue(self):
        ARRAY_TREES = initBinaryTrees()
        for i, tree in enumerate(ARRAY_TREES):
            if len(OUTPUT_DATA_CASE5) != len(ARRAY_TREES):
                break
            real_data = [[] for k in range(len(OUTPUT_DATA_CASE5))]
            real_data[i] = tree.traversal_lvl_order_queue(tree, real_data[i])
            test_date = OUTPUT_DATA_CASE5[i]
            self.assertListEqual(real_data[i], test_date)
    
    #CASE 6
    # Testing the node find function
    def test_findNode(self):
        ARRAY_TREES = initBinaryTrees()
        for i, tree in enumerate(ARRAY_TREES):
            for j, val_find in enumerate(INPUT_DATA_CASE6[i]):
                self.assertEqual(tree.findNode(val_find).value_, OUTPUT_DATA_CASE6[i][j])

    #CASE 7
    # Testing the deletion node function with given value
    def test_deleteNode(self):
        ARRAY_TREES = initBinaryTrees()
        for i, tree in enumerate(ARRAY_TREES):
            for  val_del in INPUT_DATA_CASE7[i]:
                tree.deleteNode(val_del)
            real_data = [[] for k in range(len(OUTPUT_DATA_CASE7))]
            real_data[i] = tree.traversal_lvl_order_queue(tree, real_data[i])
            print(real_data[i])
            self.assertListEqual(real_data[i], OUTPUT_DATA_CASE7[i])
              

if __name__ == "__main__":
    unittest.main()