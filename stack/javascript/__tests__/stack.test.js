import { Stack } from "../stack/stack.js";

import assert from "assert";

/*
Data for testing all stack functions
 */
const TEST_DATA1 = [1, 2, 3, 4, 5, 6, 7];
const TEST_DATA2 = ["a", "e", "b", "d", "f"];
const TEST_DATA3 = [7.5, 5.3, 1.2, 4.7, 15.2, 6.3, 7.4];

/*
Data for testing stack function pop
 */
const TEST_DATA_POP1 = [1, 2, 3, 4]; // -3
const TEST_DATA_POP2 = ["a", "e", "b"]; // -2
const TEST_DATA_POP3 = [7.5, 5.3, 1.2, 4.7, 15.2, 6.3]; //- 1


/*
Helped function for testing the push stack function 
 */
function testPush(array){
    let stack = new Stack();

    array.forEach(element => {
        stack.push(element);
    });
    return stack.array_
}


/*
Helped function for testing the pop stack function 
 */
function testPop(array, size_pop){
    let stack = new Stack();

    array.forEach(element => {
        stack.push(element);
    });

    for (let i = 0; i < size_pop; i++)
        stack.pop();
    return stack.array_;
}

/*
Helped function for testing the top stack function 
 */
function testTop(array){
    let stack = new Stack();
    array.forEach(element => {
        stack.push(element);
    });
    let temp = stack.top();
    return temp;
}

/*
Helped function for testing the size stack function 
 */
function testSize(array){
    let stack = new Stack();

    array.forEach(element => {
        stack.push(element);
    });
    return stack.size();
}

/*
Helped function for testing the empty stack function 
 */
function testEmpty(array){
    let stack = new Stack();

    array.forEach(element => {
        stack.push(element);
    });
    return stack.empty();
    
}

/*
Helped function for testing the swap stack function 
 */
function testSwap(array1, array2) {
    var stack1 = new Stack();
    array1.forEach(element => {
        stack1.push(element);
    });

    var stack2 = new Stack();
    array2.forEach(element => {
        stack2.push(element);
    });
    
    stack1.swap(stack2);
    
    return stack1.array_;
}


/*
Testing the swap stack function
*/
it('Test swap with TEST_DATA1 and TEST_DATA3', () => {
    var a = TEST_DATA1;
    var b = TEST_DATA3;
    let result = testSwap(a, b);
    assert.deepEqual(result, TEST_DATA3);
})

it('Test swap with TEST_DATA3 and TEST_DATA1', () => {
    var a = TEST_DATA3;
    var b = TEST_DATA1;
    let result = testSwap(a, b);
    assert.deepEqual(result, TEST_DATA1);
})


/*
Testing the empty stack function
*/
it('Test empty function with TEST_DATA1', () => {
    assert.equal(testEmpty(TEST_DATA1), false);
})

it('Test empty function with TEST_DATA2', () => {
    assert.equal(testEmpty(TEST_DATA2), false);
})

it('Test empty function with TEST_DATA3', () => {
    assert.equal(testEmpty(TEST_DATA3), false);
})

it('Test empty function with []', () => {
    assert.equal(testEmpty([]), true);
})


/*
Testing the size stack function
*/
it('Test size function with TEST_DATA1', () => {
    assert.equal(testSize(TEST_DATA1), TEST_DATA1.length);
})
it('Test size function with TEST_DATA2', () => {
    assert.equal(testSize(TEST_DATA2), TEST_DATA2.length);
})
it('Test size function with TEST_DATA3', () => {
    assert.equal(testSize(TEST_DATA3), TEST_DATA3.length);
})


/*
Testing the push stack function
*/
it('Test push function with TEST_DATA1', () => {
    assert.deepEqual(testPush(TEST_DATA1), TEST_DATA1);
})
it('Test push function with TEST_DATA2', () => {
    assert.deepEqual(testPush(TEST_DATA2), TEST_DATA2);
})
it('Test push function with TEST_DATA3', () => {
    assert.deepEqual(testPush(TEST_DATA3), TEST_DATA3);
})


/*
Testing the pop stack function
*/
it('Test pop function with TEST_DATA1 and with TEST_DATA_POP1', () => {
    assert.deepEqual(testPop(TEST_DATA1, 3), TEST_DATA_POP1);
})

it('Test pop function with TEST_DATA2 and with TEST_DATA_POP2', () => {
    assert.deepEqual(testPop(TEST_DATA2, 2), TEST_DATA_POP2);
})

it('Test pop function with TEST_DATA3 and with TEST_DATA_POP3', () => {
    assert.deepEqual(testPop(TEST_DATA3, 1), TEST_DATA_POP3);
})


/*
Testing the top stack function
*/
it('Test top function with TEST_DATA1', () => {
    assert.equal(testTop(TEST_DATA1), TEST_DATA1[TEST_DATA1.length - 1]);
})

it('Test top function with TEST_DATA2', () => {
    assert.equal(testTop(TEST_DATA2), TEST_DATA2[TEST_DATA2.length - 1]);
})

it('Test top function with TEST_DATA3', () => {
    assert.equal(testTop(TEST_DATA3), TEST_DATA3[TEST_DATA3.length - 1]);
})

it('Test top function with undefided', () => {
    assert.equal(testTop([]), undefined);
})

