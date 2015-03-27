package main
/*
#include <stdlib.h>
#include "lib/hunspelld.h"
#cgo LDFLAGS:-Lbuild/ -lhunspell -lstdc++
*/
import "C"

import (
	"unsafe"
	"os"
	"io/ioutil"
	"log"
)

func main() {
	file, err := os.Open("./lib/dictionaries/en_US.aff")
	if err != nil {
		log.Fatal(err)
	}
	aff, e := ioutil.ReadAll(file)
	if e != nil {
		log.Fatal(e)
	}
	file.Close()
	file, err = os.Open("./lib/dictionaries/en_US.dic")
	if err != nil {
		log.Fatal(err)
	}
	dic, e2 := ioutil.ReadAll(file)
	file.Close()
	if e2 != nil {
		log.Fatal(e2)
	}

	g := NewGohun(aff, dic)
	b, l, r := g.CheckSuggestions("calor")
	log.Println("calor spelled correctly? ", b)
	for i := 0; i < l; i++ {
		log.Println(r[i])
	}
}

type Gohun struct{
	hunspell unsafe.Pointer
}

func NewGohun(aff, dic []byte) *Gohun {
	g := new(Gohun)
	g.hunspell = C.new_hunspell((*C.char)(unsafe.Pointer(&aff[0])),(*C.char)(unsafe.Pointer(&dic[0])))
	return g
}

func (g *Gohun) CheckSuggestions(word string) (bool, int, []string) {
	w := C.CString(word)
	defer C.free(unsafe.Pointer(w))
	n := C.int(0)
	var s **C.char
	b := int(C.check_suggestions(g.hunspell, w, &n, (***C.char)(unsafe.Pointer(&s)))) != 0
	l := 0
	var r []string
	if(!b) {
		l = int(n)
		sl := *(*[]*C.char)(unsafe.Pointer(&s))
		for i := 0; i < l; i++ {
			r = append(r, C.GoString(sl[i]))
		}
		defer C.free_list(g.hunspell, (***C.char)(unsafe.Pointer(&s)), n)
	}
	return b, l, r
}
