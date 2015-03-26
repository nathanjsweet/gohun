#include <hunspell.hxx>

void* new_hunspell(char*, char*);

void delete_hunspell(void*);

void free_list(void*, char***, int);

int check_suggestions(void*, char*, int*, char***);

int add_dic(void*, char*);

int add_word(void*, char*);

int remove_word(void*, char*);

int stem(void*, char*, char***);

int generate(void*, char*, char*, char***);
