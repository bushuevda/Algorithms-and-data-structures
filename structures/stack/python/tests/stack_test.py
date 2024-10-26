import importlib.util
import unittest

file = importlib.util.spec_from_file_location("stack","./stack/stack.py")
module = importlib.util.module_from_spec(file)
file.loader.exec_module(module)
Stack = module.Stack

# Test data for testing
TEST_DATA1 = [1, 2, 3, 4, 5, 6, 7]
TEST_DATA2 = ["a", "e", "b", "d", "f"]
TEST_DATA3 = [7.5, 5.3, 1.2, 4.7, 15.2, 6.3, 7.4]

# Test data for testing function pop
TEST_DATA_POP1 = [1, 2, 3, 4] # -3
TEST_DATA_POP2 = ["a", "e", "b"]; # -2
TEST_DATA_POP3 = [7.5, 5.3, 1.2, 4.7, 15.2, 6.3]; # -1


#Helper util for testing
class UtilTestStack:
    #The helped function for testing the stack function push 
    def testPush(self, array) -> list[any]:
        s = Stack()
        for e in array:
            s.push(e)
        return s.getArray()
    
    #The helped function for testing the stack function pop
    def testPop(self, array, size_pop) -> list[any]:
        s = Stack()
        for e in array:
            s.push(e)
        for i in range(size_pop):
            s.pop()
        return s.getArray()

    #The helped function for testing the stack function top
    def testTop(self, array) -> any:
        s = Stack()
        for e in array:
            s.push(e)
        return s.top()

    #The helped function for testing the stack function size 
    def testSize(self, array) -> int:
        s = Stack()
        for e in array:
            s.push(e)
        return s.size()
    
    #The helped function for testing the stack function empty 
    def testEmpty(self, array) -> bool:
        s = Stack()
        for e in array:
            s.push(e)
        return s.empty()
    
    #The helped function for testing the stack function swap 
    def testSwap(self, s1, s2) -> None:
        temp = s1
        s1 = s2
        s2 = temp
        return s1, s2


class TestStack(unittest.TestCase):
    def __init__(self, methodName: str = "runTest") -> None:
        super().__init__(methodName)
        self.utilTest_ = UtilTestStack()
    #Testing the stack function push  
    def testPush(self) -> None:
        a = self.utilTest_.testPush(TEST_DATA1)
        b = TEST_DATA1
        self.assertEqual(a, b, f"list expect {a}, list recived {b} (TEST_DATA1)")

        a = self.utilTest_.testPush(TEST_DATA2)
        b = TEST_DATA2
        self.assertEqual(a, b, f"list expect {a}, list recived {b} (TEST_DATA2)")

        a = self.utilTest_.testPush(TEST_DATA3)
        b = TEST_DATA3
        self.assertEqual(a, b, f"list expect {a}, list recived {b} (TEST_DATA3)")

    #Testing the stack function pop
    def testPop(self) -> None:
        a = self.utilTest_.testPop(TEST_DATA1, 3)
        b = TEST_DATA_POP1
        self.assertEqual(a, b, f"list expect {a}, list recived {b} (TEST_DATA1 and TEST_DATA_POP1)")

        a = self.utilTest_.testPop(TEST_DATA2, 2)
        b = TEST_DATA_POP2
        self.assertEqual(a, b, f"list expect {a}, list recived {b} (TEST_DATA2 and TEST_DATA_POP2)")

        a = self.utilTest_.testPop(TEST_DATA3, 1)
        b = TEST_DATA_POP3
        self.assertEqual(a, b, f"list expect {a}, list recived {b} (TEST_DATA3 and TEST_DATA_POP3)")

        s = Stack()
        a = s.pop()
        b = None
        self.assertEqual(a, b, f"list expect {a}, list recived {b} ([] and None)")

    #Testing the stack function top
    def testTop(self) -> None:
        a = self.utilTest_.testTop(TEST_DATA1)
        b = TEST_DATA1[len(TEST_DATA1) - 1]
        self.assertEqual(a, b, f"top expect {a}, top recived {b} (TEST_DATA1)")

        a = self.utilTest_.testTop(TEST_DATA2)
        b = TEST_DATA2[len(TEST_DATA2) - 1]
        self.assertEqual(a, b, f"top expect {a}, top recived {b} (TEST_DATA2)")

        a = self.utilTest_.testTop(TEST_DATA3)
        b = TEST_DATA3[len(TEST_DATA3) - 1]
        self.assertEqual(a, b, f"list expect {a}, top recived {b} (TEST_DATA3)")
        
    #Testing the stack function size 
    def testSize(self) -> None:
        a = self.utilTest_.testSize(TEST_DATA1)
        b = len(TEST_DATA1)
        self.assertEqual(a, b, f"size expect {a}, size recived {b} (TEST_DATA1)")

        a = self.utilTest_.testSize(TEST_DATA2)
        b = len(TEST_DATA2)
        self.assertEqual(a, b, f"size expect {a}, size recived {b} (TEST_DATA2)")

        a = self.utilTest_.testSize(TEST_DATA3)
        b = len(TEST_DATA3)
        self.assertEqual(a, b, f"size expect {a}, size recived {b} (TEST_DATA3)")
    
    #Testing the stack function empty 
    def testEmpty(self) -> None:
        a = self.utilTest_.testEmpty(TEST_DATA1)
        b = False
        self.assertEqual(a, b, f"bool expect {a}, bool recived {b} (TEST_DATA1)")

        a = self.utilTest_.testEmpty(TEST_DATA2)
        b = False
        self.assertEqual(a, b, f"bool expect {a}, bool recived {b} (TEST_DATA2)")

        a = self.utilTest_.testEmpty(TEST_DATA3)
        b = False
        self.assertEqual(a, b, f"bool expect {a}, bool recived {b} (TEST_DATA3)")

        a = self.utilTest_.testEmpty([])
        b = True
        self.assertEqual(a, b, f"bool expect {a}, bool recived {b} ([])")

    #Testing the stack function swap 
    def testSwap(self) -> None:
        a1 = TEST_DATA1
        b1, b2 = TEST_DATA3, TEST_DATA3
        a1, b1 = self.utilTest_.testSwap(a1, b1)
        self.assertEqual(a1, b2, f"list expect {a1}, list recived {b2} ([])")

        a1 = TEST_DATA3
        b1, b2 = TEST_DATA1, TEST_DATA1
        a1, b1 = self.utilTest_.testSwap(a1, b1)
        self.assertEqual(a1, b2, f"list expect {a1}, list recived {b2} ([])")

if __name__ =="__main__":
    unittest.main()