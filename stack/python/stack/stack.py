
class Stack:
    def __init__(self) -> None:
        self.array_ = list()
    
    #A function that returns stack size
    def size(self) -> int:
        return len(self.array_)
    
    #A function that checks the stack for emptiness
    def empty(self) -> bool:
        return True if self.size() <= 0 else False
    
    #A function that adds an element to the stack
    def push(self, element: any) -> None:
        self.array_.append(element)
    
    #A function that removes a top element from the stack and returns it
    def pop(self) -> any:
        if not self.empty():
            temp = self.array_[self.size() - 1]
            self.array_ = self.array_[:self.size() - 1]
            return temp
        return None
    
    #A function that swaps the stack arrays
    def swap(self, stack) -> None:
        if type(stack) == type(Stack()):
            temp = stack.array_
            stack.array_ = self.array_
            self.array_ = temp
    
    #A function that returns the top element of the stack 
    def top(self) -> any:
        if not self.empty():
            return self.array_[self.size() - 1]
        return None
    
    #A function that shows the stack arrays
    def show(self) -> None:
        print(self.array_)
    
    #A function that returns the stack array
    def getArray(self) -> list[any]:
        return self.array_