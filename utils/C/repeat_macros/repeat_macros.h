#ifndef PREPROCCER_REPEAT_MACROS
#define PREPROCCER_REPEAT_MACROS 1

#define REPEAT_MACROS(count, macros, data) do{          \
                            switch (count)              \
                            {                           \
                            case 0:                     \
                                REPEAT0(macros, data)   \
                                break;                  \
                            case 1:                     \
                                REPEAT1(macros, data)   \
                                break;                  \
                            case 2:                     \
                                REPEAT2(macros, data)   \
                                break;                  \
                            case 3:                     \
                                REPEAT3(macros, data)   \
                                break;                  \
                            case 4:                     \
                                REPEAT4(macros, data)   \
                                break;                  \
                            case 5:                     \
                                REPEAT5(macros, data)   \
                                break;                  \
                            case 6:                     \
                                REPEAT6(macros, data)   \
                                break;                  \
                            case 7:                     \
                                REPEAT7(macros, data)   \
                                break;                  \
                            case 8:                     \
                                REPEAT8(macros, data)   \
                                break;                  \
                            case 9:                     \
                                REPEAT9(macros, data)   \
                                break;                  \
                            case 10:                    \
                                REPEAT10(macros, data)  \
                                break;                  \
                            case 11:                    \
                                REPEAT11(macros, data)  \
                                break;                  \
                            case 12:                    \
                                REPEAT12(macros, data)  \
                                break;                  \
                            case 13:                    \
                                REPEAT13(macros, data)  \
                                break;                  \
                            case 14:                    \
                                REPEAT14(macros, data)  \
                                break;                  \
                            case 15:                    \
                                REPEAT15(macros, data)  \
                                break;                  \
                            case 16:                    \
                                REPEAT16(macros, data)  \
                                break;                  \
                            case 17:                    \
                                REPEAT17(macros, data)  \
                                break;                  \
                            case 18:                    \
                                REPEAT18(macros, data)  \
                                break;                  \
                            case 19:                    \
                                REPEAT19(macros, data)  \
                                break;                  \
                            case 20:                    \
                                REPEAT20(macros, data)  \
                                break;                  \
                            case 21:                    \
                                REPEAT21(macros, data)  \
                                break;                  \
                            case 22:                    \
                                REPEAT22(macros, data)  \
                                break;                  \
                            case 23:                    \
                                REPEAT23(macros, data)  \
                                break;                  \
                            case 24:                    \
                                REPEAT24(macros, data)  \
                                break;                  \
                            case 25:                    \
                                REPEAT25(macros, data)  \
                                break;                  \
                            case 26:                    \
                                REPEAT26(macros, data)  \
                                break;                  \
                            case 27:                    \
                                REPEAT27(macros, data)  \
                                break;                  \
                            case 28:                    \
                                REPEAT28(macros, data)  \
                                break;                  \
                            case 29:                    \
                                REPEAT29(macros, data)  \
                                break;                  \
                            case 30:                    \
                                REPEAT30(macros, data)  \
                                break;                  \
                            case 31:                    \
                                REPEAT31(macros, data)  \
                                break;                  \
                            case 32:                    \
                                REPEAT32(macros, data)  \
                                break;                  \
                            case 33:                    \
                                REPEAT33(macros, data)  \
                                break;                  \
                            case 34:                    \
                                REPEAT34(macros, data)  \
                                break;                  \
                            case 35:                    \
                                REPEAT35(macros, data)  \
                                break;                  \
                            case 36:                    \
                                REPEAT36(macros, data)  \
                                break;                  \
                            case 37:                    \
                                REPEAT37(macros, data)  \
                                break;                  \
                            case 38:                    \
                                REPEAT38(macros, data)  \
                                break;                  \
                            case 39:                    \
                                REPEAT39(macros, data)  \
                                break;                  \
                            case 40:                    \
                                REPEAT40(macros, data)  \
                                break;                  \
                            case 41:                    \
                                REPEAT41(macros, data)  \
                                break;                  \
                            case 42:                    \
                                REPEAT42(macros, data)  \
                                break;                  \
                            case 43:                    \
                                REPEAT43(macros, data)  \
                                break;                  \
                            case 44:                    \
                                REPEAT44(macros, data)  \
                                break;                  \
                            case 45:                    \
                                REPEAT45(macros, data)  \
                                break;                  \
                            case 46:                    \
                                REPEAT46(macros, data)  \
                                break;                  \
                            case 47:                    \
                                REPEAT47(macros, data)  \
                                break;                  \
                            case 48:                    \
                                REPEAT48(macros, data)  \
                                break;                  \
                            case 49:                    \
                                REPEAT49(macros, data)  \
                                break;                  \
                            case 50:                    \
                                REPEAT50(macros, data)  \
                                break;                  \
                            }                           \
                        } while(0)                      \








