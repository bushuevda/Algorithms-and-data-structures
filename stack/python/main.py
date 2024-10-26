from stack.stack import Stack

a = Stack()
a.push(1)
a.push(2)
a.push(3)
print(a.top())

a.pop()

print(a.array_)
print(a.top())