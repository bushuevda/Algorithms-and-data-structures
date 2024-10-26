#ifndef STACK_HPP
#define STACK_HPP 1
#include <vector>
#include <iostream>

template <typename T> class Stack{
private:
    std::vector<T> array;
public:
    Stack();
    void push(T element);
    T pop();
    bool empty();
    int size();
    T top();
    void swap(Stack* st);
    void show();
    std::vector<T> getArray(){return this->array;}
    void setArray(std::vector<T> new_array){this->array = new_array;}
};

template <typename T> Stack<T>::Stack(){};

/*
A function that returns stack size
*/
template <typename T> int Stack<T>::size(){
    return this->array.size();
}

/*
A function that checks the stack for emptiness
*/
template <typename T> bool Stack<T>::empty(){
    if (this->size() <= 0)
        return true;
    else 
        return false;
};

/*
A function that adds an element to the stack
*/
template <typename T> void Stack<T>::push(T element){
    this->array.push_back(element);
}

/*
A function that removes a top element from the stack and returns it
*/
template <typename T> T Stack<T>::pop(){
    if(!this->empty()){
        T temp = this->array[this->array.size() - 1];
        this->array.pop_back();
        return temp;
    }
    else {
        T temp;
        return temp;
    }
}

/*
A function that returns the top element of the stack 
*/
template <typename T> T Stack<T>::top(){
    if(!this->empty())
        return this->array[this->array.size() - 1];
    else {
        T temp;
        return temp;
    }
}

/*
A function that swaps the stack arrays
*/
template <typename T> void Stack<T>::swap(Stack* st){
    std::vector<T> temp = this->array;
    this->array = st->getArray();
    st->setArray(temp);
}


/*
A function that shows the stack arrays
*/
template <typename T> void Stack<T>::show(){
    std::cout<<"[";
    for(auto elem : this->array){
        std::cout<<elem<<", ";
    }
    std::cout<<"]\n";
}

#endif