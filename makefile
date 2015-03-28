SOFLAGS=-fPIC
IDIR=-I./hunspell-distributed/src/hunspell/
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
GOBLD=$(GOOS)_$(GOARCH)
GOLIBS=$(GOPATH)/pkg/$(GOBLD)/libs
GOLIBS2=$(GOPATH)/libs
LIB=./libs
OBJ=$(LIB)/obj

default: libhunspell
	go install ./

libobjects:
	mkdir -p $(LIB)
	mkdir -p $(OBJ)
	$(CXX) $(SOFLAGS) -c -o $(OBJ)/affentry.o hunspell-distributed/src/hunspell/affentry.cxx
	$(CXX) $(SOFLAGS) -c -o $(OBJ)/affixmgr.o hunspell-distributed/src/hunspell/affixmgr.cxx	
	$(CXX) $(SOFLAGS) -c -o $(OBJ)/csutil.o hunspell-distributed/src/hunspell/csutil.cxx
	$(CXX) $(SOFLAGS) -c -o $(OBJ)/dictmgr.o hunspell-distributed/src/hunspell/dictmgr.cxx
	$(CXX) $(SOFLAGS) -c -o $(OBJ)/filemgr.o hunspell-distributed/src/hunspell/filemgr.cxx
	$(CXX) $(SOFLAGS) -c -o $(OBJ)/hashmgr.o hunspell-distributed/src/hunspell/hashmgr.cxx
	$(CXX) $(SOFLAGS) -c -o $(OBJ)/hunspell.o hunspell-distributed/src/hunspell/hunspell.cxx
	$(CXX) $(SOFLAGS) -c -o $(OBJ)/hunzip.o hunspell-distributed/src/hunspell/hunzip.cxx
	$(CXX) $(SOFLAGS) -c -o $(OBJ)/phonet.o hunspell-distributed/src/hunspell/phonet.cxx
	$(CXX) $(SOFLAGS) -c -o $(OBJ)/replist.o hunspell-distributed/src/hunspell/replist.cxx
	$(CXX) $(SOFLAGS) -c -o $(OBJ)/strmgr.o hunspell-distributed/src/hunspell/strmgr.cxx
	$(CXX) $(SOFLAGS) -c -o $(OBJ)/suggestmgr.o hunspell-distributed/src/hunspell/suggestmgr.cxx

libhunspell: libobjects
	$(CXX) $(SOFLAGS) $(IDIR) -c -o $(OBJ)/hunspelld.o ./include/hunspelld.c
	ar rcs $(LIB)/libhunspell.a $(OBJ)/hunspelld.o $(OBJ)/affentry.o \
		$(OBJ)/affixmgr.o $(OBJ)/csutil.o $(OBJ)/dictmgr.o $(OBJ)/filemgr.o \
		$(OBJ)/hashmgr.o $(OBJ)/hunspell.o $(OBJ)/hunzip.o $(OBJ)/phonet.o \
		$(OBJ)/replist.o $(OBJ)/strmgr.o $(OBJ)/suggestmgr.o
	mkdir -p $(GOPATH)/pkg
	mkdir -p $(GOPATH)/pkg/$(GOBLD)
	mkdir -p $(GOPATH)/pkg/$(GOBLD)/libs
	cp $(LIB)/libhunspell.a $(GOLIBS)/libhunspell.a
	mkdir -p $(GOPATH)/bin
	mkdir -p $(GOPATH)/bin/libs
	cp $(LIB)/libhunspell.a $(GOPATH)/bin/libs

test.c: libhunspell
	$(CXX) $(IDIR) -o ./lib/test ./include/test.c $(LIB)/libhunspell.a
