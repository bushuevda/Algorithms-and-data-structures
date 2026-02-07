#include "string/string.h"
#include <stdio.h>


int main(int argc, char** argv){
	String* s = string_new("Hello, sds, abs");
	SplitString* sp = string_split(s, ',');
	for(int i = 0; i < sp->size; i++){
		printf(split_string_at(sp, i)->data);
		printf("\n");
	}
	printf(string_join(sp, ',')->data);
	split_string_free(sp);
	return 0;
}
