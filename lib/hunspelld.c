#include "hunspelld.h"

void* new_hunspell(char* aff, char* dic)
{
  void* h =  (void*)(new Hunspell(aff, dic, NULL, true));
  return h;
}

void free_list(void* h, char** sugg, int num)
{
  ((Hunspell*)h)->free_list(&sugg, num);
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
