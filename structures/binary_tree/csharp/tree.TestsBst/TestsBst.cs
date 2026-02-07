using System.Collections;
using tree;
using Xunit;


namespace tree.TestsBst;


public class TestsBst
{


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
    public List<int> arr1 = [11, 12, 23, 4, 5, 16, 27, 8, 9, 20];

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
    public List<int> arr2 = [41, 32, 17, 8, 5, 12, 23, 6, 49, 11];

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
    public List<int> arr3 = [5, 4, 3, 2, 1, 10, 9, 8, 7, 6];

    //			30
    //		  /	   \
    //		 /		\
    //		20		40
    //	   /  \    /  \
    //	  19  21  39  41
    public List<int> arr4 = [30, 40, 20, 39, 41, 21, 19];

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
    public List<int> arr5 = [15, 5, 16, 3, 12, 20, 10, 13, 18, 23, 6, 7];

    // Функция-фабрика bst для тестирования
    public BST<int, int> CreateIntBST(List<int> arr, int data){
        BST<int, int> bst = new BST<int, int>(arr[0], data);
        foreach (var element in arr.Slice(1, arr.Count - 1)){
            bst.Insert(element, data);
        }
        return bst;
    }

    // Всп. функция для проверки целостности ссылок между node при удалении
    public void UtilForCheckPtr(BST<int, int> node){
        int rootVisit = 0;
        List<bool> resultsCheck = [];
        void visitForCheckPtr(BST<int, int> node)
        {
            if (node.Parent is null)
            {
                rootVisit += 1;
                if (rootVisit > 1)
                {
                    resultsCheck.Add(false);
                }
            }
            else if (!ReferenceEquals(node.Parent.Right, node) && !ReferenceEquals(node.Parent.Left, node))
            {
                resultsCheck.Add(false);
            }
        }

        void PreorderForCheckPtr(BST<int, int> node){
            if(node is not null){
                visitForCheckPtr(node);
                PreorderForCheckPtr(node.Left!);
                PreorderForCheckPtr(node.Right!);
            }
        }

        PreorderForCheckPtr(node);

        Assert.DoesNotContain(false, resultsCheck);
    }


    [Fact]
    // Тестирование обхода preorder
    public void TestPreOrder()
    {
        BST<int, int> bst1 = CreateIntBST(arr1, 111);
        BST<int, int> bst2 = CreateIntBST(arr2, 111);
        BST<int, int> bst3 = CreateIntBST(arr3, 111);
        BST<int, int> bst4 = CreateIntBST(arr4, 111);
        BST<int, int> bst5 = CreateIntBST(arr5, 111);

        List<int> arrExcept1 = [11, 4, 5, 8, 9, 12, 23, 16, 20, 27];
        List<int> arrExcept2 = [41, 32, 17, 8, 5, 6, 12, 11, 23, 49];
        List<int> arrExcept3 = [5, 4, 3, 2, 1, 10, 9, 8, 7, 6];
        List<int> arrExcept4 = [30, 20, 19, 21, 40, 39, 41];
        List<int> arrExcept5 = [15, 5, 3, 12, 10, 6, 7, 13, 16, 20, 18, 23];

        List<int> arrRes1 = [];
        List<int> arrRes2 = [];
        List<int> arrRes3 = [];
        List<int> arrRes4 = [];
        List<int> arrRes5 = [];

        bst1.PreOrderRecursive(bst1, (int key, int _) => arrRes1.Add(key));
        bst2.PreOrderRecursive(bst2, (int key, int _) => arrRes2.Add(key));
        bst3.PreOrderRecursive(bst3, (int key, int _) => arrRes3.Add(key));
        bst4.PreOrderRecursive(bst4, (int key, int _) => arrRes4.Add(key));
        bst5.PreOrderRecursive(bst5, (int key, int _) => arrRes5.Add(key));

        Assert.Equal(arrExcept1, arrRes1);
        Assert.Equal(arrExcept2, arrRes2);
        Assert.Equal(arrExcept3, arrRes3);
        Assert.Equal(arrExcept4, arrRes4);
        Assert.Equal(arrExcept5, arrRes5);

    }

