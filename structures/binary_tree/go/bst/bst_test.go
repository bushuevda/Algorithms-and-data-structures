package bst

import (
	"testing"
)

const (
	SIZE_NOT_EQUAL = -1
	ARRAYS_EQUAL   = -2
)

// Для хранения состояния проверки целостности ссылок между node при удалении
type stateCheckPtr struct {
	rootVisit int
}

//			  11
//		   /	 \
//		  /	      \
//		 4	  	   12
//		  \			 \
//	       5		  23
//	 		\		 /  \
//			 8 		16  27
//			  \		 \
//			   9	 20
var arr1 = []int{11, 12, 23, 4, 5, 16, 27, 8, 9, 20}

//					  41
//				    /	\
//				  32	49
//				  /
//				17
//				/ \
//			   8   23
//			/  	 \
//	   	   5	  12
//			\	  /
//			 6	 11
var arr2 = []int{41, 32, 17, 8, 5, 12, 23, 6, 49, 11}

//							5
//						   /  \
//						  4	   \
//						 /		\
//						3		10
//					   /	   /
//					  2		  9
//					 /		 /
//	     		   1		8
//						   /
//						  7
//						 /
//						6
var arr3 = []int{5, 4, 3, 2, 1, 10, 9, 8, 7, 6}

//			30
//		  /	   \
//		 /		\
//		20		40
//	   /  \    /  \
//	  19  21  39  41
var arr4 = []int{30, 40, 20, 39, 41, 21, 19}

//				 15
//			/	 	  \
//	       5		  16
//		/	\			\
//	   3	12			20
//			/ \			 /\
//		  10  13		18 23
//		 /
//		6
//		 \
//		  7
var arr5 = []int{15, 5, 16, 3, 12, 20, 10, 13, 18, 23, 6, 7}

