#include <hunspell.hxx>

void* new_hunspell(char*, char*);

void free_list(void*, char**, int);

int check_suggestions(void*, char*, int*, char***);

