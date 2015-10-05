#ifdef __cplusplus
extern "C" {
  #endif
void* new_hunspell(char*, char*);

void delete_hunspell(void*);

void free_list(void*, char***, int);

char** check_suggestions(void*, char*, int*, int*);

  int is_correct(void*, char*);

int add_dic(void*, char*);

int add_word(void*, char*);

int remove_word(void*, char*);

char** stem(void*, char*, int*);

char** generate(void*, char*, char*, int*);

char** analyze(void*, char*, int*);
  #ifdef __cplusplus
}
#endif
