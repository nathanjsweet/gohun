#include "hunspelld.h"

void* new_hunspell(char* aff, char* dic)
{
  void* h =  (void*)(new Hunspell(aff, dic, NULL, true));
  return h;
}

void delete_hunspell(void* h)
{
  delete h;
}

void free_list(void* h, char*** sugg, int num)
{
  ((Hunspell*)h)->free_list(sugg, num);
}

int check_suggestions(void* h, char* word, int* num, char*** sugg)
{
  Hunspell *hun = (Hunspell*)h;
  if(hun->spell(word)) {
    *num = 0;
    return 1;
  }
  else {
    *num = hun->suggest(sugg, word);
    return 0;
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

int stem(void* h, char* word, char*** sugg)
{
  Hunspell *hun = (Hunspell*)h;
  return hun->stem(sugg, word);
}

int generate(void* h, char* word1, char* word2, char*** sugg)
{
  Hunspell *hun = (Hunspell*)h;
  return hun->generate(sugg, word1, word2);
}

int analyze(void* h, char* word, char*** sugg)
{
  Hunspell *hun = (Hunspell*)h;
  return hun->analyze(sugg, word);
}
