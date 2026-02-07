package bst

//BST binary search tree переименовать

import (
	"errors"
)

const (
	EQUAL     = 0
	LEFT_MORE = 1
	LEFT_LESS = -1
)

type Comparator[T any] func(T, T) int

type BST[T, K any] struct {
	Key  T
	Data K
	//a > b return 1, a == b return 0, a < b return -1
	ComparatorKeys Comparator[T]
	parent         *BST[T, K]
	left           *BST[T, K]
	right          *BST[T, K]
}

// Вставка элемента в дерево
func (tn *BST[T, K]) Insert(key T, data K) error {

	if tn == nil {

		return errors.New("node is null")

	} else if tn.ComparatorKeys(tn.Key, key) == EQUAL {

		return errors.New("value exists")

	} else if tn.ComparatorKeys(tn.Key, key) == LEFT_MORE {

		if tn.left == nil {
			tn.left = &BST[T, K]{Key: key, Data: data, ComparatorKeys: tn.ComparatorKeys, parent: tn}
			return nil

		}

		return tn.left.Insert(key, data)

	} else if tn.ComparatorKeys(tn.Key, key) == LEFT_LESS {

		if tn.right == nil {

			tn.right = &BST[T, K]{Key: key, Data: data, ComparatorKeys: tn.ComparatorKeys, parent: tn}
			return nil

		}

		return tn.right.Insert(key, data)

	}
	return nil
}

// Рекурсивный обход дерева InOrder (left->root->right)
func (tr *BST[T, K]) InOrderRecursive(visit func(key T)) {
	if tr != nil {
		tr.left.InOrderRecursive(visit)
		visit(tr.Key)
		tr.right.InOrderRecursive(visit)
	}
}

// Рекурсивный обход дерева PreOrder (root->left->right)
func (tr *BST[T, K]) PreOrderRecursive(visit func(key T)) {
	if tr != nil {
		visit(tr.Key)
		tr.left.PreOrderRecursive(visit)
		tr.right.PreOrderRecursive(visit)
	}
}

// Рекурсивный обход дерева PostOrder (left->right->root)
func (tr *BST[T, K]) PostOrderRecursive(visit func(key T)) {
	if tr != nil {
		tr.left.PostOrderRecursive(visit)
		tr.right.PostOrderRecursive(visit)
		visit(tr.Key)
	}
}

// Поиск node с заданным значением
func (tr *BST[T, K]) FindNode(key T, compare func(a, b T) int) *BST[T, K] {
	if tr == nil {
		return nil
	}

	if compare(tr.Key, key) == EQUAL {
		return tr
	} else if compare(tr.Key, key) == LEFT_MORE {
		return tr.left.FindNode(key, compare)
	} else if compare(tr.Key, key) == LEFT_LESS {
		return tr.right.FindNode(key, compare)
	}
	return nil
}

// Удаление node с заданным value
func (tr *BST[T, K]) RemoveNode(key T, compare func(a, b T) int) *BST[T, K] {

	if tr == nil {
		return nil
	}

	if compare(tr.Key, key) == LEFT_MORE {
		tr.left = tr.left.RemoveNode(key, compare)
	} else if compare(tr.Key, key) == LEFT_LESS {
		tr.right = tr.right.RemoveNode(key, compare)
	} else {
		if tr.left == nil && tr.right == nil {
			return nil
		} else if tr.left == nil {
			return tr.right
		} else if tr.right == nil {
			return tr.left
		}
		successor := tr.right.findMinInNode()
		tr.Key = successor.Key
		tr.right = tr.right.deleteMin()
	}
	return tr
}

// Найти node с минимальным значением (для правого поддерева)
func (tr *BST[T, K]) findMinInNode() *BST[T, K] {
	if tr.left == nil {
		return tr
	}
	return tr.left.findMinInNode()
}

// Удалить node с минимальным значением
func (tr *BST[T, K]) deleteMin() *BST[T, K] {
	if tr == nil {
		return nil
	}

	if tr.left == nil {
		return tr.right
	}

	tr.left = tr.left.deleteMin()
	return tr
}
