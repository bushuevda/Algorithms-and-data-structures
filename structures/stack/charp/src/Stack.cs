using System.Collections;
using System.Collections.Generic;
using System.Runtime.CompilerServices;

namespace Structure{
   
    public class Stack <T>{

    protected LinkedList<T> array;
    
    public Stack() {
        array = new LinkedList<T>();
    }
   
    /*
    A function that returns a stack array
    */
    public LinkedList<T> getArray(){
        return array;
    }

    /*
    A function that set a stack array
    */
    public void setArray(LinkedList<T> new_array){
        array = new_array;
    }

    /*
    A function that swaps the stack arrays
    */
    public void swap(Stack<T> s2){
        LinkedList<T> temp = new LinkedList<T>();
        foreach(T element in s2.array){
            temp.AddLast(element);
        }
        s2.array.Clear();
        foreach(T element in array){
            s2.array.AddLast(element);
        }
        array.Clear();
        foreach(T element in temp){
            array.AddLast(element);
        }
    }

    /*
    A function that returns the top element of the stack 
    */
    public T? top(){
        if(!empty())
            return array.Last();
        return default(T);
    }

    /*
    A function that adds an element to the stack
    */
    public void push(T element){
        array.AddLast(element);
   }
   
    /*
    A function that removes a top element from the stack and returns it
    */
    public T? pop(){
        
        if (!empty()){
            T temp = array.Last();
            array.RemoveLast();
            return temp;
        }
        return default(T);
    }

    /*
    A function that returns stack size
    */
    public int size(){
        return array.Count;
    }

    /*
    A function that checks the stack for emptiness
    */
    public bool empty(){
        return size() > 0 ? false : true;
    }

    /*
    A function that shows the stack arrays
    */
    public void show(){
        if(!empty())
            foreach(T element in array){
                Console.Write(element + " ");
            }
        else
            Console.WriteLine("Stack empty");
    }
}
}
