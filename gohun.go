package gohun

/*
#cgo LDFLAGS: -L${SRCDIR}/libs -lhunspell -lstdc++ -lm
#include <stdlib.h>
#include "include/hunspelld.h"
*/
import "C"
import (
	"errors"
	"reflect"
	"runtime"
	"sync"
	"unsafe"
)

// Do not try to initialize the Gohun struct yourself,
// always use NewGohun, as a finalizer is set on the struct.
type Gohun struct {
	hunspell unsafe.Pointer
	lock     *sync.RWMutex
}

func finalizer(g *Gohun) {
	C.delete_hunspell(g.hunspell)
}

// Initializing gohun is very easy, simply add the byte arrays of an affix and
// dictionary file as the first two arguments of the constructor. The mechanics
// of the dictionaries gohun processes is fairly simple to understand. Gohun ships
// with US english and Canadian English (look in the examples folder), but tons
// of languages are available for free at open office, http://extensions.services.openoffice.org/dictionary.
func NewGohun(aff, dic []byte) *Gohun {
	g := new(Gohun)
	g.hunspell = C.new_hunspell((*C.char)(unsafe.Pointer(&aff[0])), (*C.char)(unsafe.Pointer(&dic[0])))
	g.lock = new(sync.RWMutex)
	runtime.SetFinalizer(g, finalizer)
	return g
}

// CheckSuggestions checks to see if a word is correct and returns an array
// of possible correct words if it is not.
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

// IsCorrect returns true if a word is spelled correctly.
func (g *Gohun) IsCorrect(word string) bool {
	w := C.CString(word)
	defer C.free(unsafe.Pointer(w))
	g.lock.RLock()
	i := C.is_correct(g.hunspell, w)
	g.lock.RUnlock()
	return int(i) != 0
}

// AddDictionary allows you to add another dictionary, during run time,
// To a current Gohun dictionary (a custom dictionary, for example). This
// is ephemeral.
func (g *Gohun) AddDictionary(dictionary []byte) error {
	g.lock.Lock()
	n := C.add_dic(g.hunspell, (*C.char)(unsafe.Pointer(&dictionary[0])))
	g.lock.Unlock()
	var err error
	if int(n) != 1 {
		err = errors.New("Failed to add dictionary to gohun object.")
	}
	return err
}

// AddWord allows you to add a custom or previously undefined word to the
// current Gohun dictionary, ephemerally.
func (g *Gohun) AddWord(word string) bool {
	w := C.CString(word)
	defer C.free(unsafe.Pointer(w))
	g.lock.Lock()
	b := C.add_word(g.hunspell, w)
	g.lock.Unlock()
	return int(b) == 1
}

// RemoveWrod allows you to remove a defined word from the
// current Gohun dictionary, ephemerally.
func (g *Gohun) RemoveWord(word string) bool {
	w := C.CString(word)
	defer C.free(unsafe.Pointer(w))
	g.lock.Lock()
	b := C.remove_word(g.hunspell, w)
	g.lock.Unlock()
	return int(b) == 1
}

// Stem will return a list of all possible words that a given
// valid word exists in.
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

// Generate returns the variation of a passed word,
// by matching it to the morphological structure of a second word.
// For example, if "telling" and "ran" were passed, "told" would
// be returned. It is possible to receive back the first word as the
// correct result.
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

// Analyze returns a custom array of the morphological information and possibilities
// of a given word. Consult the hunspell docs for further understanding.
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
