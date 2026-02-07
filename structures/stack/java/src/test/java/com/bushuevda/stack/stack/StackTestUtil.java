package com.bushuevda.stack.stack;
import java.util.ArrayList;

public class StackTestUtil <T> {
    
    /*
    Helped function for testing the push stack function 
    */
    ArrayList<T> testPush(ArrayList<T> array){
        Stack<T> stack = new Stack<T>();
        for (T element : array) {
            stack.push(element);
        }
        return stack.getArray();
    }

    /*
    Helped function for testing the pop stack function 
    */
    ArrayList<T> testPop(ArrayList<T> array, Integer size_pop){
        Stack<T> stack = new Stack<T>();
        for (T element : array){
            stack.push(element);
        }
        for(int i = 0; i < size_pop; i++)
            stack.pop();
        return stack.getArray();
    }

    /*
    Helped function for testing the top stack function 
    */
    T testTop(ArrayList<T> array){
        Stack<T> stack = new Stack<T>();
        for (T element : array){
            stack.push(element);
        }
        return stack.top();
    }

    /*
    Helped function for testing the size stack function 
    */
    Integer testSize(ArrayList<T> array){
        Stack<T> stack = new Stack<T>();
        for (T element : array){
            stack.push(element);
        }
        return stack.size();
    }

    /*
    Helped function for testing the empty stack function 
    */
    Boolean testEmpty(ArrayList<T> array){
        Stack<T> stack = new Stack<T>();
        for (T element : array){
            stack.push(element);
        }
        return stack.empty();
    }

    /*
    Helped function for testing the swap stack function 
    */
    ArrayList<T> testSwap(ArrayList<T> array1, ArrayList<T> array2){
        Stack<T> stack1 = new Stack<T>();
        for (T element : array1){
            stack1.push(element);
        }

        Stack<T> stack2 = new Stack<T>();
        for (T element : array2){
            stack2.push(element);
        }
        stack1.swap(stack2);
        return stack1.getArray();
    }


}
