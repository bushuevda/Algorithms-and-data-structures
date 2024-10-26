#ifndef TEST_STACK_HPP 
#define TEST_STACK_HPP  1
#include <vector>
#include "../libs/cppunit/cppunit.h"
#include "../stack/stack.hpp"
#include <iostream>

/*
Data for testing all stack functions
 */
const std::vector<int> TEST_DATA1 {1, 2, 3, 4, 5, 6, 7};
const std::vector<std::string> TEST_DATA2{"a", "e", "b", "d", "f"};
const std::vector<int> TEST_DATA3 {7, 5, 1, 4, 15, 6, 7};

/*
Data for testing stack function pop
*/
const std::vector<int> TEST_DATA_POP1 {1, 2, 3, 4};              // -3
const std::vector<std::string> TEST_DATA_POP2 {"a", "e", "b"};   // -2
const std::vector<int> TEST_DATA_POP3 {7, 5, 1, 4, 15, 6};       //- 1


class UtilTestStack{

public:
    /*
    Helped function for testing the push stack function 
    */
   template <typename T> 
   std::vector<T> testPush(std::vector<T> array){
        Stack<T> stack= Stack<T>();
        for(auto element : array){
            stack.push(element);
        }
    return stack.getArray();
   }

    /*
    Helped function for testing the pop stack function 
    */
   template <typename T>
   std::vector<T> testPop(std::vector<T> array, int size_pop){
        Stack<T> stack= Stack<T>();
        for(auto element : array){
            stack.push(element);
        }
        for(int i = 0; i < size_pop; i++)
            stack.pop();
        return stack.getArray();
   }

    /*
    Helped function for testing the top stack function 
    */
    template <typename T>
    T testTop(std::vector<T> array){
        Stack<T> stack= Stack<T>();
        for(auto element : array){
            stack.push(element);
        }
        return stack.top();
    }


    /*
    Helped function for testing the size stack function 
    */
    template <typename T>
    int testSize(std::vector<T> array){
        Stack<T> stack= Stack<T>();
        for(auto element : array){
            stack.push(element);
        }
        return stack.size();
    }


    /*
    Helped function for testing the empty stack function 
    */
    template <typename T>
    bool testEmpty(std::vector<T> array){
        Stack<T> stack = Stack<T>();
        for(auto element : array){
            stack.push(element);
        }
        return stack.empty();
    }

    /*
    Helped function for testing the swap stack function 
    */
    template <typename T>
    std::vector<T> testSwap(std::vector<T> array1, std::vector<T> array2){
        Stack<T> stack1 = Stack<T>();
        for(auto element : array1){
            stack1.push(element);
        }

        Stack<T> stack2 = Stack<T>();
        for(auto element : array2){
            stack2.push(element);
        }

        stack1.swap(&stack2);
        return stack1.getArray();
    }


    template <typename T>
    bool assertEqualArrays(std::vector<T> array1, std::vector<T> array2){
        int size = array1.size();
        size =  array2.size() < size ? array2.size() : size;

        if(array1.size() != array2.size())
            return false;
        for(int i = 0; i < size; i++){
            if (array1.at(i) != array2.at(i))
                 return false;
        }
        return true;
    }

    template <typename T>
    bool assertEqual(T value1, T value2){
        if(value1 != value2)
            return false;
        else
            return true;
    }
};

class TestStack: public Cppunit{

private:
    UtilTestStack util;

public:

    /*
    Testing the swap stack function
    */
    void testSwap(){
        bool result = util.assertEqualArrays(util.testSwap(TEST_DATA1, TEST_DATA3), TEST_DATA3);
        CHECK(result, true);

        result = util.assertEqualArrays(util.testSwap(TEST_DATA3, TEST_DATA1), TEST_DATA1);
        CHECK(result, true);

    }

    /*
    Testing the empty stack function
    */
    void testEmpty(){
        CHECK(util.testEmpty(TEST_DATA1), false);
        CHECK(util.testEmpty(TEST_DATA2), false);
        CHECK(util.testEmpty(TEST_DATA3), false);
        CHECK(util.testEmpty(std::vector<int>(0)), true);
    }

    /*
    Testing the size stack function
    */
    void testSize(){
        CHECK(util.testSize(TEST_DATA1), TEST_DATA1.size());
        CHECK(util.testSize(TEST_DATA2), TEST_DATA2.size());
        CHECK(util.testSize(TEST_DATA3), TEST_DATA3.size());
    }

    /*
    Testing the push stack function
    */
    void testPush(){
        bool result = util.assertEqualArrays(util.testPush(TEST_DATA1), TEST_DATA1);
        CHECK(result, true);
        
        result = util.assertEqualArrays(util.testPush(TEST_DATA2), TEST_DATA2);
        CHECK(result, true);

        result = util.assertEqualArrays(util.testPush(TEST_DATA3), TEST_DATA3);
        CHECK(result, true);
    }

    /*
    Testing the pop stack function
    */
    void testPop(){
        CHECK(util.assertEqualArrays(util.testPop(TEST_DATA1, 3), TEST_DATA_POP1), true); // -3
        CHECK(util.assertEqualArrays(util.testPop(TEST_DATA2, 2), TEST_DATA_POP2), true); // -2
        CHECK(util.assertEqualArrays(util.testPop(TEST_DATA3, 1), TEST_DATA_POP3), true); // -1
    }

    /*
    Testing the top stack function
    */
    void testTop(){
        CHECK(util.assertEqual(util.testTop(TEST_DATA1), TEST_DATA1.at(TEST_DATA1.size() - 1)), true);
        CHECK(util.assertEqual(util.testTop(TEST_DATA2), TEST_DATA2.at(TEST_DATA2.size() - 1)), true);
        CHECK(util.assertEqual(util.testTop(TEST_DATA3), TEST_DATA3.at(TEST_DATA3.size() - 1)), true);
    }

    void test_list(){
        testSwap();
        testEmpty();
        testSize();
        testPush();
        testPop();
        testTop();
        
    }
};



#endif