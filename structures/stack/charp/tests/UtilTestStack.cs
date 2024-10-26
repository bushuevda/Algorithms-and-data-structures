namespace Testing;
using Structure;

public class UtilTestStack <T>{

    /*
    Helped function for testing the push stack function 
    */
    public LinkedList<T> testPush(LinkedList<T> array){
        Stack<T> stack = new Stack<T>();
        foreach(T element in array){
            stack.push(element);
        }
        return stack.getArray();
    }

    /*
    Helped function for testing the pop stack function 
    */
    public LinkedList<T> testPop(LinkedList<T> array, int size_pop){
        Stack<T> stack = new Stack<T>();
        foreach(T element in array){
            stack.push(element);
        }
        for(int i = 0; i < size_pop; i++)
            stack.pop();
        return stack.getArray();
    }

    /*
    Helped function for testing the top stack function 
    */
    public T? testTop(LinkedList<T> array){
        Stack<T> stack = new Stack<T>();
        foreach(T element in array){
            stack.push(element);
        }
        return stack.top();
    }


    /*
    Helped function for testing the size stack function 
    */
    public int testSize(LinkedList<T> array){
        Stack<T> stack = new Stack<T>();
        foreach(T element in array){
            stack.push(element);
        }
        return stack.size();
    }

    /*
    Helped function for testing the empty stack function 
    */
    public bool testEmpty(LinkedList<T> array){
        Stack<T> stack = new Stack<T>();
        foreach(T element in array){
            stack.push(element);
        }
        return stack.empty();
    }

    /*
    Helped function for testing the swap stack function 
    */
    public LinkedList<T> testSwap(LinkedList<T> array1, LinkedList<T> array2){
        Stack<T> stack1 = new Stack<T>();
        foreach(T element in array1){
            stack1.push(element);
        }

        Stack<T> stack2 = new Stack<T>();
        foreach(T element in array2){
            stack2.push(element);
        }
        stack1.swap(stack2);

        return stack1.getArray();
    }
}