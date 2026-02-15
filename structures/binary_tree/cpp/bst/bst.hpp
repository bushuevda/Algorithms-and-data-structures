#ifndef BST_HPP
#define BST_HPP 1

#include <memory>


template <typename T, typename K>
class Bst{
public:
    T Key;
    std::unique_ptr<K> Data;
    std::unique_ptr<Bst <T, K>> Parent;
    std::unique_ptr<Bst <T, K>> Left;
    std::unique_ptr<Bst, <T, K>> Right;


    bool compare(T key);



    // Вставка элемента в дерево
    //public void Insert(T key, K data)
    void insert(T key, K data);

    // Рекурсивный обход дерева InOrder (left->root->right)
    // void in_order_recursive(Bst<T, K>? node, Action<T,K> visit);
    void in_order_recursive(Bst<T, K>? node, Action<T,K> visit);
    // Рекурсивный обход дерева PreOrder (root->left->right)
    //void pre_order_recursive(BST<T, K>? node, Action<T,K> visit )

    // Рекурсивный обход дерева PostOrder (left->right->root)
    // void post_order_recursive(BST<T, K>? node, Action<T,K> visit)


    // Удаление node с заданным value
    //BST<T, K>? remove_node(BST<T, K>? node, T key)


    // Найти node с минимальным значением (для правого поддерева)
    //private  BST<T, K>? find_min_node()

    // Удалить node с минимальным значением
    //private BST<T, K>? DeleteMin()


    // Поиск node с заданным значением
    //public BST<T, K>? FindNode(BST<T, K> node, T key)

};

template <typename T, typename K> 
bool Bst<int, K>::compare(int key){
    
};

template <typename T, typename K>
void Bst<T, K>::insert(T key, K data){

}


#endif