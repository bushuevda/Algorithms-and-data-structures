#ifndef BST_HPP
#define BST_HPP 1

#include <memory>


template <typename T, typename K>
class Bst{
public:
    T key_;
    K data_;
    Bst <T, K>* parent_;
    std::unique_ptr<Bst <T, K>> left_ = nullptr;
    std::unique_ptr<Bst <T, K>> right_ = nullptr;
    Bst <T, K>(T key, K data, Bst<T,K>* parent = nullptr): key_(key), data_(data), parent_(parent){
    };
    
    void insert(T key, K data){
    if(this->key_ < key){
        if(!this->left_){
            this->left_ = std::unique_ptr<Bst<T, K>>(new Bst<T, K>(std::move(key), std::move(data)));
            this->left_ = 4;
        } else {
            
        }

    }
    
}

};






#endif