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

# An array of objects of the BinaryTree class.
# The dimension must be equal to ARRAY_VALUES.
ARRAY_TREES = [
    BinaryTree(array[0]) for array in ARRAY_VALUES
]

# Required function for filling ARRAY_TREES objects with data.
def initBinaryTrees():
    if len(ARRAY_TREES) == len(ARRAY_VALUES):
        for i in range(len(ARRAY_TREES)):
            for value in ARRAY_VALUES[i]:
                ARRAY_TREES[i].insert(value)

initBinaryTrees()


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
    [], #1
    [], #2
    [], #3
    [], #4
    [], #5
    [], #6
    [], #7
    [], #8
    [], #9
    [], #10
]

#CASE 4
OUTPUT_DATA_CASE4 = [
    [], #1
    [], #2
    [], #3
    [], #4
    [], #5
    [], #6
    [], #7
    [], #8
    [], #9
    [], #10
]

#CASE 5
OUTPUT_DATA_CASE5 = [
    [], #1
    [], #2
    [], #3
    [], #4
    [], #5
    [], #6
    [], #7
    [], #8
    [], #9
    [], #10
]




class TestBinaryTree(unittest.TestCase):
    #CASE 1
    #Testing the direct bypass function using stack.
    def test_traversal_preorder_stack(self):
        for i, tree in enumerate(ARRAY_TREES):
            if len(OUTPUT_DATA_CASE1) != len(ARRAY_TREES):
                break
            real_data = tree.traversal_preorder_stack()
            test_date = OUTPUT_DATA_CASE1[i]
            self.assertListEqual(real_data, test_date)

    #CASE 2
    #Testing the direct traversal function using recursion.
    def test_traversal_preorder_recursive(self):
        for i, tree in enumerate(ARRAY_TREES):
            if len(OUTPUT_DATA_CASE2) != len(ARRAY_TREES):
                break
            real_data = [[] for k in range(len(OUTPUT_DATA_CASE2))]
            real_data[i] = tree.traversal_preorder_recursive(tree, real_data[i])
            test_date = OUTPUT_DATA_CASE2[i]
            print(real_data[i], test_date)
            self.assertListEqual(real_data[i], test_date)


if __name__ == "__main__":
    unittest.main()