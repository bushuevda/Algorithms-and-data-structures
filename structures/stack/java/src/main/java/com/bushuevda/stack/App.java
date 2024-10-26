package com.bushuevda.stack;

import com.bushuevda.stack.stack.Stack;

public class App {
    public static void main(String[] args) {
        Stack <Integer> b = new Stack<Integer>();
        b.push(1);
        b.push(2);
        System.out.println(b.top());
        b.push(3);
        System.out.println(b.top());

        Stack <Integer> a = new Stack<Integer>();
        a.push(4);
        a.push(5);
        a.push(6);
        a.pop();
        a.show();

    }
}
