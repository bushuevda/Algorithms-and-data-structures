#ifndef QUADRATIC_EQUATION_H
#define QUADRATIC_EQUATION_H 1

#ifdef __cplusplus
extern "C"{
#endif

typedef struct _QuadraticEquation{
    double a_;
    double b_;
    double c_;
    double x1_;
    double x2_;
    int root_count_;
} QuadraticEquation;

/** Получение x1 */
extern double get_x1(QuadraticEquation* qe);
/*Получение x2*/
extern double get_x2(QuadraticEquation* qe);
/** Получение кол-ва корней */
extern int get_root_count(QuadraticEquation* qe);
/** Задание переменных уравнения */
extern void set(QuadraticEquation* qe, double a, double b, double c);
/**Вычисление корней уравнения */
extern void calculate(QuadraticEquation* qe);
/**Вывод информации об уравнении */
extern void show(QuadraticEquation* qe);
/**Выделение памяти под новое уравнение */
extern QuadraticEquation* create_qe();
/**Очистка памяти */
extern int free_qe(QuadraticEquation * qe);


#ifdef __cplusplus
}
#endif

#endif