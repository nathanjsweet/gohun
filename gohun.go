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
)

type Gohun struct{
	hunspell unsafe.Pointer
}

func finalizer(g *Gohun) {
	C.delete_hunspell(g.hunspell)
}

func NewGohun(aff, dic []byte) *Gohun {
	g := new(Gohun)
	g.hunspell = C.new_hunspell((*C.char)(unsafe.Pointer(&aff[0])),(*C.char)(unsafe.Pointer(&dic[0])))
	runtime.SetFinalizer(g, finalizer)
	return g
}

func (g *Gohun) CheckSuggestions(word string) (bool, int, []string) {
	w := C.CString(word)
	defer C.free(unsafe.Pointer(w))
	n := C.int(0)
	b := C.int(0)
	s := C.check_suggestions(g.hunspell, w, &n, &b)
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

func (g *Gohun) AddDictionary(dictionary []byte) bool {
	n := C.add_dic(g.hunspell, (*C.char)(unsafe.Pointer(&dictionary[0])));
	return int(n) == 1;
}

func (g *Gohun) AddDictionaryString(dictionary string) bool {
	d := C.CString(dictionary);
	defer C.free(unsafe.Pointer(d));
	n := C.add_dic(g.hunspell, d);
	return int(n) == 1;
}
