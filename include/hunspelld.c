#include "hunspelld.h"
#include <hunspell.hxx>

void* new_hunspell(char* aff, char* dic)
{
  void* h =  (void*)(new Hunspell(aff, dic, NULL, true));
  return h;
}

void delete_hunspell(void* h)
{
  delete (Hunspell*)h;
}

void free_list(void* h, char*** sugg, int num)
{
  ((Hunspell*)h)->free_list(sugg, num);
}

char** check_suggestions(void* h, char* word, int* num, int* cor)
{
  Hunspell *hun = (Hunspell*)h;
  if(hun->spell(word)) {
    *num = 0;
    *cor = 1;
    return NULL;
  }
  else {
    char** sugg;
    *num = hun->suggest(&sugg, word);
    *cor = 0;
    return sugg;
  } 
}

int add_dic(void* h, char* dic)
{
  Hunspell *hun = (Hunspell*)h;
  return hun->add_dic(dic) == 0 ? 1 : 0;
}

int add_word(void* h, char* word)
{
  Hunspell *hun = (Hunspell*)h;
  return hun->add(word) == 0 ? 1 : 0;
}

int remove_word(void* h, char* word)
{
  Hunspell *hun = (Hunspell*)h;
  return hun->remove(word) == 0 ? 1 : 0;
}

char** stem(void* h, char* word, int* count)
{
  Hunspell *hun = (Hunspell*)h;
  char** sugg;
  *count = hun->stem(&sugg, word);
  return sugg;
}

char** generate(void* h, char* word1, char* word2, int* count)
{
  Hunspell *hun = (Hunspell*)h;
  char** sugg;
  *count = hun->generate(&sugg, word1, word2);
  return sugg;
}

char** analyze(void* h, char* word, int* count)
{
  Hunspell *hun = (Hunspell*)h;
  char** sugg
  *count = hun->analyze(&sugg, word);
  return count;
}
