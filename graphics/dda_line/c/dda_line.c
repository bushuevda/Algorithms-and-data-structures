#include "dda_line.h"
#include "math.h"

//Конструктор list_float
struct list_float new_list_float(int size){
    struct list_float lf ;
    lf.list = calloc(sizeof(struct list_float), size);
    lf.size_list = size;
    return lf;
}

//Деструктор list_float
void free_list_float(struct list_float* list){
    free(list->list);
    list->list = NULL;
}

void free_pair_list_float(struct pair_list_float plf){
    free(plf.list1.list);
    plf.list1.list = NULL;
    free(plf.list2.list);
    plf.list2.list = NULL;
}


void show_result(struct pair_list_float plfl){
    printf("%d\n", plfl.list1.size_list);
    for(int i = 0, k = plfl.list1.size_list; i < k; i++){
        printf("x: %f, y: %f\n", plfl.list1.list[i], plfl.list2.list[i] );
    }

}

struct pair_list_float dda_line(float x1, float x2, float y1, float y2){
    int i = 0;

    struct list_float x;
    struct list_float y;

    float x_start = roundf(x1);
    float y_start = roundf(y1);
    float x_end = roundf(x2);
    float y_end = roundf(y2);

    int length = __max(abs(x_end - x_start), abs(y_end - y_start));

    float dX = (x2 - x1) / length;
    float dY = (y2 - y1) / length;
    printf("%d\n", length);
    x = new_list_float(length);
    y = new_list_float(length);

    x.list[0] = x1;
    y.list[0] = y1;
    i++;


    while(i < length){
        x.list[i] = x.list[i - 1] + dX;
        y.list[i] = y.list[i - 1] + dY;
        i++;
    }


    x.list[i] = x2;
    y.list[i] = y2;


    struct pair_list_float p_list = {x, y};

    return p_list;
}