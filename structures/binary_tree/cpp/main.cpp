#include "bst/bst.hpp"
#include <iostream>

int main() {
    auto bst = std::unique_ptr<Bst<int, int>>(new Bst<int, int>(5, 1));
    bst->insert(4, 1);
    bst->insert(9, 1);
    bst->insert(3, 1);
    bst->insert(10, 1);
    bst->insert(11, 1);
    bst->post_order_recursive(bst, [](int key, int data){
        std::cout<<key<<" ---- " <<data<<"\n";
    });
    bst = bst->remove_node(bst, 11);
    bst->post_order_recursive(bst, [](int key, int data){
        std::cout<<key<<" ---- " <<data<<"\n";
    });


    std::cout<<bst->find_node(bst, 10)->key_;



    return  0;
}