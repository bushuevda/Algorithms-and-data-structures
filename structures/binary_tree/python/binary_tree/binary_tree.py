from queue import Queue
from collections import deque
from typing import Union

class BinaryTree:
    #Node search +
    #Inserting a node +
    #Deleting a node +
    #Прямой обход (КЛП) Preorder +
    #Центрированный обход (ЛКП) Inorder +
    #Обратный обход (ЛПК) Postorder +
    #Обход в ширину Level-order +
    def __init__(self, value, parent = None) -> None:
        self.left_: BinaryTree = None
        self.right_: BinaryTree = None
        self.parent_: BinaryTree = parent
        self.value_: Union[int, float] = value
        self.queue_: Queue = None
        self.stack_: deque = None

    # Test+ A function for inserting value into a tree
    def insert(self, value):
        #Inner function to find node to insert
        def find_insert_node(value: Union[int, float], parent) -> None: 
            if value < parent.value_:
                if parent.left_ is None:
                    parent.left_ = BinaryTree(value, parent)
                else:
                    find_insert_node(value, parent.left_)
            elif value > parent.value_:
                if parent.right_ is None:
                    parent.right_ = BinaryTree(value, parent)
                else:
                    find_insert_node(value, parent.right_)
        
        find_insert_node(value, self)

    # Test+ A function to search node to insert value
    def find_insert_node(self, value: Union[int, float], parent) -> None: 
        if value < parent.value_:
            if parent.left_ is None:
                parent.left_ = BinaryTree(value, parent)
            else:
                self.find_insert_node(value, parent.left_)
        elif value > parent.value_:
            if parent.right_ is None:
                parent.right_ = BinaryTree(value, parent)
            else:
                self.find_insert_node(value, parent.right_)
    
    # Test+ Traversal using stack (Direct Traversal (Preorder): main node -> left node -> right node)
    def traversal_preorder_stack(self):
        self.stack_ = deque()
        self.stack_.append(self)
        array_traversal = []

        while len(self.stack_) > 0:
            node = self.stack_.pop()
            array_traversal.append(node.value_)
            if node.right_:
                self.stack_.append(node.right_)
            if node.left_:
                self.stack_.append(node.left_)
        return array_traversal
   
    # Test+ Traversal using recursion (Direct Traversal (Preorder): main node -> left node-> right node)
    def traversal_preorder_recursive(self, node, array_traversal = None):
        if node != None:
            if array_traversal != None:
                array_traversal.append(node.value_)
            self.traversal_preorder_recursive(node.left_, array_traversal)
            self.traversal_preorder_recursive(node.right_, array_traversal)
        return array_traversal

    # Test+ Traversal using recursion (Reverse bypass (Postorder))
    def traversal_postorder_recursive(self, node, array_traversal = None):
        if node != None:
            self.traversal_postorder_recursive(node.left_, array_traversal)
            self.traversal_postorder_recursive(node.right_, array_traversal)
            if array_traversal != None:
                array_traversal.append(node.value_)
        return array_traversal

    # Test+ Traversal using recursion (Centered bypass (Inorder))
    def traversal_inorder_recursive(self, node, array_traversal = None):
        if node != None:
            self.traversal_inorder_recursive(node.left_, array_traversal)
            if array_traversal != None:
                array_traversal.append(node.value_)
            self.traversal_inorder_recursive(node.right_, array_traversal)
        
        return array_traversal
          
    # Test+ Breadth traversal using a queue (Level-order)
    def traversal_lvl_order_queue(self, node, array_traversal = None):
        self.queue_ = Queue()
        self.queue_.put(self)
        
        while not self.queue_.empty():
            node = self.queue_.get()
            if array_traversal != None:
                array_traversal.append(node.value_)
            if node.left_:
                self.queue_.put(node.left_)
            if node.right_:
                self.queue_.put(node.right_)
        return array_traversal

    # The function of searching for a node with a given value
    def findNode(self, value):
        node = self
        while node:
            if value == node.value_:
                return node
            if value < node.value_:
                node = node.left_
            if value > node.value_:
                node = node.right_
        return None
        
    # The function to remove a node with a given value  
    def deleteNode(self, value):
        #Inner function to remove a node with a given value
        def remove(node):
            if node.parent_:
                if node.parent_.left_ == node:
                    node.parent_.left_ = None
                elif node.parent_.right_ == node:
                    node.parent_.right_ = None
            else:
                node.value_ = None

        #Inner function for search right node with a max value
        def findMinRight(node):
            if node:
                if node.right_:
                    temp_node = node.right_
                    while temp_node:
                        if temp_node.left_:
                            temp_node = temp_node.left_
                        elif temp_node.right_:
                            temp_node = temp_node.right_
                        else:
                            break
                    node.value_ = temp_node.value_
                    remove(temp_node)

        #Inner function for search left node with a min value
        def findMaxLeft(node):
            if node:
                if node.left_:
                    temp_node = node.left_
                    while temp_node.right_ or temp_node.left_:
                        if temp_node.right_:
                            temp_node = temp_node.right_
                        elif temp_node.left_:
                            temp_node = temp_node.left_
                        else:
                            break
                    node.value_ = temp_node.value_
                    remove(temp_node)

        node = self.findNode(value)
        #case 1 (not childs)
        if node.right_ == None and node.left_ == None:
            print('------------------CASE1')
            remove(node)
        
        #case 2 (child left xor child right)
        elif node.right_ == None and node.left_ or node.right_ and node.left_ == None:
            print('------------------CASE2')
            if node.right_:
                node.value_ = node.right_.value_
                node.right_ = node.right_.right_
            else:
                node.value_ = node.left_.value_
                node.left_ = node.left_.left_
        
        #case 3 (child left and child right)
        else:
            # for main node (find max from left)
            if not node.parent_:
                findMinRight(node)

            # find max from left
            elif node.parent_.left_ == node:
                findMaxLeft(node)
            # find min from right
            elif node.parent_.right_ == node:
                findMinRight(node)


if __name__ == "__main__":
    b = BinaryTree(25)


    b.insert(40)
    b.insert(42)
    b.insert(45)
    b.insert(43)
    b.insert(48)
    print(b.traversal_lvl_order_queue(b, []))
    b.deleteNode(45)
    print(b.traversal_lvl_order_queue(b, []))
    b.deleteNode(40)
    print(b.traversal_lvl_order_queue(b, []))

