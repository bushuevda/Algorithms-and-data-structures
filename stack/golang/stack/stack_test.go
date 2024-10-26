package stack

import "testing"

/*
Testing the swap stack function
*/
func TestSwap(t *testing.T) {
	utilInt := testStackUtil[int]{}
	result := utilInt.testSwap(TEST_DATA1, TEST_DATA3)
	if !utilInt.assertEqualsArrays(result, TEST_DATA3) {
		t.Error("\nExcepted: ", TEST_DATA3, " \nActually: ", result)
	}

	result = utilInt.testSwap(TEST_DATA3, TEST_DATA1)
	if !utilInt.assertEqualsArrays(result, TEST_DATA1) {
		t.Error("\nExcepted: ", TEST_DATA1, " \nActually: ", result)
	}
}

/*
Testing the empty stack function
*/
func TestEmpty(t *testing.T) {
	utilInt := testStackUtil[int]{}
	utilString := testStackUtil[string]{}
	utilBool := testStackUtil[bool]{}
	result := utilInt.testEmpty(TEST_DATA1)
	if utilBool.assertEqual(result, false) {
		t.Error("\nExcepted: ", false, " \nActually: ", result)
	}

	result2 := utilString.testEmpty(TEST_DATA2)
	if utilBool.assertEqual(result2, false) {
		t.Error("\nExcepted: ", false, " \nActually: ", result2)
	}

	result = utilInt.testEmpty(TEST_DATA3)
	if utilBool.assertEqual(result, false) {
		t.Error("\nExcepted: ", false, " \nActually: ", result)
	}

	result = utilInt.testEmpty([]int{})
	if utilBool.assertEqual(result, true) {
		t.Error("\nExcepted: ", true, " \nActually: ", result)
	}
}

/*
Testing the size stack function
*/
func TestSize(t *testing.T) {
	utilInt := testStackUtil[int]{}
	utilString := testStackUtil[string]{}

	result := utilInt.testSize(TEST_DATA1)
	if !utilInt.assertEqual(result, len(TEST_DATA1)) {
		t.Error("\nExcepted: ", len(TEST_DATA1), " \nActually: ", result)
	}

	result2 := utilString.testSize(TEST_DATA2)
	if !utilInt.assertEqual(result2, len(TEST_DATA2)) {
		t.Error("\nExcepted: ", len(TEST_DATA2), " \nActually: ", result2)
	}

	result = utilInt.testSize(TEST_DATA3)
	if !utilInt.assertEqual(result, len(TEST_DATA3)) {
		t.Error("\nExcepted: ", len(TEST_DATA3), " \nActually: ", result)
	}
}

/*
Testing the push stack function
*/
func TestPush(t *testing.T) {
	utilInt := testStackUtil[int]{}
	utilString := testStackUtil[string]{}

	result := utilInt.testPush(TEST_DATA1)
	if !utilInt.assertEqualsArrays(result, TEST_DATA1) {
		t.Error("\nExcepted: ", TEST_DATA1, " \nActually: ", result)
	}

	result2 := utilString.testPush(TEST_DATA2)
	if !utilString.assertEqualsArrays(result2, TEST_DATA2) {
		t.Error("\nExcepted: ", TEST_DATA2, " \nActually: ", result2)
	}

	result = utilInt.testPush(TEST_DATA3)
	if !utilInt.assertEqualsArrays(result, TEST_DATA3) {
		t.Error("\nExcepted: ", TEST_DATA3, " \nActually: ", result)
	}

}

/*
Testing the pop stack function
*/
func TestPop(t *testing.T) {
	utilInt := testStackUtil[int]{}
	utilString := testStackUtil[string]{}

	result := utilInt.testPop(TEST_DATA1, 3)
	if !utilInt.assertEqualsArrays(result, TEST_DATA_POP1) {
		t.Error("\nExcepted: ", TEST_DATA_POP1, " \nActually: ", result)
	}

	result2 := utilString.testPop(TEST_DATA2, 2)
	if !utilString.assertEqualsArrays(result2, TEST_DATA_POP2) {
		t.Error("\nExcepted: ", TEST_DATA_POP2, " \nActually: ", result)
	}

	result = utilInt.testPop(TEST_DATA3, 1)
	if !utilInt.assertEqualsArrays(result, TEST_DATA_POP3) {
		t.Error("\nExcepted: ", TEST_DATA_POP3, " \nActually: ", result)
	}

}

/*
Testing the top stack function
*/
func TestTop(t *testing.T) {
	utilInt := testStackUtil[int]{}
	utilString := testStackUtil[string]{}

	result := utilInt.testTop(TEST_DATA1)
	if !utilInt.assertEqual(result, TEST_DATA1[len(TEST_DATA1)-1]) {
		t.Error("\nExcepted: ", TEST_DATA1[len(TEST_DATA1)-1], " \nActually: ", result)
	}

	result2 := utilString.testTop(TEST_DATA2)
	if !utilString.assertEqual(result2, TEST_DATA2[len(TEST_DATA2)-1]) {
		t.Error("\nExcepted: ", TEST_DATA2[len(TEST_DATA2)-1], " \nActually: ", result2)
	}
}
