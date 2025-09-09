#include <stdio.h>
#include "quadratic_education/quadratic_education.h"

int main(){

    QuadraticEquation* qe = create_qe();
    set(qe, 6, 2, 3);
    calculate(qe);
    show(qe);
    set(qe, -1, 6, 7);
    calculate(qe);
    show(qe);
    return 0;
}