#ifndef DDA_LINE_H
#define DDA_LINE_H 1

#include <stdio.h>
#include <stdlib.h>

struct list_float
{
    float* list;
    int size_list;
};


struct pair_list_float{
    struct list_float list1;
    struct list_float list2;
};


#ifdef __cplusplus
extern "C"{
#endif

//Конструктор list_float
extern struct list_float new_list_float(int size);
//Деструктор list_float
extern void free_list_float(struct list_float* list);
//Деструктро pair_list_float
extern void free_pair_list_float(struct pair_list_float plf);

//Функция dda_line
extern struct pair_list_float dda_line(float x1, float x2, float y1, float y2);

void show_result(struct pair_list_float plfl);

#ifdef __cplusplus
}
#endif

#endif