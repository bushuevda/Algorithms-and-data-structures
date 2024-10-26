package com.bushuevda.stack.stack;

import java.util.ArrayList;

public class Stack<T>{
    ArrayList<T> array;

    public Stack(){
        array = new ArrayList<T>();
    }

    /*
    A function that returns a stack array
    */
    public ArrayList<T> getArray() {
        return array;
    }

    /*
    A function that set a stack array
    */
    private void setArray(ArrayList<T> array){
        this.array = array;
    }

    /*
    A function that swaps the stack arrays
    */
    public void swap(Stack<T> stack){
        ArrayList<T> temp = stack.getArray();
        stack.setArray(this.array);
        this.array = temp;
    }

    /*
    A function that returns stack size
    */
    public Integer size(){
        return this.array.size();
    };

    /*
    A function that checks the stack for emptiness
    */
    public Boolean empty(){
        if(this.size() > 0)
            return false;
        else
            return true;
    }

    /*
    A function that removes a top element from the stack and returns it
    */
    public T pop(){
        T temp = this.array.get(this.array.size() - 1);
        this.array.remove(this.array.size() - 1);
        return temp;
    }

    /*
    A function that returns the top element of the stack 
    */
    public T top(){
        return this.array.get(this.array.size() - 1);
    }

    /*
    A function that adds an element to the stack
    */
    public void push(T element){
        this.array.add(element);
    }
    
    /*
    A function that shows the stack arrays
    */
    public void show(){
        System.err.println(this.array);
    }

}