    [Fact]
    // Тестирование обхода postorder
    public void TestPostOrder()
    {
        BST<int, int> bst1 = CreateIntBST(arr1, 111);
        BST<int, int> bst2 = CreateIntBST(arr2, 111);
        BST<int, int> bst3 = CreateIntBST(arr3, 111);
        BST<int, int> bst4 = CreateIntBST(arr4, 111);
        BST<int, int> bst5 = CreateIntBST(arr5, 111);

        List<int> arrExcept1 = [9, 8, 5, 4, 20, 16, 27, 23, 12, 11];
        List<int> arrExcept2 = [6, 5, 11, 12, 8, 23, 17, 32, 49, 41];
        List<int> arrExcept3 = [1, 2, 3, 4, 6, 7, 8, 9, 10, 5];
        List<int> arrExcept4 = [19, 21, 20, 39, 41, 40, 30];
        List<int> arrExcept5 = [3, 7, 6, 10, 13, 12, 5, 18, 23, 20, 16, 15];

        List<int> arrRes1 = [];
        List<int> arrRes2 = [];
        List<int> arrRes3 = [];
        List<int> arrRes4 = [];
        List<int> arrRes5 = [];

        bst1.PostOrderRecursive(bst1, (int key, int _) => arrRes1.Add(key));
        bst2.PostOrderRecursive(bst2, (int key, int _) => arrRes2.Add(key));
        bst3.PostOrderRecursive(bst3, (int key, int _) => arrRes3.Add(key));
        bst4.PostOrderRecursive(bst4, (int key, int _) => arrRes4.Add(key));
        bst5.PostOrderRecursive(bst5, (int key, int _) => arrRes5.Add(key));

        Assert.Equal(arrExcept1, arrRes1);
        Assert.Equal(arrExcept2, arrRes2);
        Assert.Equal(arrExcept3, arrRes3);
        Assert.Equal(arrExcept4, arrRes4);
        Assert.Equal(arrExcept5, arrRes5);

    }

    [Fact]
    // Тестирование обхода inorder
    public void TestInOrder()
    {
        BST<int, int> bst1 = CreateIntBST(arr1, 111);
        BST<int, int> bst2 = CreateIntBST(arr2, 111);
        BST<int, int> bst3 = CreateIntBST(arr3, 111);
        BST<int, int> bst4 = CreateIntBST(arr4, 111);
        BST<int, int> bst5 = CreateIntBST(arr5, 111);

        List<int> arrExcept1 = [4, 5, 8, 9, 11, 12, 16, 20, 23, 27];
        List<int> arrExcept2 = [5, 6, 8, 11, 12, 17, 23, 32, 41, 49];
        List<int> arrExcept3 = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
        List<int> arrExcept4 = [19, 20, 21, 30, 39, 40, 41];
        List<int> arrExcept5 = [3, 5, 6, 7, 10, 12, 13, 15, 16, 18, 20, 23];

        List<int> arrRes1 = [];
        List<int> arrRes2 = [];
        List<int> arrRes3 = [];
        List<int> arrRes4 = [];
        List<int> arrRes5 = [];

        bst1.InOrderRecursive(bst1, (int key, int _) => arrRes1.Add(key));
        bst2.InOrderRecursive(bst2, (int key, int _) => arrRes2.Add(key));
        bst3.InOrderRecursive(bst3, (int key, int _) => arrRes3.Add(key));
        bst4.InOrderRecursive(bst4, (int key, int _) => arrRes4.Add(key));
        bst5.InOrderRecursive(bst5, (int key, int _) => arrRes5.Add(key));

        Assert.Equal(arrExcept1, arrRes1);
        Assert.Equal(arrExcept2, arrRes2);
        Assert.Equal(arrExcept3, arrRes3);
        Assert.Equal(arrExcept4, arrRes4);
        Assert.Equal(arrExcept5, arrRes5);

    }

