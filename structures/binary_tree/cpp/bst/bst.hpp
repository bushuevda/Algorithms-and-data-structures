#ifndef BST_HPP
#define BST_HPP 1

#include <memory>



template <typename T, typename K>
class Bst{

private:

    // Найти node с минимальным значением (для правого поддерева, используется в remove_node)
    Bst<T, K>* fin_min_node(){
        if (!this->left_) {
            return this;
        }
        return this->left_->fin_min_node();
    }


    //используется в remove_node
    std::unique_ptr<Bst<T,K>> delete_min(std::unique_ptr<Bst<T, K>> bst){
        if(!bst){
            return nullptr;
        }

        if(!bst->left_){
            return std::move(bst->right_);
        }

        bst->left_ = delete_min(std::move(bst->right_));

        return std::move(bst);
    }

public:
    T key_;
    K data_;
    Bst <T, K>* parent_;
    std::unique_ptr<Bst <T, K>> left_ = nullptr;
    std::unique_ptr<Bst <T, K>> right_ = nullptr;
    Bst <T, K>(T key, K data, Bst<T,K>* parent = nullptr): key_(key), data_(data), parent_(parent){};
    


    // Удаление node с заданным key
    std::unique_ptr<Bst<T,K>> remove_node(std::unique_ptr<Bst<T,K>>& node, T key){
        if(!node){
            return nullptr;
        }

        if (node->key_ > key){
            node->left_ = std::move(remove_node(node->left_, key));
        } else if (node->key_ < key){
            node->right_ = std::move(remove_node(node->right_, key));
        } else {
            if(!node->left_ && !node->right_){
                return std::move(node->left_);
            } 
            if (!node->left_){
                return std::move(node->right_);
            } 
            if (!node->right_){
                return std::move(node->left_);
            }

            auto successor = node->right_->fin_min_node();
            node->key_ = successor->key_;
            node->right_ = std::move(delete_min(std::move(node->right_)));
        }
        return std::move(node);
    }



    //Поиск node с заданным значением

    Bst<T, K>* find_node(std::unique_ptr<Bst<T, K>>& node, T key){
        if(!node){
            return nullptr;
        }
        if(node->key_ == key){
            return node.get();
        } else if(node->key_ > key){
            return find_node(node->left_, key);
        } else if (node->key_ < key) {
            return find_node(node->right_, key);
        }
        return nullptr;
    }
    
    // Вставка элемента в дерево
    void insert(T key, K data){
        if(this->key_ > key){
            if(!this->left_){
                this->left_ = std::unique_ptr<Bst<T, K>>(new Bst<T, K>(std::move(key), std::move(data), this));
            } else {
                this->left_->insert(key, data);
            }

        } else if(this->key_ < key){
            if(!this->right_){
                this->right_ = std::unique_ptr<Bst<T, K>>(new Bst<T, K>(std::move(key), std::move(data), this));
            } else {
                this->right_->insert(key, data);
            }
        }
        
    }


    //Рекурсивный обход дерева InOrder (left->root->right)
    void in_order_recursive(std::unique_ptr<Bst<T, K>>& bst, void visit(T,K)){
        if(bst){
            this->in_order_recursive(bst->left_, visit);
            visit(bst->key_, bst->data_);
            this->in_order_recursive(bst->right_, visit);
        }
    }

    // Рекурсивный обход дерева PreOrder (root->left->right)
    void pre_order_recursive(std::unique_ptr<Bst<T, K>>& bst, void visit(T,K)){
        if(bst){
            visit(bst->key_, bst->data_);
            this->in_order_recursive(bst->left_, visit);
            this->in_order_recursive(bst->right_, visit);
        }
    }


    // Рекурсивный обход дерева PostOrder (left->right->root)
    void post_order_recursive(std::unique_ptr<Bst<T, K>>& bst, void visit(T,K)){
        if(bst){
            this->in_order_recursive(bst->left_, visit);
            this->in_order_recursive(bst->right_, visit);
            visit(bst->key_, bst->data_);
        }
    }




};


template class Bst<int, int>; 



#endif