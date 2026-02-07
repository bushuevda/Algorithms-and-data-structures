package com.bushuevda.stack.stack;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertTrue;

import java.util.ArrayList;

import org.junit.jupiter.api.Test;


public class StackTest {
    /*
    Data for testing all stack functions
    */
    final ArrayList<Integer> TEST_DATA1 = new ArrayList<Integer>(){{
        add(1); add(2); add(3); add(4); add(5); add(6); add(7);
    }};
    final ArrayList<String> TEST_DATA2 = new ArrayList<String>(){{
        add("a"); add("e"); add("b"); add("d"); add("f");}
    };    
    final ArrayList<Integer> TEST_DATA3 = new ArrayList<Integer>(){{
        add(7); add(5); add(1); add(4); add(15); add(6); add(7);
    }};        
     
    /*
    Data for testing stack function pop
    */
    final ArrayList<Integer> TEST_DATA_POP1 = new ArrayList<Integer>(){{  // -3
        add(1); add(2); add(3); add(4);
    }};

    final ArrayList<String> TEST_DATA_POP2 = new ArrayList<String>() {{ // -2
        add("a"); add("e"); add("b");
    }};

    final ArrayList<Integer> TEST_DATA_POP3 = new ArrayList<Integer>(){{ // -1
        add(7); add(5); add(1); add(4); add(15); add(6);
    }};
 
    
    /*
    Testing the swap stack function
    */
    @Test
    public void testSwap() {
        StackTestUtil<Integer> utilTest = new StackTestUtil<Integer>();

        ArrayList<Integer> result = utilTest.testSwap(TEST_DATA1, TEST_DATA3);
        assertArrayEquals(TEST_DATA3.toArray(), result.toArray());

        result = utilTest.testSwap(TEST_DATA3, TEST_DATA1);
        assertArrayEquals(TEST_DATA1.toArray(), result.toArray());
    }


    /*
    Testing the empty stack function
    */
    @Test
    public void testEmpty(){
        StackTestUtil<Integer> utilTestInteger = new StackTestUtil<Integer>();
        StackTestUtil<String> utilTestString = new StackTestUtil<String>();
        assertFalse(utilTestInteger.testEmpty(TEST_DATA1));
        assertFalse(utilTestString.testEmpty(TEST_DATA2));
        assertFalse(utilTestInteger.testEmpty(TEST_DATA3));
        assertTrue(utilTestInteger.testEmpty(new ArrayList<Integer>()));
    }

    /*
    Testing the size stack function
    */
    @Test
    public void testSize(){
        StackTestUtil<Integer> utilTestInteger = new StackTestUtil<Integer>();
        StackTestUtil<String> utilTestString = new StackTestUtil<String>();
        assertEquals(utilTestInteger.testSize(TEST_DATA1), TEST_DATA1.size());
        assertEquals(utilTestString.testSize(TEST_DATA2), TEST_DATA2.size());
        assertEquals(utilTestInteger.testSize(TEST_DATA3), TEST_DATA3.size());
    }

    /*
    Testing the push stack function
    */
    @Test 
    public void testPush(){
        StackTestUtil<Integer> utilTestInteger = new StackTestUtil<Integer>();
        StackTestUtil<String> utilTestString = new StackTestUtil<String>();
        assertArrayEquals(utilTestInteger.testPush(TEST_DATA1).toArray(), TEST_DATA1.toArray());
        assertArrayEquals(utilTestString.testPush(TEST_DATA2).toArray(), TEST_DATA2.toArray());
        assertArrayEquals(utilTestInteger.testPush(TEST_DATA3).toArray(), TEST_DATA3.toArray());
    }

    /*
    Testing the pop stack function
    */
    @Test
    public void testPop(){
        StackTestUtil<Integer> utilTestInteger = new StackTestUtil<Integer>();
        StackTestUtil<String> utilTestString = new StackTestUtil<String>();
        assertArrayEquals(utilTestInteger.testPop(TEST_DATA1, 3).toArray(), TEST_DATA_POP1.toArray());
        assertArrayEquals(utilTestString.testPop(TEST_DATA2, 2).toArray(), TEST_DATA_POP2.toArray());
        assertArrayEquals(utilTestInteger.testPop(TEST_DATA3, 1).toArray(), TEST_DATA_POP3.toArray());
    }

    /*
    Testing the top stack function
    */
    @Test
    public void testTop(){
        StackTestUtil<Integer> utilTestInteger = new StackTestUtil<Integer>();
        StackTestUtil<String> utilTestString = new StackTestUtil<String>();
        assertEquals(utilTestInteger.testTop(TEST_DATA1), TEST_DATA1.get(TEST_DATA1.size() - 1));
        assertEquals(utilTestString.testTop(TEST_DATA2), TEST_DATA2.get(TEST_DATA2.size() - 1));
        assertEquals(utilTestInteger.testTop(TEST_DATA3), TEST_DATA3.get(TEST_DATA3.size() - 1));
    }

}
