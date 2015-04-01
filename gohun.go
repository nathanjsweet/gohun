package gohun
/*
#cgo pkg-config: hunspell
#include <stdlib.h>
#include <hunspelld.h>
*/
import "C"
import (
	"unsafe"
	"reflect"
	"runtime"
	"errors"
	"sync"
)

type Gohun struct{
	hunspell unsafe.Pointer
	lock *sync.RWMutex
}

func finalizer(g *Gohun) {
	C.delete_hunspell(g.hunspell)
}

func NewGohun(aff, dic []byte) *Gohun {
	g := new(Gohun)
	g.hunspell = C.new_hunspell((*C.char)(unsafe.Pointer(&aff[0])),(*C.char)(unsafe.Pointer(&dic[0])))
	g.lock = new(sync.RWMutex)
	runtime.SetFinalizer(g, finalizer)
	return g
}

func (g *Gohun) CheckSuggestions(word string) (bool, int, []string) {
	w := C.CString(word)
	defer C.free(unsafe.Pointer(w))
	n := C.int(0)
	b := C.int(0)
	g.lock.RLock()
	s := C.check_suggestions(g.hunspell, w, &n, &b)
	g.lock.RUnlock()
	l := 0
	var r []string
	bo := int(b) == 1
	if !bo {
		l = int(n)
		hdr := reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(s)),
			Len:  l,
			Cap:  l,
		}
		sl := *(*[]*C.char)(unsafe.Pointer(&hdr))
		for i := 0; i < l; i++ {
			r = append(r, C.GoString(sl[i]))
		}
		defer C.free_list(g.hunspell, (***C.char)(unsafe.Pointer(&s)), n)
	}
	return bo, l, r
}

func (g *Gohun) AddDictionary(dictionary []byte) error {
	g.lock.Lock()
	n := C.add_dic(g.hunspell, (*C.char)(unsafe.Pointer(&dictionary[0])));
	g.lock.Unlock()
	var err error 
	if int(n) != 1 {
		err = errors.New("Failed to add dictionary to gohun object.")
	}
	return err
}

func (g *Gohun) AddWord(word string) bool {
	w := C.CString(word)
	defer C.free(unsafe.Pointer(w))
	g.lock.Lock()
	b := C.add_word(g.hunspell, w)
	g.lock.Unlock()
	return int(b) == 1
}

func (g *Gohun) RemoveWord(word string) bool {
	w := C.CString(word)
	defer C.free(unsafe.Pointer(w))
	g.lock.Lock()
	b := C.remove_word(g.hunspell, w)
	g.lock.Unlock()
	return int(b) == 1
}

func (g *Gohun) Stem(word string) (int, []string) {
	w := C.CString(word)
	defer C.free(unsafe.Pointer(w))
	n := C.int(0)
	g.lock.RLock()
	s := C.stem(g.hunspell, w, &n)
	g.lock.RUnlock()
	l := int(n)
	var res []string
	if n > 0 {
		hdr := reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(s)),
			Len:  l,
			Cap:  l,
		}
		sl := *(*[]*C.char)(unsafe.Pointer(&hdr))
		for i := 0; i < l; i++ {
			res = append(res, C.GoString(sl[i]))
		}
		defer C.free_list(g.hunspell, (***C.char)(unsafe.Pointer(&s)), n)
	}
	return l, res
}

func (g *Gohun) Generate(word1, word2 string) (int, []string) {
	w1 := C.CString(word1)
	w2 := C.CString(word2)
	defer C.free(unsafe.Pointer(w1))
	defer C.free(unsafe.Pointer(w2))
	n := C.int(0)
	g.lock.RLock()
	s := C.generate(g.hunspell, w1, w2, &n)
	g.lock.RUnlock()
	l := int(n)
	var res []string
	if n > 0 {
		hdr := reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(s)),
			Len:  l,
			Cap:  l,
		}
		sl := *(*[]*C.char)(unsafe.Pointer(&hdr))
		for i := 0; i < l; i++ {
			res = append(res, C.GoString(sl[i]))
		}
		defer C.free_list(g.hunspell, (***C.char)(unsafe.Pointer(&s)), n)
	}
	return l, res
}

func (g *Gohun) Analyze(word string) (int, []string) {
	w := C.CString(word)
	defer C.free(unsafe.Pointer(w))
	n := C.int(0)
	g.lock.RLock()
	s := C.analyze(g.hunspell, w, &n)
	g.lock.RUnlock()
	l := int(n)
	var res []string
	if n > 0 {
		hdr := reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(s)),
			Len:  l,
			Cap:  l,
		}
		sl := *(*[]*C.char)(unsafe.Pointer(&hdr))
		for i := 0; i < l; i++ {
			res = append(res, C.GoString(sl[i]))
		}
		defer C.free_list(g.hunspell, (***C.char)(unsafe.Pointer(&s)), n)
	}
	return l, res
}
