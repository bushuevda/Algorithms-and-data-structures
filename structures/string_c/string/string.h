#ifndef STRING_H
#define STRING_H 1
#include <stdlib.h>

#ifdef __cplusplus
extern "C"{
#endif

typedef struct _String{		
	int size;				
	int capacity;			
	char *data;			
}String;					

typedef struct _SplitString{
	int size;
	int capacity;			
	String** strs;
}SplitString;

///--------String funtions---------
///--------------------------------
//---------CREATE STRING-----------
extern String* string_new(char* data);
//---------CREATE STRING END-------


//---------FREE STRING-----------
extern int string_free(String* str);
//---------FREE STRING END---------


//-----COPY AND MOVING STRING------
extern String* string_copy(String* str);
extern void string_move(String* str1, String* str2);
//--COPY AND MOVING STRING END-----

//------ATTRIBUTES STRING ---------
extern char string_at(String* str, int index);
extern char* string_data(String* str);
extern String* string_sub(String* str, int begin, int end);

extern int empty(String* str);
extern size_t string_size(String* str);
extern int string_capacity(String* str);

//----MODIFIERS STRING----------
extern String* string_concat(String* str1, String* str2);
extern void string_push_back(String* str, char ch);
extern void string_resize(String* str, int new_size);
extern void string_insert(String* str, char* s, int pos); 
extern void string_replace(String* str, char ch_patt, char s_repl); 
extern void string_clear(String* str); 
//----MODIFIERS STRING END------


///--------SplitString funtions---------
extern SplitString* string_split(String* str, char symbol); 
extern String* string_join(SplitString* sp_str, char symbol); 
extern int split_string_free(SplitString* sp_str);
extern String* split_string_at(SplitString* sp_str, int index);
///--------------------------------

#ifdef __cplusplus
}
#endif

#endif
