#include "protothread/protothread.h"
#include <stdio.h>
#include <stddef.h>
#include <stdlib.h>



const int next(struct pt_c* p)
{
    BEGIN(p);
        YIELD(p, 1);
        YIELD(p, 2);
        YIELD(p, 3);
    END;
    return 0;
}

int main(){
    struct pt_c* p = malloc(sizeof(struct pt_c));
    DECLARE(p);
    int i = 0;
    while(1){
        const int val = next(p);
        if(i > 2)
            break;
        else
            printf("%d\n", val);
        printf("%s\n", "op");
        i++;
    }

    return 0;
}