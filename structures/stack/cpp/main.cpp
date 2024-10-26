#include "stack/stack.hpp"
#include <iostream>

int main(){
    Stack<int> a;
    a.push(1);
    a.push(2);
    a.push(3);
    
    Stack<int> b;
    b.push(4);
    b.push(5);
    b.push(6);

    std::cout<<"before\n";
    a.show();
    b.show();
    a.swap(&b);
    std::cout<<"after\n";
    a.show();
    b.show();


}