    [Fact]
    // Тестирование удаления node
    public void TestRemove(){
        BST<int, int> bst1 = CreateIntBST(arr1, 111);
        BST<int, int> bst2 = CreateIntBST(arr2, 111);
        BST<int, int> bst3 = CreateIntBST(arr3, 111);
        BST<int, int> bst4 = CreateIntBST(arr4, 111);
        BST<int, int> bst5 = CreateIntBST(arr5, 111);

        List<int> arrExcept1_1 = [11, 4, 5, 8, 9, 12, 27, 16, 20];
        List<int> arrExcept1_2 = [11, 4, 5, 9, 12, 27, 16, 20];
        List<int> arrExcept1_3 = [11, 4, 5, 9, 12, 16, 20];
        List<int> arrExcept1_4 = [11, 4, 5, 9, 16, 20];
        List<int> arrExcept1_5 = [11, 5, 9, 16, 20];

        List<int> arrExcept2_1 = [41, 32, 23, 8, 5, 6, 12, 11, 49];
        List<int> arrExcept2_2 = [41, 32, 23, 8, 5, 6, 11, 49];
        List<int> arrExcept2_3 = [41, 32, 23, 11, 5, 6, 49];
        List<int> arrExcept2_4 = [41, 32, 11, 5, 6, 49];
        List<int> arrExcept2_5 = [49, 32, 11, 5, 6];

        List<int> arrExcept3_1 = [6, 4, 3, 2, 1, 10, 9, 8, 7];
        List<int> arrExcept3_2 = [6, 4, 3, 2, 10, 9, 8, 7];
        List<int> arrExcept3_3 = [6, 4, 2, 10, 9, 8, 7];
        List<int> arrExcept3_4 = [7, 4, 2, 10, 9, 8];
        List<int> arrExcept3_5 = [7, 4, 2, 9, 8];

        List<int> arrExcept4_1 = [30, 21, 19, 40, 39, 41];
        List<int> arrExcept4_2 = [30, 21, 19, 41, 39];
        List<int> arrExcept4_3 = [39, 21, 19, 41];
        List<int> arrExcept4_4 = [41, 21, 19];
        List<int> arrExcept4_5 = [21, 19];
        List<int> arrExcept4_6 = [19];
        List<int> arrExcept4_7 = [];


        List<int> arrExcept5_1 = [15, 5, 3, 12, 10, 6, 7, 13, 16, 23, 18];
        List<int> arrExcept5_2 = [16, 5, 3, 12, 10, 6, 7, 13, 23, 18];
        List<int> arrExcept5_3 = [16, 5, 3, 12, 10, 6, 7, 13, 18];
        List<int> arrExcept5_4 = [16, 5, 3, 12, 10, 6, 7, 13];
        List<int> arrExcept5_5 = [5, 3, 12, 10, 6, 7, 13];


        List<int> arrRes1_1 = [];
        List<int> arrRes1_2 = [];
        List<int> arrRes1_3 = [];
        List<int> arrRes1_4 = [];
        List<int> arrRes1_5 = [];

        List<int> arrRes2_1 = [];
        List<int> arrRes2_2 = [];
        List<int> arrRes2_3 = [];
        List<int> arrRes2_4 = [];
        List<int> arrRes2_5 = [];

        List<int> arrRes3_1 = [];
        List<int> arrRes3_2 = [];
        List<int> arrRes3_3 = [];
        List<int> arrRes3_4 = [];
        List<int> arrRes3_5 = [];

        List<int> arrRes4_1 = [];
        List<int> arrRes4_2 = [];
        List<int> arrRes4_3 = [];
        List<int> arrRes4_4 = [];
        List<int> arrRes4_5 = [];
        List<int> arrRes4_6 = [];
        List<int> arrRes4_7 = [];


        List<int> arrRes5_1 = [];
        List<int> arrRes5_2 = [];
        List<int> arrRes5_3 = [];
        List<int> arrRes5_4 = [];
        List<int> arrRes5_5 = [];

        bst1 = bst1.RemoveNode(bst1, 23)!;
        bst1.PreOrderRecursive(bst1, (int key, int _) => arrRes1_1.Add(key));
        UtilForCheckPtr(bst1);
        bst1 = bst1.RemoveNode(bst1, 8)!;
        bst1.PreOrderRecursive(bst1, (int key, int _) => arrRes1_2.Add(key));
        UtilForCheckPtr(bst1);
        bst1 = bst1.RemoveNode(bst1, 27)!;
        bst1.PreOrderRecursive(bst1, (int key, int _) => arrRes1_3.Add(key));
        UtilForCheckPtr(bst1);
        bst1 = bst1.RemoveNode(bst1, 12)!;
        bst1.PreOrderRecursive(bst1, (int key, int _) => arrRes1_4.Add(key));
        UtilForCheckPtr(bst1);
        bst1 = bst1.RemoveNode(bst1, 4)!;
        bst1.PreOrderRecursive(bst1, (int key, int _) => arrRes1_5.Add(key));
        UtilForCheckPtr(bst1);
       
        bst2 = bst2.RemoveNode(bst2, 17)!;
        bst2.PreOrderRecursive(bst2, (int key, int _) => arrRes2_1.Add(key));
        UtilForCheckPtr(bst2);
        bst2 = bst2.RemoveNode(bst2, 12)!;
        bst2.PreOrderRecursive(bst2, (int key, int _) => arrRes2_2.Add(key));
        UtilForCheckPtr(bst2);
        bst2 = bst2.RemoveNode(bst2, 8)!;
        bst2.PreOrderRecursive(bst2, (int key, int _) => arrRes2_3.Add(key));
        UtilForCheckPtr(bst2);
        bst2 = bst2.RemoveNode(bst2, 23)!;
        bst2.PreOrderRecursive(bst2, (int key, int _) => arrRes2_4.Add(key));
        UtilForCheckPtr(bst2);
        bst2 = bst2.RemoveNode(bst2, 41)!;
        bst2.PreOrderRecursive(bst2, (int key, int _) => arrRes2_5.Add(key));
        UtilForCheckPtr(bst2);

        bst3 = bst3.RemoveNode(bst3, 5)!;
        bst3.PreOrderRecursive(bst3, (int key, int _) => arrRes3_1.Add(key));
        UtilForCheckPtr(bst3);
        bst3 = bst3.RemoveNode(bst3, 1)!;
        bst3.PreOrderRecursive(bst3, (int key, int _) => arrRes3_2.Add(key));
        UtilForCheckPtr(bst3);
        bst3 = bst3.RemoveNode(bst3, 3)!;
        bst3.PreOrderRecursive(bst3, (int key, int _) => arrRes3_3.Add(key));
        UtilForCheckPtr(bst3);
        bst3 = bst3.RemoveNode(bst3, 6)!;
        bst3.PreOrderRecursive(bst3, (int key, int _) => arrRes3_4.Add(key));
        UtilForCheckPtr(bst3);
        bst3 = bst3.RemoveNode(bst3, 10)!;
        bst3.PreOrderRecursive(bst3, (int key, int _) => arrRes3_5.Add(key));
        UtilForCheckPtr(bst3);

        bst4 = bst4.RemoveNode(bst4, 20)!;
        bst4.PreOrderRecursive(bst4, (int key, int _) => arrRes4_1.Add(key));
        UtilForCheckPtr(bst4);    
        bst4 = bst4.RemoveNode(bst4, 40)!;
        bst4.PreOrderRecursive(bst4, (int key, int _) => arrRes4_2.Add(key));
        UtilForCheckPtr(bst4);
        bst4 = bst4.RemoveNode(bst4, 30)!;
        bst4.PreOrderRecursive(bst4, (int key, int _) => arrRes4_3.Add(key));
        UtilForCheckPtr(bst4);
        bst4 = bst4.RemoveNode(bst4, 39)!;
        bst4.PreOrderRecursive(bst4, (int key, int _) => arrRes4_4.Add(key));
        UtilForCheckPtr(bst4);
        bst4 = bst4.RemoveNode(bst4, 41)!;
        bst4.PreOrderRecursive(bst4, (int key, int _) => arrRes4_5.Add(key));
        UtilForCheckPtr(bst4);
        bst4 = bst4.RemoveNode(bst4, 21)!;
        bst4.PreOrderRecursive(bst4, (int key, int _) => arrRes4_6.Add(key));
        UtilForCheckPtr(bst4);
        bst4 = bst4.RemoveNode(bst4, 19)!;


        bst5 = bst5.RemoveNode(bst5, 20)!;
        bst5.PreOrderRecursive(bst5, (int key, int _) => arrRes5_1.Add(key));
        UtilForCheckPtr(bst5);
        bst5 = bst5.RemoveNode(bst5, 15)!;
        bst5.PreOrderRecursive(bst5, (int key, int _) => arrRes5_2.Add(key));
        UtilForCheckPtr(bst5);
        bst5 = bst5.RemoveNode(bst5, 23)!;
        bst5.PreOrderRecursive(bst5, (int key, int _) => arrRes5_3.Add(key));
        UtilForCheckPtr(bst5);
        bst5 = bst5.RemoveNode(bst5, 18)!;
        bst5.PreOrderRecursive(bst5, (int key, int _) => arrRes5_4.Add(key));
        UtilForCheckPtr(bst5);
        bst5 = bst5.RemoveNode(bst5, 16)!;
        bst5.PreOrderRecursive(bst5, (int key, int _) => arrRes5_5.Add(key));
        UtilForCheckPtr(bst5);


        Assert.Equal(arrExcept1_1, arrRes1_1);
        Assert.Equal(arrExcept1_2, arrRes1_2);
        Assert.Equal(arrExcept1_3, arrRes1_3);
        Assert.Equal(arrExcept1_4, arrRes1_4);
        Assert.Equal(arrExcept1_5, arrRes1_5);

        Assert.Equal(arrExcept2_1, arrRes2_1);
        Assert.Equal(arrExcept2_2, arrRes2_2);
        Assert.Equal(arrExcept2_3, arrRes2_3);
        Assert.Equal(arrExcept2_4, arrRes2_4);
        Assert.Equal(arrExcept2_5, arrRes2_5);

        Assert.Equal(arrExcept3_1, arrRes3_1);
        Assert.Equal(arrExcept3_2, arrRes3_2);
        Assert.Equal(arrExcept3_3, arrRes3_3);
        Assert.Equal(arrExcept3_4, arrRes3_4);
        Assert.Equal(arrExcept3_5, arrRes3_5);

        Assert.Equal(arrExcept4_1, arrRes4_1);
        Assert.Equal(arrExcept4_2, arrRes4_2);
        Assert.Equal(arrExcept4_3, arrRes4_3);
        Assert.Equal(arrExcept4_4, arrRes4_4);
        Assert.Equal(arrExcept4_5, arrRes4_5);
        Assert.Equal(arrExcept4_6, arrRes4_6);
        Assert.Equal(arrExcept4_7, arrRes4_7);

        Assert.Equal(arrExcept5_1, arrRes5_1);
        Assert.Equal(arrExcept5_2, arrRes5_2);
        Assert.Equal(arrExcept5_3, arrRes5_3);
        Assert.Equal(arrExcept5_4, arrRes5_4);
        Assert.Equal(arrExcept5_5, arrRes5_5);

    }


