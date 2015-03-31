SOFLAGS=-fPIC
IDIR=-I./hunspell-distributed/src/hunspell/
INCLUDE=$(GOPATH)/include
LIB=$(GOPATH)/pkg/$(shell go env GOOS)_$(shell go env GOARCH)
OBJ=./obj

default: libhunspell
	mkdir -p $(GOPATH)/pkgconfig
	sed 's#$${GOPATH}#$(GOPATH)#g;s#$${GOOS}#$(shell go env GOOS)#g;s#$${GOARCH}#$(shell go env GOARCH)#g;' \
		./include/hunspell.pc > $(GOPATH)/pkgconfig/hunspell.pc
	mkdir -p $(GOPATH)/include
	cp ./include/hunspelld.h $(GOPATH)/include/hunspelld.h

libobjects:
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
	mkdir -p $(GOPATH)/pkg
	mkdir -p $(LIB)
	$(CXX) $(SOFLAGS) $(IDIR) -c -o $(OBJ)/hunspelld.o ./include/hunspelld.c
	ar rcs $(LIB)/libhunspell.a $(OBJ)/hunspelld.o $(OBJ)/affentry.o \
		$(OBJ)/affixmgr.o $(OBJ)/csutil.o $(OBJ)/dictmgr.o $(OBJ)/filemgr.o \
		$(OBJ)/hashmgr.o $(OBJ)/hunspell.o $(OBJ)/hunzip.o $(OBJ)/phonet.o \
		$(OBJ)/replist.o $(OBJ)/strmgr.o $(OBJ)/suggestmgr.o
	rm -rf $(OBJ)

test.c: libhunspell
	$(CXX) $(IDIR) -o ./lib/test ./include/test.c $(LIB)/libhunspell.a
