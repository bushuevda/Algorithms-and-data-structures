public class Program{
    static void Main(string[] args){
        
        Structure.Stack<int> s = new Structure.Stack<int>();
        s.push(1);
        s.push(2);
        s.push(3);
        
        Structure.Stack<int> s2 = new Structure.Stack<int>();
        s2.push(4);
        s2.push(5);
        s2.push(6);


        s.show();
        s2.show();
        s.swap(s2);
        s.show();
        s2.show();
    }
}