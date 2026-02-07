import {Stack} from "./stack/stack";



let s: Stack<number> = new Stack<number>;

s.push(1);
s.push(2);
s.push(3);
s.show();
console.log(s.top);
s.pop();
s.show();