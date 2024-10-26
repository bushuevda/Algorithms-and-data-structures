package graphics_algorithms;
import java.util.ArrayList;

public class DDA_line {
    public ArrayList<ArrayList<Float>> calculate(float x1, float y1, float x2, float y2){
        ArrayList<Float> x = new ArrayList<Float>();
        ArrayList<Float> y = new ArrayList<Float>();
        ArrayList<ArrayList<Float>> result = new ArrayList<ArrayList<Float>>();

        int i = 0;

        float x_start = Math.round(x1);
        float y_start = Math.round(y1);
        float x_end = Math.round(x2);
        float y_end = Math.round(y2);

        float length = Math.max(Math.abs(x_end - x_start), Math.abs(y_end - y_start));

        float dX = (x2 - x1) / length;
        float dY = (y2 - y1) / length;
        x.add(x1);
        y.add(y1);
        i++;
        
        while(i < length){
            x.add(x.get(i - 1) + dX);
            y.add(y.get(i - 1) + dY);
            i++;
        }
        x.add(x2);
        y.add(y2);
        result.add(x);
        result.add(y);
        return result;
    }
}
