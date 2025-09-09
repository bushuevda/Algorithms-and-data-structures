
namespace Mathematics{

    public class QuadraticEquation
    {
        static public List<double> Calculate(double a, double b, double c){
            var (a_, b_, c_) = (a, b, c);
            double x1, x2;
            double D = b * b - 4 * a * c_;
            List<double> result = [];
            if(D > 0){
                //calculate two roots
                x1 = (-b_ - Math.Sqrt(D)) / (2 * a_);
                x2 = (-b_ + Math.Sqrt(D)) / (2 * a_);
                result.Add(x1);
                result.Add(x2);
            } else if (D == 0){
                //calculate one root
                x1 = -b_ / (2 * a_);
                result.Add(x1);
            }
            
            return result;
        }



        
    } 
}