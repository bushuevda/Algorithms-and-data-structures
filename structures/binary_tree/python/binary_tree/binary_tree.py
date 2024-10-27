from queue import Queue
from collections import deque
from typing import Union

class BinaryTree:
    #Поиск узла
    #Вставка узла
    #Удаление узла
    #Прямой обход (КЛП) Preorder +
    #Центрированный обход (ЛКП) Inorder +
    #Обратный обход (ЛПК) Postorder +
    #Обход в ширину Level-order +
    def __init__(self, value, parent = None) -> None:
        self.leftTree_: BinaryTree = None
        self.rightTree_: BinaryTree = None
        self.parentTree_: BinaryTree = parent
        self.value_: Union[int, float] = value
        self.queue_: Queue = None
        self.stack_: deque = None

    def insert(self, value):
        self.find_insert_node(value, self)
        
    def find_insert_node(self, value: Union[int, float], parent) -> None: 
        if value < parent.value_:
            if parent.leftTree_ is None:
                parent.leftTree_ = BinaryTree(value, parent)
            else:
                self.find_insert_node(value, parent.leftTree_)
        elif value > parent.value_:
            if parent.rightTree_ is None:
                parent.rightTree_ = BinaryTree(value, parent)
            else:
                self.find_insert_node(value, parent.rightTree_)
    
    # Test+ Обход со стэком (Прямой обход (КЛП или Preorder): корень->левое поддерево->правое поддерево->)
    def traversal_preorder_stack(self):
        self.stack_ = deque()
        self.stack_.append(self)
        array_traversal = []

        while len(self.stack_) > 0:
            current_node = self.stack_.pop()
            array_traversal.append(current_node.value_)
            if current_node.rightTree_:
                self.stack_.append(current_node.rightTree_)
            if current_node.leftTree_:
                self.stack_.append(current_node.leftTree_)
        return array_traversal
   
    # Test+ Обход рекурсией (Прямой обход (КЛП или Preorder): корень->левое поддерево->правое поддерево->)
    def traversal_preorder_recursive(self, node, array_traversal = None):
        
        if node != None:
            if array_traversal != None:
                array_traversal.append(node.value_)
            #print(f'node = {node.value_}')
            self.traversal_preorder_recursive(node.leftTree_, array_traversal)
            self.traversal_preorder_recursive(node.rightTree_, array_traversal)
        return array_traversal

    #Обход рекурсией (Обратный обход)
    def traversal_postorder_recursive(self, node, ):
        if node != None:
            self.traversal_postorder_recursive(node.leftTree_)
            self.traversal_postorder_recursive(node.rightTree_)
            print(f'node = {node.value_}')

    #Обход рекурсией (Центрированный обход)
    def traversal_inorder_recursive(self, node):
        if node != None:
            self.traversal_inorder_recursive(node.leftTree_)
            print(f'node = {node.value_}')
            self.traversal_inorder_recursive(node.rightTree_)
          

    #Обход в ширину с помощью очереди Level-order
    def traversal_lvl_order_queue(self, node):
        self.queue_.put(self)
        
        while not self.queue_.empty():
            current_node = self.queue_.get()
            print(current_node.value_)
            if current_node.leftTree_:
                self.queue_.put(current_node.leftTree_)
            if current_node.rightTree_:
                self.queue_.put(current_node.rightTree_)

    
    def findValue(self, value):
        if self.value_ == value:
            return value
        if self.rightTree_:
            if self.rightTree_.value_ < value:
                return self.rightTree_.findValue(value)
            elif self.rightTree_.value_ == value:
                return self.rightTree_.value_
        if self.leftTree_:
            if self.leftTree_.value_ > value:
                return self.leftTree_.findValue(value)
            elif self.leftTree_.value_ == value:
                return self.leftTree_.value_
        ...

        
    
if __name__ == "__main__":
    b = BinaryTree(10)
    b.insert(7)
    b.insert(12)
    b.insert(6)
    b.insert(8)
    b.insert(11)
    #b.insert(13)
    b.traversal_preorder_stack()
    print('f', b.findValue(11))