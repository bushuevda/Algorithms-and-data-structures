import { Stack } from "../stack/stack";

/*
Data for testing all stack functions
*/
const TEST_DATA1: Array<number> = [1, 2, 3, 4, 5, 6, 7];
const TEST_DATA2: Array<string> = ["a", "e", "b", "d", "f"];
const TEST_DATA3: Array<number> = [7.5, 5.3, 1.2, 4.7, 15.2, 6.3, 7.4];

/*
Data for testing stack function pop
*/
const TEST_DATA_POP1: Array<number> = [1, 2, 3, 4]; // -3
const TEST_DATA_POP2: Array<string> = ["a", "e", "b"]; // -2
const TEST_DATA_POP3: Array<number> = [7.5, 5.3, 1.2, 4.7, 15.2, 6.3]; //- 1

/*
Helped function for testing the push stack function 
*/
function testPush<T>(array: Array<T>): T[]{
    let stack: Stack<T> = new Stack<T>();

    array.forEach(element => {
        stack.push(element);
    });
    return stack.getArray()
}

/*
Helped function for testing the pop stack function 
*/
function testPop<T>(array: Array<T>, size_pop: number): T[]{
    let stack: Stack<T> = new Stack<T>();

    array.forEach(element => {
        stack.push(element);
    });

    for (let i = 0; i < size_pop; i++)
        stack.pop();
    return stack.getArray();
}

/*
Helped function for testing the top stack function 
*/
function testTop<T>(array: Array<T>): T | undefined{
    let stack: Stack<T> = new Stack<T>();
    array.forEach(element => {
        stack.push(element);
    });
    let temp: T | undefined = stack.top();
    return temp;
}

/*
Helped function for testing the size stack function 
*/
function testSize<T>(array: Array<T>): number{
    let stack: Stack<T> = new Stack<T>();

    array.forEach(element => {
        stack.push(element);
    });
    return stack.size();
}

/*
Helped function for testing the empty stack function 
*/
function testEmpty<T>(array: Array<T>): boolean{
    let stack: Stack<T> = new Stack<T>();

    array.forEach(element => {
        stack.push(element);
    });
    return stack.empty();
    
}

/*
Helped function for testing the swap stack function 
*/
function testSwap<T>(array1: Array<T>, array2: Array<T>): T[]{
    let stack1: Stack<T> = new Stack<T>();
    array1.forEach(element => {
        stack1.push(element);
    });

    let stack2: Stack<T> = new Stack<T>();
    array2.forEach(element => {
        stack2.push(element);
    });

    stack1.swap(stack2);
    return stack1.getArray();
}

/*
Testing the swap stack function
*/
test("Test swap with TEST_DATA1 and TEST_DATA3", () => {
    expect(testSwap(TEST_DATA1, TEST_DATA3)).toStrictEqual(TEST_DATA3);
})
test("Test swap with TEST_DATA3 and TEST_DATA1", () => {
    expect(testSwap(TEST_DATA3, TEST_DATA1)).toStrictEqual(TEST_DATA1);
})

/*
Testing the empty stack function
*/
test("Test empty function with TEST_DATA1", () => {
    expect(testEmpty(TEST_DATA2)).toBe(false);
})
test("Test empty function with TEST_DATA2", () => {
    expect(testEmpty(TEST_DATA2)).toBe(false);
})
test("Test empty function with TEST_DATA3", () => {
    expect(testEmpty(TEST_DATA3)).toBe(false);
})
test("Test empty function with []", () => {
    expect(testEmpty([])).toBe(true);
})

/*
Testing the size stack function
*/
test("Test size function with TEST_DATA1", () => {
    expect(testSize(TEST_DATA1)).toBe(TEST_DATA1.length);
})
test("Test size function with TEST_DATA2", () => {
    expect(testSize(TEST_DATA2)).toBe(TEST_DATA2.length);
})
test("Test size function with TEST_DATA3", () => {
    expect(testSize(TEST_DATA3)).toBe(TEST_DATA3.length);
})


/*
Testing the push stack function
*/
test("Test push function with TEST_DATA1", () => {
    expect(testPush(TEST_DATA1)).toStrictEqual(TEST_DATA1);
})
test("Test push function with TEST_DATA2", () => {
    expect(testPush(TEST_DATA2)).toStrictEqual(TEST_DATA2);
})
test("Test push function with TEST_DATA3", () => {
    expect(testPush(TEST_DATA3)).toStrictEqual(TEST_DATA3);    
})


/*
Testing the pop stack function
*/
test("Test pop function with TEST_DATA1 and with TEST_DATA_POP1", () => {
    expect(testPop(TEST_DATA1, 3)).toStrictEqual(TEST_DATA_POP1);
})
test("Test pop function with TEST_DATA2 and with TEST_DATA_POP2", () => {
    expect(testPop(TEST_DATA2, 2)).toStrictEqual(TEST_DATA_POP2);
})
test("Test pop function with TEST_DATA3 and with TEST_DATA_POP3", () => {
    expect(testPop(TEST_DATA3, 1)).toStrictEqual(TEST_DATA_POP3);
})
test("Test pop function with undefided", () => {
    let stack: Stack<number> = new Stack<number>();
    expect(stack.pop()).toBe(undefined);
})


/*
Testing the top stack function
*/
test("Test top function with TEST_DATA1", () => {
    expect(testTop(TEST_DATA1)).toBe(TEST_DATA1[TEST_DATA1.length - 1]);
})
test("Test top function with TEST_DATA2", () => {
    expect(testTop(TEST_DATA2)).toBe(TEST_DATA2[TEST_DATA2.length - 1]);
})
test("Test top function with TEST_DATA3", () => {
    expect(testTop(TEST_DATA3)).toBe(TEST_DATA3[TEST_DATA3.length - 1]);
})
test("Test top function with undefided", () => {
    expect(testTop([])).toBe(undefined);
})