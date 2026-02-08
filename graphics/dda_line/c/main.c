#include <stdio.h>
#include "dda_line.h"

int main(){

    struct pair_list_float plf = dda_line(13, 0, -4, 0);
    show_result(plf);
    free_pair_list_float(plf);

    return 0;
}