#include "quadratic_education.h"
#include <stdio.h>
#include <stdlib.h>
#include <math.h>


/** Получение x1 */
double get_x1(QuadraticEquation* qe){
    return qe -> x1_;
}

/*Получение x2*/
double get_x2(QuadraticEquation* qe){
    return qe -> x2_;
}

/** Получение кол-ва корней */
int get_root_count(QuadraticEquation* qe){
    return qe -> root_count_;
}

/** Задание переменных уравнения */
void set(QuadraticEquation* qe, double a, double b, double c){
    qe -> a_ = a;
    qe -> b_ = b;
    qe -> c_ = c;
}

/**Вычисление корней уравнения */
void calculate(QuadraticEquation* qe){
    double D = qe -> b_ * qe -> b_ - 4 * qe -> a_ * qe -> c_;

    if(D > 0){
        //два корня
        qe -> x1_ = (-qe -> b_ - sqrt(D)) / (2 * qe -> a_);
        qe -> x2_ = (-qe -> b_ + sqrt(D)) / (2 * qe -> a_);
        qe -> root_count_ = 2;
    } else if (D == 0){
        //один корень
        qe -> x1_ = -qe -> b_ / (2 * qe -> a_);
        qe -> x2_ = NAN;
        qe -> root_count_ = 1;
    } else {
        qe -> x1_ = NAN;
        qe -> x2_ = NAN;
        qe -> root_count_ = 0;
    }
}

/**Выделение памяти под новое уравнение */
QuadraticEquation* create_qe(){
    QuadraticEquation* qe = malloc(sizeof(QuadraticEquation));
    return qe;
}

/**Очистка памяти */
int free_qe(QuadraticEquation * qe){
    free(qe);
    return 1;
}

/**Вывод информации об уравнении */
void show(QuadraticEquation* qe){
    printf("x1: %.2f, x2: %.2f, roots: %.2f\n", qe->x1_, qe->x2_, qe->root_count_);
}
