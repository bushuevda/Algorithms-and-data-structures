#ifndef PROTO_THREAD_H
#define PROTO_THREAD_H 1


struct pt_c{ 
    unsigned short l;
};


#define DECLARE(pt) pt->l = 0

//Инициализация Pthread
#define BEGIN(pt) switch (pt->l) { \
                      case 0:
                      
#define YIELD(pt, val)  pt->l = __LINE__;   \
                    return val;         \
                    case __LINE__:      \


//Завершение выполнения Pthread
#define END }



#endif