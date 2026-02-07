#include "string.h"
#include <stdio.h>
#include <assert.h>

///--------String funtions---------
///--------------------------------
//---------CREATE STRING-----------
/*
Get size string
*/
static size_t size__(char* data){
	assert(data != NULL && "data must not be null");

	size_t count = 0;
	for(; *(data + count) != '\0'; count++){}
	return count;
}


/*
Memory allocation for string
*/
static String* new__(char* data){
	assert(data != NULL && "str must not be null");

	String* str = (String*) malloc(sizeof(String));
	assert(str != NULL && "str must not be null");
	
	int count = size__(data);
	
	str->data = (char*) calloc(count, sizeof(char));
	assert(str->data != NULL && "str->data must not be null");
	for(int i = 0; i < count; i++)
		*(str->data + i) = *(data + i);
	
	str->size = str->capacity = count;
	
	return str;
}


/*
Create new string
*/
String* string_new(char* data){
	assert(data != NULL && "data must not be null");

	return new__(data);
}
//--------CREATE STRING END--------


//---------FREE STRING-----------
/*
Free memory by string
*/
static int free__(String* str){
	assert(str != NULL && "str must not be null");

	free(str->data);
	free(str);
	str = NULL;
	return 1;
}

/*
Delete string
*/
int string_free(String* str){
	free__(str);
}
//---------END FREE STRING--------


//--------COPY AND MOVING STRING--------------

String* string_copy(String* str){
	return new__(str->data);
}


/*
Moving str1* to str2*
*/
void string_move(String* str1, String* str2){
	assert(str1 != NULL && str2 != NULL && "str1 and str2 must not be null");

	if(str2->data != NULL)
		free(str2->data);
	str2->data = str1->data;
	str2->size = str1->size;
	str2->capacity = str1->capacity;
	str1->data = NULL;
	free(str1);
}

/*
Moving pointer char* to String* data
*/
static String* move_data__(char* data){
	assert(data != NULL && "data must not be null");

	String* str = (String*) malloc(sizeof(String));
	assert(str != NULL);
	
	str->data = data;
	str->size = str->capacity = size__(data);
	
	return str;
}

//--------COPY AND MOVING STRING END----------


//------ATTRIBUTES STRING ---------
/*
Get char string by index
*/
char string_at(String* str, int index){
	assert(str != NULL && "str must not be null");

	return (str->size < index) ? '\0' : *(str->data + index);
}


/*
Get pointer string data
*/
char* string_data(String* str){
	assert(str != NULL && "str must not be null");

	return str->data;
}


/*
Get substring 
*/
String* string_sub(String* str, int begin, int end){
	assert(str != NULL && "str must not be null");
	
	assert(str->size >= end && str->size > begin && "size >= end > begin");
	
	size_t size_sub = end - begin;

	char* sub = (char*) calloc(size_sub, sizeof(char));
	
	for(int i = begin, j = 0; i < end; i++, j++)
		*(sub + j) = *(str->data + i);
	
	return move_data__(sub);
}


/*
Check string in empty
*/
int empty(String* str){
	assert(str != NULL && "str must not be null");

	return (str->size == 0) ? 1 : 0;
}

/*
Get size string
*/
size_t string_size(String* str){
	assert(str != NULL && "str must not be null");

	return str->size;
}

/*
Get capacity string
*/
int string_capacity(String* str){
	assert(str != NULL && "str must not be null");
	
	return str->capacity;
}
//------ATTRIBUTES STRING END------



//----MODIFIERS STRING----------
/*
Concatenation of two strings
*/
String* string_concat(String* str1, String* str2){
	assert(str1 != NULL && str2 != NULL && "str1 and str2 must not be null");

	size_t size_str1 = size__(str1->data);
	size_t size_str2 = size__(str2->data);
	
	char* ndata = (char*) calloc(size_str1 + size_str2, sizeof(char));
	
	int i = 0, j = 0;
	while(i < size_str1)
		*(ndata + i) = *(str1->data + i++);
	
	while(j < size_str2)
		*(ndata + i++) = *(str2->data + j++);
	
	return move_data__(ndata);
}


/*
Push back char to string
*/
void string_push_back(String* str, char ch){
	assert(str != NULL && "str must not be null");
	if(str->size == str->capacity){
		if(str->capacity == 0)
			str->capacity++;
		str->capacity *= 2;
		str->data = (char*) realloc(str->data, str->capacity * sizeof(char));
	}
	*(str->data + str->size++) = ch;
}