// Компаратор для int
func comparator(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

// Функция-фабрика bst для тестирования
func createBST[T, K any](arr []T, compare func(a, b T) int, data K) *BST[T, K] {
	bst := &BST[T, K]{Key: arr[0], ComparatorKeys: compare}
	for _, element := range arr[1:] {
		bst.Insert(element, data)
	}
	return bst
}

// Функция сравнения двух массивов
func compareTwoArrays[T any](arr1 []T, arr2 []T, compare func(a, b T) int) int {

	if len(arr1) != len(arr2) {
		return SIZE_NOT_EQUAL
	}

	for index := range arr1 {
		if compare(arr1[index], arr2[index]) != EQUAL {
			return index
		}
	}

	return ARRAYS_EQUAL
}

// Функция вывода ошибки для сравнения массивов
func showStateTest[T any](t *testing.T, res int, nameArr string, arrRes []T, arrExcept []T) {
	switch {
	case res == SIZE_NOT_EQUAL:
		t.Errorf("size arrays except and actually by %s not equal. \nExcepted: %v\nActually: %v\n", nameArr, arrExcept, arrRes)
	case res >= 0:
		t.Errorf("%s elements arrays except and actually with index %d not equal. \nExcepted: %v\nActually: %v\n", nameArr, res, arrExcept, arrRes)
	}
}

// Функция вывода ошибки для поиска node в tree с типом int
func showTestFindInt(t *testing.T, actual *BST[int, int], except *BST[int, int], nameTest string) {

	if actual != except && except == nil {
		t.Errorf("error test %s:\nActual: %v\nExcept: %v", nameTest, actual, except)
	}

	if actual != nil && except != nil {
		if actual.Key != except.Key {
			t.Errorf("error test %s:\nActual: %d\nExcept: %d", nameTest, actual.Key, except.Key)
		}
	}
}

// Всп. функция для проверки целостности ссылок между node при удалении
func utilForCheckPtr(t *testing.T, tn *BST[int, int], nTest string) {
	stateCheck := &stateCheckPtr{}
	preorderForCheckPtr(tn, stateCheck, t, nTest)
}

// Всп. функция для проверки целостности ссылок - обход preorder
func preorderForCheckPtr[T, K any](tn *BST[T, K], stCheck *stateCheckPtr, t *testing.T, nTest string) {
	if tn != nil {
		visitForCheckPtr(t, tn, stCheck, nTest)
		preorderForCheckPtr(tn.left, stCheck, t, nTest)
		preorderForCheckPtr(tn.right, stCheck, t, nTest)
	}
}

// Всп. функция для проверки целостности ссылок - посещениу node
func visitForCheckPtr[T, K any](t *testing.T, tn *BST[T, K], stCheck *stateCheckPtr, nTest string) {

	if tn.parent == nil {
		stCheck.rootVisit += 1
		if stCheck.rootVisit > 1 {
			t.Errorf("test %s error check ptr, not root node with %v have not parent ", nTest, tn.Key)
		}
	} else if tn.parent.right != tn && tn.parent.left != tn {
		t.Errorf("test %s error check ptr, node with %v have not incorrect ptr \nExcept: %v | %v\nActually: %v ", nTest, tn.Key, &tn.parent.right, &tn.parent.left, &tn)
	}

}

// Тестирование обхода preorder
func TestPreOrder(t *testing.T) {

	bst1 := createBST[int, int](arr1, comparator, 1111)
	bst2 := createBST[int, int](arr2, comparator, 1111)
	bst3 := createBST[int, int](arr3, comparator, 1111)
	bst4 := createBST[int, int](arr4, comparator, 1111)
	bst5 := createBST[int, int](arr5, comparator, 1111)

	arrExcept1 := []int{11, 4, 5, 8, 9, 12, 23, 16, 20, 27}
	arrExcept2 := []int{41, 32, 17, 8, 5, 6, 12, 11, 23, 49}
	arrExcept3 := []int{5, 4, 3, 2, 1, 10, 9, 8, 7, 6}
	arrExcept4 := []int{30, 20, 19, 21, 40, 39, 41}
	arrExcept5 := []int{15, 5, 3, 12, 10, 6, 7, 13, 16, 20, 18, 23}

	var arrRes1 []int
	bst1.PreOrderRecursive(func(key int) { arrRes1 = append(arrRes1, key) })
	showStateTest(t, compareTwoArrays(arrRes1, arrExcept1, comparator), "arr1", arrRes1, arrExcept1)

	var arrRes2 []int
	bst2.PreOrderRecursive(func(key int) { arrRes2 = append(arrRes2, key) })
	showStateTest(t, compareTwoArrays(arrRes2, arrExcept2, comparator), "arr2", arrRes1, arrExcept1)

	var arrRes3 []int
	bst3.PreOrderRecursive(func(key int) { arrRes3 = append(arrRes3, key) })
	showStateTest(t, compareTwoArrays(arrRes3, arrExcept3, comparator), "arr3", arrRes1, arrExcept1)

	var arrRes4 []int
	bst4.PreOrderRecursive(func(key int) { arrRes4 = append(arrRes4, key) })
	showStateTest(t, compareTwoArrays(arrRes4, arrExcept4, comparator), "arr4", arrRes1, arrExcept1)

	var arrRes5 []int
	bst5.PreOrderRecursive(func(key int) { arrRes5 = append(arrRes5, key) })
	showStateTest(t, compareTwoArrays(arrRes5, arrExcept5, comparator), "arr5", arrRes1, arrExcept1)

}

// Тестирование обхода postorder
func TestPostOrder(t *testing.T) {
	bst1 := createBST[int, int](arr1, comparator, 1111)
	bst2 := createBST[int, int](arr2, comparator, 1111)
	bst3 := createBST[int, int](arr3, comparator, 1111)
	bst4 := createBST[int, int](arr4, comparator, 1111)
	bst5 := createBST[int, int](arr5, comparator, 1111)

	arrExcept1 := []int{9, 8, 5, 4, 20, 16, 27, 23, 12, 11}
	arrExcept2 := []int{6, 5, 11, 12, 8, 23, 17, 32, 49, 41}
	arrExcept3 := []int{1, 2, 3, 4, 6, 7, 8, 9, 10, 5}
	arrExcept4 := []int{19, 21, 20, 39, 41, 40, 30}

	arrExcept5 := []int{3, 7, 6, 10, 13, 12, 5, 18, 23, 20, 16, 15}

	var arrRes1 []int
	bst1.PostOrderRecursive(func(key int) { arrRes1 = append(arrRes1, key) })
	showStateTest(t, compareTwoArrays(arrRes1, arrExcept1, comparator), "arr1", arrRes1, arrExcept1)

	var arrRes2 []int
	bst2.PostOrderRecursive(func(key int) { arrRes2 = append(arrRes2, key) })
	showStateTest(t, compareTwoArrays(arrRes2, arrExcept2, comparator), "arr2", arrRes1, arrExcept1)

	var arrRes3 []int
	bst3.PostOrderRecursive(func(key int) { arrRes3 = append(arrRes3, key) })
	showStateTest(t, compareTwoArrays(arrRes3, arrExcept3, comparator), "arr3", arrRes1, arrExcept1)

	var arrRes4 []int
	bst4.PostOrderRecursive(func(key int) { arrRes4 = append(arrRes4, key) })
	showStateTest(t, compareTwoArrays(arrRes4, arrExcept4, comparator), "arr4", arrRes1, arrExcept1)

	var arrRes5 []int
	bst5.PostOrderRecursive(func(key int) { arrRes5 = append(arrRes5, key) })
	showStateTest(t, compareTwoArrays(arrRes5, arrExcept5, comparator), "arr5", arrRes1, arrExcept1)
}

// Тестирование обхода inorder
func TestInOrder(t *testing.T) {
	bst1 := createBST[int, int](arr1, comparator, 1111)
	bst2 := createBST[int, int](arr2, comparator, 1111)
	bst3 := createBST[int, int](arr3, comparator, 1111)
	bst4 := createBST[int, int](arr4, comparator, 1111)
	bst5 := createBST[int, int](arr5, comparator, 1111)

	arrExcept1 := []int{4, 5, 8, 9, 11, 12, 16, 20, 23, 27}
	arrExcept2 := []int{5, 6, 8, 11, 12, 17, 23, 32, 41, 49}
	arrExcept3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arrExcept4 := []int{19, 20, 21, 30, 39, 40, 41}
	arrExcept5 := []int{3, 5, 6, 7, 10, 12, 13, 15, 16, 18, 20, 23}

	var arrRes1 []int
	bst1.InOrderRecursive(func(key int) { arrRes1 = append(arrRes1, key) })
	showStateTest(t, compareTwoArrays(arrRes1, arrExcept1, comparator), "arr1", arrRes1, arrExcept1)

	var arrRes2 []int
	bst2.InOrderRecursive(func(key int) { arrRes2 = append(arrRes2, key) })
	showStateTest(t, compareTwoArrays(arrRes2, arrExcept2, comparator), "arr2", arrRes2, arrExcept2)

	var arrRes3 []int
	bst3.InOrderRecursive(func(key int) { arrRes3 = append(arrRes3, key) })
	showStateTest(t, compareTwoArrays(arrRes3, arrExcept3, comparator), "arr3", arrRes3, arrExcept3)

	var arrRes4 []int
	bst4.InOrderRecursive(func(key int) { arrRes4 = append(arrRes4, key) })
	showStateTest(t, compareTwoArrays(arrRes4, arrExcept4, comparator), "arr4", arrRes4, arrExcept4)

	var arrRes5 []int
	bst5.InOrderRecursive(func(key int) { arrRes5 = append(arrRes5, key) })
	showStateTest(t, compareTwoArrays(arrRes5, arrExcept5, comparator), "arr5", arrRes5, arrExcept5)
}

// Тестирование удаления node
func TestRemove(t *testing.T) {
	//Тестовые данные 1
	bst1 := createBST[int, int](arr1, comparator, 1111)
	//1_1
	arrExcept1_1 := []int{11, 4, 5, 8, 9, 12, 27, 16, 20}
	var arrRes1_1 []int
	bst1 = bst1.RemoveNode(23, comparator)
	bst1.PreOrderRecursive(func(key int) { arrRes1_1 = append(arrRes1_1, key) })
	showStateTest(t, compareTwoArrays(arrRes1_1, arrExcept1_1, comparator), "arr1_1", arrRes1_1, arrExcept1_1)
	utilForCheckPtr(t, bst1, "arr1_1")
	//1_2
	arrExcept1_2 := []int{11, 4, 5, 9, 12, 27, 16, 20}
	var arrRes1_2 []int
	bst1 = bst1.RemoveNode(8, comparator)
	bst1.PreOrderRecursive(func(key int) { arrRes1_2 = append(arrRes1_2, key) })
	showStateTest(t, compareTwoArrays(arrRes1_2, arrExcept1_2, comparator), "arr1_2", arrRes1_2, arrExcept1_2)
	utilForCheckPtr(t, bst1, "arr1_2")
	//1_3
	arrExcept1_3 := []int{11, 4, 5, 9, 12, 16, 20}
	var arrRes1_3 []int
	bst1 = bst1.RemoveNode(27, comparator)
	bst1.PreOrderRecursive(func(key int) { arrRes1_3 = append(arrRes1_3, key) })
	showStateTest(t, compareTwoArrays(arrRes1_3, arrExcept1_3, comparator), "arr1_3", arrRes1_3, arrExcept1_3)
	utilForCheckPtr(t, bst1, "arr1_3")
	//1_4
	arrExcept1_4 := []int{11, 4, 5, 9, 16, 20}
	var arrRes1_4 []int
	bst1 = bst1.RemoveNode(12, comparator)
	bst1.PreOrderRecursive(func(key int) { arrRes1_4 = append(arrRes1_4, key) })
	showStateTest(t, compareTwoArrays(arrRes1_4, arrExcept1_4, comparator), "arr1_4", arrRes1_4, arrExcept1_4)
	utilForCheckPtr(t, bst1, "arr1_4")
	//1_5
	arrExcept1_5 := []int{11, 5, 9, 16, 20}
	var arrRes1_5 []int
	bst1 = bst1.RemoveNode(4, comparator)
	bst1.PreOrderRecursive(func(key int) { arrRes1_5 = append(arrRes1_5, key) })
	showStateTest(t, compareTwoArrays(arrRes1_5, arrExcept1_5, comparator), "arr1_5", arrRes1_5, arrExcept1_5)
	utilForCheckPtr(t, bst1, "arr1_5")

	//Тестовые данные 2
	bst2 := createBST[int, int](arr2, comparator, 1111)
	//2_1
	arrExcept2_1 := []int{41, 32, 23, 8, 5, 6, 12, 11, 49}
	var arrRes2_1 []int
	bst2 = bst2.RemoveNode(17, comparator)
	bst2.PreOrderRecursive(func(key int) { arrRes2_1 = append(arrRes2_1, key) })
	showStateTest(t, compareTwoArrays(arrRes2_1, arrExcept2_1, comparator), "arr2_1", arrRes2_1, arrExcept2_1)
	utilForCheckPtr(t, bst2, "arr2_1")
	//2_2
	arrExcept2_2 := []int{41, 32, 23, 8, 5, 6, 11, 49}
	var arrRes2_2 []int
	bst2 = bst2.RemoveNode(12, comparator)
	bst2.PreOrderRecursive(func(key int) { arrRes2_2 = append(arrRes2_2, key) })
	showStateTest(t, compareTwoArrays(arrRes2_2, arrExcept2_2, comparator), "arr2_2", arrRes2_2, arrExcept2_2)
	utilForCheckPtr(t, bst2, "arr2_2")
	//2_3
	arrExcept2_3 := []int{41, 32, 23, 11, 5, 6, 49}
	var arrRes2_3 []int
	bst2 = bst2.RemoveNode(8, comparator)
	bst2.PreOrderRecursive(func(key int) { arrRes2_3 = append(arrRes2_3, key) })
	showStateTest(t, compareTwoArrays(arrRes2_3, arrExcept2_3, comparator), "arr2_3", arrRes2_3, arrExcept2_3)
	utilForCheckPtr(t, bst2, "arr2_3")
	//2_4
	arrExcept2_4 := []int{41, 32, 11, 5, 6, 49}
	var arrRes2_4 []int
	bst2 = bst2.RemoveNode(23, comparator)
	bst2.PreOrderRecursive(func(key int) { arrRes2_4 = append(arrRes2_4, key) })
	showStateTest(t, compareTwoArrays(arrRes2_4, arrExcept2_4, comparator), "arr2_4", arrRes2_4, arrExcept2_4)
	utilForCheckPtr(t, bst2, "arr2_4")
	//2_5
	arrExcept2_5 := []int{49, 32, 11, 5, 6}
	var arrRes2_5 []int
	bst2 = bst2.RemoveNode(41, comparator)
	bst2.PreOrderRecursive(func(key int) { arrRes2_5 = append(arrRes2_5, key) })
	showStateTest(t, compareTwoArrays(arrRes2_5, arrExcept2_5, comparator), "arr2_5", arrRes2_5, arrExcept2_5)
	utilForCheckPtr(t, bst2, "arr2_5")

	//Тестовые данные 3
	bst3 := createBST[int, int](arr3, comparator, 1111)
	//3_1
	arrExcept3_1 := []int{6, 4, 3, 2, 1, 10, 9, 8, 7}
	var arrRes3_1 []int
	bst3 = bst3.RemoveNode(5, comparator)
	bst3.PreOrderRecursive(func(key int) { arrRes3_1 = append(arrRes3_1, key) })
	showStateTest(t, compareTwoArrays(arrRes3_1, arrExcept3_1, comparator), "arr3_1", arrRes3_1, arrExcept3_1)
	utilForCheckPtr(t, bst3, "arr3_1")
	//3_2
	arrExcept3_2 := []int{6, 4, 3, 2, 10, 9, 8, 7}
	var arrRes3_2 []int
	bst3 = bst3.RemoveNode(1, comparator)
	bst3.PreOrderRecursive(func(key int) { arrRes3_2 = append(arrRes3_2, key) })
	showStateTest(t, compareTwoArrays(arrRes3_2, arrExcept3_2, comparator), "arr3_2", arrRes3_2, arrExcept3_2)
	utilForCheckPtr(t, bst3, "arr3_2")
	//3_3
	arrExcept3_3 := []int{6, 4, 2, 10, 9, 8, 7}
	var arrRes3_3 []int
	bst3 = bst3.RemoveNode(3, comparator)
	bst3.PreOrderRecursive(func(key int) { arrRes3_3 = append(arrRes3_3, key) })
	showStateTest(t, compareTwoArrays(arrRes3_3, arrExcept3_3, comparator), "arr3_3", arrRes3_3, arrExcept3_3)
	utilForCheckPtr(t, bst3, "arr3_3")
	//3_4
	arrExcept3_4 := []int{7, 4, 2, 10, 9, 8}
	var arrRes3_4 []int
	bst3 = bst3.RemoveNode(6, comparator)
	bst3.PreOrderRecursive(func(key int) { arrRes3_4 = append(arrRes3_4, key) })
	showStateTest(t, compareTwoArrays(arrRes3_4, arrExcept3_4, comparator), "arr3_4", arrRes3_4, arrExcept3_4)
	utilForCheckPtr(t, bst3, "arr3_4")
	//3_5
	arrExcept3_5 := []int{7, 4, 2, 9, 8}
	var arrRes3_5 []int
	bst3 = bst3.RemoveNode(10, comparator)
	bst3.PreOrderRecursive(func(key int) { arrRes3_5 = append(arrRes3_5, key) })
	showStateTest(t, compareTwoArrays(arrRes3_5, arrExcept3_5, comparator), "arr3_5", arrRes3_5, arrExcept3_5)
	utilForCheckPtr(t, bst3, "arr3_5")

	//Тестовые данные 4
	bst4 := createBST[int, int](arr4, comparator, 1111)
	//4_1
	arrExcept4_1 := []int{30, 21, 19, 40, 39, 41}
	var arrRes4_1 []int
	bst4 = bst4.RemoveNode(20, comparator)
	bst4.PreOrderRecursive(func(key int) { arrRes4_1 = append(arrRes4_1, key) })
	showStateTest(t, compareTwoArrays(arrRes4_1, arrExcept4_1, comparator), "arr4_1", arrRes4_1, arrExcept4_1)
	utilForCheckPtr(t, bst4, "arr4_1")
	//4_2
	arrExcept4_2 := []int{30, 21, 19, 41, 39}
	var arrRes4_2 []int
	bst4 = bst4.RemoveNode(40, comparator)
	bst4.PreOrderRecursive(func(key int) { arrRes4_2 = append(arrRes4_2, key) })
	showStateTest(t, compareTwoArrays(arrRes4_2, arrExcept4_2, comparator), "arr4_2", arrRes4_2, arrExcept4_2)
	utilForCheckPtr(t, bst4, "arr4_2")
	//4_3
	arrExcept4_3 := []int{39, 21, 19, 41}
	var arrRes4_3 []int
	bst4 = bst4.RemoveNode(30, comparator)
	bst4.PreOrderRecursive(func(key int) { arrRes4_3 = append(arrRes4_3, key) })
	showStateTest(t, compareTwoArrays(arrRes4_3, arrExcept4_3, comparator), "arr4_3", arrRes4_3, arrExcept4_3)
	utilForCheckPtr(t, bst4, "arr4_3")
	//4_4
	arrExcept4_4 := []int{41, 21, 19}
	var arrRes4_4 []int
	bst4 = bst4.RemoveNode(39, comparator)
	bst4.PreOrderRecursive(func(key int) { arrRes4_4 = append(arrRes4_4, key) })
	showStateTest(t, compareTwoArrays(arrRes4_4, arrExcept4_4, comparator), "arr4_4", arrRes4_4, arrExcept4_4)
	utilForCheckPtr(t, bst4, "arr4_4")
	//4_5
	arrExcept4_5 := []int{21, 19}
	var arrRes4_5 []int
	bst4 = bst4.RemoveNode(41, comparator)
	bst4.PreOrderRecursive(func(key int) { arrRes4_5 = append(arrRes4_5, key) })
	showStateTest(t, compareTwoArrays(arrRes4_5, arrExcept4_5, comparator), "arr4_5", arrRes4_5, arrExcept4_5)
	utilForCheckPtr(t, bst4, "arr4_5")
	//4_6
	arrExcept4_6 := []int{19}
	var arrRes4_6 []int
	bst4 = bst4.RemoveNode(21, comparator)
	bst4.PreOrderRecursive(func(key int) { arrRes4_6 = append(arrRes4_6, key) })
	showStateTest(t, compareTwoArrays(arrRes4_6, arrExcept4_6, comparator), "arr4_6", arrRes4_6, arrExcept4_6)
	utilForCheckPtr(t, bst4, "arr4_6")
	//4_7
	arrExcept4_7 := []int{}
	var arrRes4_7 []int
	bst4 = bst4.RemoveNode(19, comparator)
	bst4.PreOrderRecursive(func(key int) { arrRes4_7 = append(arrRes4_7, key) })
	showStateTest(t, compareTwoArrays(arrRes4_7, arrExcept4_7, comparator), "arr4_7", arrRes4_7, arrExcept4_7)
	utilForCheckPtr(t, bst4, "arr4_7")

	//Тестовые данные 5
	bst5 := createBST[int, int](arr5, comparator, 1111)
	//5_1
	arrExcept5_1 := []int{15, 5, 3, 12, 10, 6, 7, 13, 16, 23, 18}
	var arrRes5_1 []int
	bst5 = bst5.RemoveNode(20, comparator)
	bst5.PreOrderRecursive(func(key int) { arrRes5_1 = append(arrRes5_1, key) })
	showStateTest(t, compareTwoArrays(arrRes5_1, arrExcept5_1, comparator), "arr5_1", arrRes5_1, arrExcept5_1)
	utilForCheckPtr(t, bst5, "arr5_1")
	//5_2
	arrExcept5_2 := []int{16, 5, 3, 12, 10, 6, 7, 13, 23, 18}
	var arrRes5_2 []int
	bst5 = bst5.RemoveNode(15, comparator)
	bst5.PreOrderRecursive(func(key int) { arrRes5_2 = append(arrRes5_2, key) })
	showStateTest(t, compareTwoArrays(arrRes5_2, arrExcept5_2, comparator), "arr5_2", arrRes5_2, arrExcept5_2)
	utilForCheckPtr(t, bst5, "arr5_2")
	//5_3
	arrExcept5_3 := []int{16, 5, 3, 12, 10, 6, 7, 13, 18}
	var arrRes5_3 []int
	bst5 = bst5.RemoveNode(23, comparator)
	bst5.PreOrderRecursive(func(key int) { arrRes5_3 = append(arrRes5_3, key) })
	showStateTest(t, compareTwoArrays(arrRes5_3, arrExcept5_3, comparator), "arr5_3", arrRes5_3, arrExcept5_3)
	utilForCheckPtr(t, bst5, "arr5_3")
	//5_4
	arrExcept5_4 := []int{16, 5, 3, 12, 10, 6, 7, 13}
	var arrRes5_4 []int
	bst5 = bst5.RemoveNode(18, comparator)
	bst5.PreOrderRecursive(func(key int) { arrRes5_4 = append(arrRes5_4, key) })
	showStateTest(t, compareTwoArrays(arrRes5_4, arrExcept5_4, comparator), "arr5_4", arrRes5_4, arrExcept5_4)
	utilForCheckPtr(t, bst5, "arr5_4")
	//5_5
	arrExcept5_5 := []int{5, 3, 12, 10, 6, 7, 13}
	var arrRes5_5 []int
	bst5 = bst5.RemoveNode(16, comparator)
	bst5.PreOrderRecursive(func(key int) { arrRes5_5 = append(arrRes5_5, key) })
	showStateTest(t, compareTwoArrays(arrRes5_5, arrExcept5_5, comparator), "arr5_5", arrRes5_5, arrExcept5_5)
	utilForCheckPtr(t, bst5, "arr5_5")
}

// Тестирование поиска node
func TestFind(t *testing.T) {

	bst1 := createBST[int, int](arr1, comparator, 1111)
	bst2 := createBST[int, int](arr2, comparator, 1111)
	bst3 := createBST[int, int](arr3, comparator, 1111)
	bst4 := createBST[int, int](arr4, comparator, 1111)
	bst5 := createBST[int, int](arr5, comparator, 1111)

	keyActually1_1 := bst1.FindNode(5, comparator)
	keyActually1_2 := bst1.FindNode(16, comparator)
	keyActually1_3 := bst1.FindNode(111, comparator)

	keyExcept1_1 := &BST[int, int]{Key: 5}
	keyExcept1_2 := &BST[int, int]{Key: 16}
	var keyExcept1_3 *BST[int, int]

	showTestFindInt(t, keyActually1_1, keyExcept1_1, "1_1")
	showTestFindInt(t, keyActually1_2, keyExcept1_2, "1_2")
	showTestFindInt(t, keyActually1_3, keyExcept1_3, "1_3")

	keyActually2_1 := bst2.FindNode(17, comparator)
	keyActually2_2 := bst2.FindNode(12, comparator)
	keyActually2_3 := bst2.FindNode(223, comparator)

	keyExcept2_1 := &BST[int, int]{Key: 17}
	keyExcept2_2 := &BST[int, int]{Key: 12}
	var keyExcept2_3 *BST[int, int]

	showTestFindInt(t, keyActually2_1, keyExcept2_1, "2_1")
	showTestFindInt(t, keyActually2_2, keyExcept2_2, "2_2")
	showTestFindInt(t, keyActually2_3, keyExcept2_3, "2_3")

	keyActually3_1 := bst3.FindNode(4, comparator)
	keyActually3_2 := bst3.FindNode(6, comparator)
	keyActually3_3 := bst3.FindNode(344, comparator)

	keyExcept3_1 := &BST[int, int]{Key: 4}
	keyExcept3_2 := &BST[int, int]{Key: 6}
	var keyExcept3_3 *BST[int, int]

	showTestFindInt(t, keyActually3_1, keyExcept3_1, "3_1")
	showTestFindInt(t, keyActually3_2, keyExcept3_2, "3_2")
	showTestFindInt(t, keyActually3_3, keyExcept3_3, "3_3")

	keyActually4_1 := bst4.FindNode(20, comparator)
	keyActually4_2 := bst4.FindNode(41, comparator)
	keyActually4_3 := bst4.FindNode(541, comparator)

	keyExcept4_1 := &BST[int, int]{Key: 20}
	keyExcept4_2 := &BST[int, int]{Key: 41}
	var keyExcept4_3 *BST[int, int]

	showTestFindInt(t, keyActually4_1, keyExcept4_1, "3_1")
	showTestFindInt(t, keyActually4_2, keyExcept4_2, "3_2")
	showTestFindInt(t, keyActually4_3, keyExcept4_3, "3_3")

	keyActually5_1 := bst5.FindNode(13, comparator)
	keyActually5_2 := bst5.FindNode(16, comparator)
	keyActually5_3 := bst5.FindNode(611, comparator)

	keyExcept5_1 := &BST[int, int]{Key: 13}
	keyExcept5_2 := &BST[int, int]{Key: 16}
	var keyExcept5_3 *BST[int, int]

	showTestFindInt(t, keyActually5_1, keyExcept5_1, "5_1")
	showTestFindInt(t, keyActually5_2, keyExcept5_2, "5_2")
	showTestFindInt(t, keyActually5_3, keyExcept5_3, "5_3")

}
