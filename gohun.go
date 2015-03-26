package gohun
/*
#include "lib/hunspelld.h"
#cgo CFLAGS: -I./hunspell-distributed/src/hunspell/
#cgo LDFLAGS: -L ./build/ -lhunspell
*/
import "C"

import (
	"unsafe"
	"sync"
	"io/ioutil"
)

func main() {
	
}


