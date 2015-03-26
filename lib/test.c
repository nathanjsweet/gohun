#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>
#include "hunspelld.h"

extern int errno;

int main(void)
{
  FILE * aFile;
  FILE * dFile;
  aFile = fopen ("lib/dictionaries/en_US.aff" , "r");
  if(aFile == NULL) {
    printf("failed to open affix dictionary\n%s\n", strerror(errno));
    return 1;
  }
  fseek(aFile, 0, SEEK_END);
  long size = ftell(aFile);
  char* aff = (char*)calloc(size + 1, sizeof(char));
  if(!aff){
    fclose(aFile);
    printf("failed to allocate affix buffer\n");
    return 1;
  }
  fseek(aFile, 0, SEEK_SET);
  int s = fread(aff, sizeof(char), size + 1,  aFile);
  if(s != size) {
    fclose(aFile);
    printf("failed to copy to affix buffer:%d,%d\n", s, size);
    free(aff);
    return 1;
  }
  fclose(aFile);
  
  dFile = fopen ("lib/dictionaries/en_US.dic" , "r");
  if(dFile == NULL) {
    printf("failed to open word dictionary\n");
    return 1;
  }
  fseek(dFile, 0, SEEK_END);
  size = ftell(aFile);
  char* dic = (char*)calloc(size + 1, sizeof(char));
  if(!dic){
    fclose(dFile);
    printf("failed to allocate dictionary buffer\n");
    return 1;
  }
  fseek(dFile, 0, SEEK_SET);
  s = fread(dic, sizeof(char), size + 1,  dFile);
  if(s != size) {
    fclose(dFile);
    printf("failed to copy to dictionary buffer\n%d,%d", s, size);
    return 1;
  }
  fclose(dFile);
  
  void* h = new_hunspell(aff, dic);
  int num = 0;
  char** sugg;
  char* w = "calor";
  int correct = check_suggestions(h, w, &num, &sugg);
  if(!correct) {
    printf("\"%s\" is spelled incorrectly\nHere are some suggestions:\n", w);
    int i;
    for(i = 0; i < num; i++)
      printf("%s\n", sugg[i]);
  }
  else {
    printf("\"%s\" is spelled correctly\n", w);
  }
  return 0;
}