/*
Resize string
*/
void string_resize(String* str, int new_size){
	assert(str != NULL && "str must not be null");

	char* new_data = (char*) calloc(new_size, sizeof(char));
	for(int i = 0; i < new_size; i++)
		*(new_data + i) = *(str->data + i);
	free(str->data);
	str->data = new_data;
	str->size = new_size;
	str->capacity = new_size;
}


/*
Insert string1 to string2 by position
*/
void string_insert(String* str, char* s, int pos){
	assert(str != NULL && s != NULL && "str and s must not be null");
	
	size_t size_s = size__(s);
	size_t size_str = str->size;
	size_t new_size = size_s + size_str;

	String* ns;
	if(pos > size_str - 1){
		ns = string_concat(str, string_new(s));
	} else {
		String* sub1 = string_sub(str, 0, pos);
		String* sub2 = string_new(s);
		String* sub3 = string_sub(str, pos, str->size);
		ns = string_concat(string_concat(sub1, sub2), sub3);
		string_free(sub1);
		string_free(sub2);
		string_free(sub3);
	}
	string_move(ns, str);
	ns = NULL;
}

/*
Replace symbol in string
*/
void string_replace(String* str, char ch_patt, char s_repl){
	assert(str != NULL && "str must not be null");

	for(int i = 0; *(str->data + i) != '\0'; i++){
		if(*(str->data + i) == ch_patt)
			*(str->data + i) = s_repl;
	}
}

/*
Clear string
*/
void string_clear(String* str){
	assert(str != NULL && "str must not be null");

	string_resize(str, 1);
	*(str->data + 0) = ' ';
	str->size = 0;
}

//----MODIFIERS STRING END---------
///-----STRING FUNCTIONS END-------
///--------------------------------



///--------SPLITSTRING FUNCTIONS---------
///--------------------------------------

/*
Create new SplitString
*/
static SplitString*  split_str_new__(){
	SplitString* sp_str = (SplitString*) malloc(sizeof(SplitString));
	assert(sp_str != NULL && "SplitString must be not null");
	
	sp_str->strs = (String**) malloc(sizeof(String*));

	sp_str->size = 0;
	sp_str->capacity = 1;
	return sp_str;
}


/*
Push back String to SplitString
*/
static void split_str_push_back__(SplitString* sp_str, String* str){
	assert(sp_str != NULL && str != NULL && "sp_str and str must not be null");

	if(sp_str->capacity == sp_str->size){
		if(sp_str->capacity == 0)
			sp_str->capacity++;
		sp_str->capacity *= 2;
		sp_str->strs = (String**) realloc(sp_str->strs, sp_str->capacity * sizeof(String*));
	}
	*(sp_str->strs + sp_str->size++) = str;
	
}


/*
Split string
*/
SplitString* string_split(String* str, char symbol){
	
	assert(str != NULL && "str must not be null");

	SplitString* sp_str = split_str_new__();
	String* nstr = string_new("");
	for(int i = 0; *(str->data + i) != '\0'; i++){
		if(symbol != *(str->data + i)){
			string_push_back(nstr, *(str->data + i));
		} else {
			if(nstr->size > 0)
				split_str_push_back__(sp_str, nstr);
			nstr = string_new("");
		}
	}
	if(nstr->size > 0)
		split_str_push_back__(sp_str, nstr);
	
	return sp_str;
}

/*
Join string
*/
String* string_join(SplitString* sp_str, char symbol){
	
	assert(sp_str != NULL && "sp_str must not be null");

	String* nstr = string_new("");
	for(int i = 0; i < sp_str->size; i++){
		if(i < sp_str->size - 1)
			string_push_back(*(sp_str->strs + i), symbol);	
		nstr = string_concat(nstr, *(sp_str->strs + i));
	}
	return nstr;
}

/*
Delete SplitString 
*/
int split_string_free(SplitString* sp_str){
	
	assert(sp_str != NULL && "sp_str must not be null");

	
	for(int i = 0; i < sp_str->size; i++){
		string_free(*(sp_str->strs + i));
	}
	free(sp_str);
	return 1;
}


/*
Get char string by index
*/
String* split_string_at(SplitString* sp_str, int index){
	
	assert(sp_str != NULL && "sp_str must not be null");

	return *(sp_str->strs + index);
}

///-------SPLITSTRING FUNCTIONS END------
///--------------------------------------
