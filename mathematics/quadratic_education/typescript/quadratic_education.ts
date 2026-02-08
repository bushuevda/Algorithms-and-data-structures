export type QuadraticEquation = {
    a: number;
    b: number;
    c: number;
    x1?: number;
    x2?: number; 
    root_count?: number;
};


export class QuadraticEquationUtil{
 
    static calculate(qe: QuadraticEquation){

        let D: number = qe.b * qe.b - 4 * qe.a * qe.c;

        if(D > 0){
            qe.x1 = (-qe.b - Math.sqrt(D)) / (2 * qe.a);
            qe.x2 = (-qe.b + Math.sqrt(D)) / (2 * qe.a);
            qe.root_count = 2;
        } else if (D == 0){
            qe.x1 = -qe.b / (2 * qe.a);
            qe.x2 = NaN;
            qe.root_count = 1;
        } else {
            qe.x1 = NaN;
            qe.x2 = NaN;
            qe.root_count = 0;
        }
        return qe;
    }

}