namespace tree;


enum CMP
{
    EQUAL = 0,
    LEFT_MORE = 1,
    LEFT_LESS = -1
}

public class BST<T, K> where T : IComparable<T>
{
    public T Key{ get; set; }

    public K Data{ get; set; }

    public BST<T, K>? Parent{ get; set; }

    public BST<T, K>? Left{ get; set; }

    public BST<T, K>? Right{ get; set; }


    public BST(T key, K data){
        this.Key = key;
        this.Data = data;
    }

   private BST(T key, K data, BST<T,K> parent){
        this.Key = key;
        this.Data = data;
        this.Parent = parent;
    }

    // Вставка элемента в дерево
    public void Insert(T key, K data) {
        if (this.Key.CompareTo(key) == ((int)CMP.LEFT_MORE)){
            if (this.Left is null){
                this.Left = new BST<T, K>(key, data, this);
            } else {
                this.Left.Insert(key, data);
            }
        } else if (this.Key.CompareTo(key) == ((int)CMP.LEFT_LESS)){
            if (this.Right is null){
                this.Right = new BST<T, K>(key, data, this);
            } else {
                this.Right.Insert(key, data);
            }
        }
    }
    
    // Рекурсивный обход дерева InOrder (left->root->right)
    public void InOrderRecursive(BST<T, K>? node, Action<T,K> visit ){
        if (node is not null){
            this.InOrderRecursive(node.Left, visit);
            visit(node.Key, node.Data);
            this.InOrderRecursive(node.Right, visit);
        }
    }

    // Рекурсивный обход дерева PreOrder (root->left->right)
    public void PreOrderRecursive(BST<T, K>? node, Action<T,K> visit ){
        if (node is not null){
            visit(node.Key, node.Data);
            this.PreOrderRecursive(node.Left, visit);
            this.PreOrderRecursive(node.Right, visit);
        }
    }

    // Рекурсивный обход дерева PostOrder (left->right->root)
    public void PostOrderRecursive(BST<T, K>? node, Action<T,K> visit){
        if (node is not null){
            this.PostOrderRecursive(node.Left, visit);
            this.PostOrderRecursive(node.Right, visit);            
            visit(node.Key, node.Data);
        }
    }

    // Удаление node с заданным value
    public BST<T, K>? RemoveNode(BST<T, K>? node, T key){
        if(node is null){
            return null;
        }

        if (node.Key.CompareTo(key) == ((int)CMP.LEFT_MORE)){
            node.Left = node!.Left!.RemoveNode(node.Left, key);
        } else if (node.Key.CompareTo(key) == ((int)CMP.LEFT_LESS)){
            node.Right = node!.Right!.RemoveNode(node.Right, key);
        } else {
            if(node.Left is null && node.Right is null){
                return null;
            } else if (node.Left is null){
                return node.Right;
            } else if (node.Right is null){
                return node.Left;
            }

            var successor = node.Right.FindMinNode();
            node.Key = successor!.Key;
            node.Right = node.Right.DeleteMin();
        }
        return node;
    }

    // Найти node с минимальным значением (для правого поддерева)
    private BST<T, K>? FindMinNode(){
        if (this.Left is null) {
            return this;
        }
        return this.Left.FindMinNode();
    }

    // Удалить node с минимальным значением
    private BST<T, K>? DeleteMin(){
        if (this is null){
            return null;
        }

        if (this.Left is null){
            return this.Right;
        }

        this.Left = this.Left.DeleteMin();
        return this;
    }

    // Поиск node с заданным значением
    public BST<T, K>? FindNode(BST<T, K> node, T key){
        if(node is null){
            return null;
        }
        if(node.Key.CompareTo(key) == ((int)CMP.EQUAL)){
            return node;
        } else if(node.Key.CompareTo(key) == ((int)CMP.LEFT_MORE)){
            return FindNode(node.Left, key);
        } else if(node.Key.CompareTo(key) == ((int)CMP.LEFT_LESS)){
            return FindNode(node.Right, key);
        }

        return null;
    }

}