#define REPEAT0(macros, data) macros(data[0]);                           \

#define REPEAT1(macros, data) macros(data[1]); REPEAT0(macros, data)       \

#define REPEAT2(macros, data) macros(data[2]); REPEAT1(macros, data)       \

#define REPEAT3(macros, data) macros(data[3]); REPEAT2(macros, data)       \

#define REPEAT4(macros, data) macros(data[4]); REPEAT3(macros, data)       \

#define REPEAT5(macros, data) macros(data[5]); REPEAT4(macros, data)       \

#define REPEAT6(macros, data) macros(data[6]); REPEAT5(macros, data)      \

#define REPEAT7(macros, data) macros(data[7]); REPEAT6(macros, data)       \

#define REPEAT8(macros, data) macros(data[8]); REPEAT7(macros, data)       \

#define REPEAT9(macros, data) macros(data[9]); REPEAT8(macros, data)       \

#define REPEAT10(macros, data) macros(data[10]); REPEAT9(macros, data)       \

#define REPEAT11(macros, data) macros(data[11]); REPEAT10(macros, data)       \

#define REPEAT12(macros, data) macros(data[12]); REPEAT11(macros, data)       \

#define REPEAT13(macros, data) macros(data[13]); REPEAT12(macros, data)       \

#define REPEAT14(macros, data) macros(data[14]); REPEAT13(macros, data)       \

#define REPEAT15(macros, data) macros(data[15]); REPEAT14(macros, data)       \

#define REPEAT16(macros, data) macros(data[16]); REPEAT15(macros, data)       \

#define REPEAT17(macros, data) macros(data[17]); REPEAT16(macros, data)       \

#define REPEAT18(macros, data) macros(data[18]); REPEAT17(macros, data)       \

#define REPEAT19(macros, data) macros(data[19]); REPEAT18(macros, data)       \

#define REPEAT20(macros, data) macros(data[20]); REPEAT19(macros, data)       \

#define REPEAT21(macros, data) macros(data[21]); REPEAT20(macros, data)       \

#define REPEAT22(macros, data) macros(data[22]); REPEAT21(macros, data)       \

#define REPEAT23(macros, data) macros(data[23]); REPEAT22(macros, data)       \

#define REPEAT24(macros, data) macros(data[24]); REPEAT23(macros, data)       \

#define REPEAT25(macros, data) macros(data[25]); REPEAT24(macros, data)       \

#define REPEAT26(macros, data) macros(data[26]); REPEAT25(macros, data)       \

#define REPEAT27(macros, data) macros(data[27]); REPEAT26(macros, data)       \

#define REPEAT28(macros, data) macros(data[28]); REPEAT27(macros, data)       \

#define REPEAT29(macros, data) macros(data[29]); REPEAT28(macros, data)       \

#define REPEAT30(macros, data) macros(data[30]); REPEAT29(macros, data)       \

#define REPEAT31(macros, data) macros(data[31]); REPEAT30(macros, data)       \

#define REPEAT32(macros, data) macros(data[32]); REPEAT31(macros, data)       \

#define REPEAT33(macros, data) macros(data[33]); REPEAT32(macros, data)       \

#define REPEAT34(macros, data) macros(data[34]); REPEAT33(macros, data)       \

#define REPEAT35(macros, data) macros(data[35]); REPEAT34(macros, data)       \

#define REPEAT36(macros, data) macros(data[36]); REPEAT35(macros, data)       \

#define REPEAT37(macros, data) macros(data[37]); REPEAT36(macros, data)       \

#define REPEAT38(macros, data) macros(data[38]); REPEAT37(macros, data)       \

#define REPEAT39(macros, data) macros(data[39]); REPEAT38(macros, data)       \

#define REPEAT40(macros, data) macros(data[40]); REPEAT39(macros, data)       \

#define REPEAT41(macros, data) macros(data[41]); REPEAT40(macros, data)       \

#define REPEAT42(macros, data) macros(data[42]); REPEAT41(macros, data)       \

#define REPEAT43(macros, data) macros(data[43]); REPEAT42(macros, data)       \

#define REPEAT44(macros, data) macros(data[44]); REPEAT43(macros, data)       \

#define REPEAT45(macros, data) macros(data[45]); REPEAT44(macros, data)       \

#define REPEAT46(macros, data) macros(data[46]); REPEAT45(macros, data)       \

#define REPEAT47(macros, data) macros(data[47]); REPEAT46(macros, data)       \

#define REPEAT48(macros, data) macros(data[48]); REPEAT47(macros, data)       \

#define REPEAT49(macros, data) macros(data[49]); REPEAT48(macros, data)       \

#define REPEAT50(macros, data) macros(data[50]); REPEAT49(macros, data)       \

#endif