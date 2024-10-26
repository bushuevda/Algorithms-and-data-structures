import graphics_algorithms.DDA_line;
import java.util.ArrayList;
public class Main{

    public static void main(String[] args) {
        DDA_line dda = new DDA_line();
        ArrayList<ArrayList<Float>> result = dda.calculate(1,2,3,4);
        System.out.println(result);
    }
}