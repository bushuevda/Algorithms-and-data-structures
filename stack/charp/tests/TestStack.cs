using Structure;

namespace Testing;

public class TestStack
{
    /*
    Data for testing all stack functions
    */
    public LinkedList<int> TEST_DATA1 = new LinkedList<int>(new[]{1, 2, 3, 4, 5, 6, 7});
    public LinkedList<string> TEST_DATA2 = new LinkedList<string>(new[]{"a", "e", "b", "d", "f"});
    public LinkedList<int> TEST_DATA3 = new LinkedList<int>(new[]{7, 5, 1, 4, 15, 6, 7});

    /*
    Data for testing stack function pop
    */
    public LinkedList<int> TEST_DATA_POP1 = new LinkedList<int>(new[]{1, 2, 3, 4});             // -3
    public LinkedList<string> TEST_DATA_POP2 = new LinkedList<string>(new[]{"a", "e", "b"});    // -2
    public LinkedList<int> TEST_DATA_POP3 = new LinkedList<int>(new[]{7, 5, 1, 4, 15, 6});      // -1


    public UtilTestStack<int> utilInt = new UtilTestStack<int>();
    public UtilTestStack<string> utilString = new UtilTestStack<string>();
    public UtilTestStack<bool> utilBool = new UtilTestStack<bool>();

    [Fact]
    /*
    Testing the swap stack function
    */
    public void TestSwap(){
        Assert.Equal<LinkedList<int>>(utilInt.testSwap(TEST_DATA1, TEST_DATA3), TEST_DATA3);
        Assert.Equal<LinkedList<int>>(utilInt.testSwap(TEST_DATA3, TEST_DATA1), TEST_DATA1);
    }

    [Fact]
    /*
    Testing the empty stack function
    */
    public void TestEmpty(){
        Assert.False(utilInt.testEmpty(TEST_DATA1));
        Assert.False(utilString.testEmpty(TEST_DATA2));
        Assert.False(utilInt.testEmpty(TEST_DATA3));
        Assert.True(utilInt.testEmpty(new LinkedList<int>{}));
    }

    [Fact]
    /*
    Testing the size stack function
    */
    public void TestSize(){
        Assert.Equal(utilInt.testSize(TEST_DATA1), TEST_DATA1.Count);
        Assert.Equal(utilString.testSize(TEST_DATA2), TEST_DATA2.Count);
        Assert.Equal(utilInt.testSize(TEST_DATA3), TEST_DATA3.Count);
    }

    [Fact]
    /*
    Testing the push stack function
    */
    public void TestPush(){
        Assert.Equal(utilInt.testPush(TEST_DATA1), TEST_DATA1);
        Assert.Equal(utilString.testPush(TEST_DATA2), TEST_DATA2);
        Assert.Equal(utilInt.testPush(TEST_DATA3), TEST_DATA3);
    }

    [Fact]
    /*
    Testing the pop stack function
    */
    public void TestPop(){
        Assert.Equal(utilInt.testPop(TEST_DATA1, 3), TEST_DATA_POP1);           //-3
        Assert.Equal(utilString.testPop(TEST_DATA2, 2), TEST_DATA_POP2);        //-2
        Assert.Equal(utilInt.testPop(TEST_DATA3, 1), TEST_DATA_POP3);           //-1
    }



    [Fact]
    /*
    Testing the top stack function
    */
    public void TestTop(){
        Assert.Equal(utilInt.testTop(TEST_DATA1), TEST_DATA1.Last());
        Assert.Equal(utilString.testTop(TEST_DATA2), TEST_DATA2.Last());
        Assert.Equal(utilInt.testTop(TEST_DATA3), TEST_DATA3.Last());
    }

}