    [Fact]
    // Тестирование поиска node
    public void TestFind(){
        BST<int, int> bst1 = CreateIntBST(arr1, 1111);
        BST<int, int> bst2 = CreateIntBST(arr2, 1111);
        BST<int, int> bst3 = CreateIntBST(arr3, 1111);
        BST<int, int> bst4 = CreateIntBST(arr4, 1111);
        BST<int, int> bst5 = CreateIntBST(arr5, 1111);

        int keyExcept1_1 = 5;
        int keyExcept1_2 = 16;
        int keyExcept1_3 = 0;

        int keyActually1_1 = bst1.FindNode(bst1, 5)!.Key;
        int keyActually1_2 = bst1.FindNode(bst1, 16)!.Key;
        int keyActually1_3 = bst1.FindNode(bst1, 111) is not null ? bst1.FindNode(bst1, 111)!.Key : 0;

        int keyExcept2_1 = 17;
        int keyExcept2_2 = 12;
        int keyExcept2_3 = 0;

        int keyActually2_1 = bst2.FindNode(bst2, 17)!.Key;
        int keyActually2_2 = bst2.FindNode(bst2, 12)!.Key;
        int keyActually2_3 = bst2.FindNode(bst2, 111) is not null ? bst2.FindNode(bst2, 111)!.Key : 0;

        int keyExcept3_1 = 4;
        int keyExcept3_2 = 6;
        int keyExcept3_3 = 0;

        int keyActually3_1 = bst3.FindNode(bst3, 4)!.Key;
        int keyActually3_2 = bst3.FindNode(bst3, 6)!.Key;
        int keyActually3_3 = bst3.FindNode(bst3, 111) is not null ? bst3.FindNode(bst3, 111)!.Key : 0;

        int keyExcept4_1 = 20;
        int keyExcept4_2 = 41;
        int keyExcept4_3 = 0;

        int keyActually4_1 = bst4.FindNode(bst4, 20)!.Key;
        int keyActually4_2 = bst4.FindNode(bst4, 41)!.Key;
        int keyActually4_3 = bst4.FindNode(bst4, 111) is not null ? bst4.FindNode(bst4, 111)!.Key : 0;

        int keyExcept5_1 = 13;
        int keyExcept5_2 = 16;
        int keyExcept5_3 = 0;

        int keyActually5_1 = bst5.FindNode(bst5, 13)!.Key;
        int keyActually5_2 = bst5.FindNode(bst5, 16)!.Key;
        int keyActually5_3 = bst5.FindNode(bst5, 111) is not null ? bst5.FindNode(bst5, 111)!.Key : 0;

        Assert.Equal(keyExcept1_1, keyActually1_1);
        Assert.Equal(keyExcept1_2, keyActually1_2);
        Assert.Equal(keyExcept1_3, keyActually1_3);

        Assert.Equal(keyExcept2_1, keyActually2_1);
        Assert.Equal(keyExcept2_2, keyActually2_2);
        Assert.Equal(keyExcept2_3, keyActually2_3);

        Assert.Equal(keyExcept3_1, keyActually3_1);
        Assert.Equal(keyExcept3_2, keyActually3_2);
        Assert.Equal(keyExcept3_3, keyActually3_3);

        Assert.Equal(keyExcept4_1, keyActually4_1);
        Assert.Equal(keyExcept4_2, keyActually4_2);
        Assert.Equal(keyExcept4_3, keyActually4_3);

        Assert.Equal(keyExcept5_1, keyActually5_1);
        Assert.Equal(keyExcept5_2, keyActually5_2);
        Assert.Equal(keyExcept5_3, keyActually5_3);

    }
}
