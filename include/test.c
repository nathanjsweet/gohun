#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>
#include "hunspelld.h"

extern int errno;

void test_suggest(void*);
//void test_add_dic(void*);

int main(void)
{
  FILE * aFile;
  FILE * dFile;
  aFile = fopen ("include/dictionaries/en_US.aff" , "r");
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
  
  dFile = fopen ("include/dictionaries/en_US.dic" , "r");
  if(dFile == NULL) {
    printf("failed to open word dictionary\n%s\n", strerror(errno));
    return 1;
  }
  fseek(dFile, 0, SEEK_END);
  size = ftell(dFile);
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
  test_suggest(h);
  //  test_add_dic(h);
  return 0;
}

void test_suggest(void* h)
{
  int num = 0;
  char* w = "calor";
  int correct = 0;
  char** sugg = check_suggestions(h, w, &num, &correct);
  printf("test suggestions:\n");
  if(!correct) {
    printf("\t\"%s\" is spelled incorrectly\n\tHere are some suggestions:\n", w);
    int i;
    for(i = 0; i < num; i++)
      printf("\t\t%s\n", sugg[i]);
    free_list(h, &sugg, num);
  }
  else {
    printf("\t\"%s\" is spelled correctly\n", w);
  }
}
/*
void test_add_dic(void* h)
{
  int size = 0;
  printf("\nadd_dic test:\n");
  FILE *dFile = fopen ("include/dictionaries/en_CA.dic" , "r");
  if(dFile == NULL) {
    printf("\tfailed to open word dictionary: %s\n", strerror(errno));
    return;
  }
  fseek(dFile, 0, SEEK_END);
  size = ftell(dFile);
  char* dic = (char*)calloc(size + 1, sizeof(char));
  if(!dic){
    fclose(dFile);
    printf("\tfailed to allocate dictionary buffer\n");
    return;
  }
  fseek(dFile, 0, SEEK_SET);
  int s = fread(dic, sizeof(char), size + 1,  dFile);
  if(s != size) {
    fclose(dFile);
    printf("\tfailed to copy to dictionary buffer\n%d,%d\n", s, size);
    return;
  }
  fclose(dFile);
  int num = 0;
  char** sugg;
  char* w = "colour";
  int correct = check_suggestions(h, w, &num, &sugg);
  printf("\tbefore en_CA is added ");
  if(!correct) {
    printf("\"%s\" is spelled incorrectly\n\tHere are some suggestions:\n", w);
    int i;
    for(i = 0; i < num; i++)
      printf("\t\t%s\n", sugg[i]);
    free_list(h, &sugg, num);
  }
  else {
    printf("\"%s\" is spelled correctly\n", w);
  }

  correct = add_dic(h, dic);
  if(!correct) {
    printf("\t for someone reason hunspell failed to add the dictionary.");
    return;
  }

  correct = check_suggestions(h, w, &num, &sugg);
  printf("\tafter en_CA is added ");
  if(!correct) {
    printf("\"%s\" is spelled incorrectly\n\tHere are some suggestions:\n", w);
    int i;
    for(i = 0; i < num; i++)
      printf("\t\t%s\n", sugg[i]);
    free_list(h, &sugg, num);
  }
  else {
    printf("\"%s\" is spelled correctly\n", w);
  }
  
}
*/
