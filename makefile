OBJFlAGS=-fPIC
IDIR=-I./hunspell-distributed/src/hunspell/
INCLUDE=$(GOPATH)/include
LIB=./libs
OBJ=./obj

default: libhunspell

libobjects:
	mkdir -p $(OBJ)
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJ)/affentry.o hunspell-distributed/src/hunspell/affentry.cxx
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJ)/affixmgr.o hunspell-distributed/src/hunspell/affixmgr.cxx	
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJ)/csutil.o hunspell-distributed/src/hunspell/csutil.cxx
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJ)/dictmgr.o hunspell-distributed/src/hunspell/dictmgr.cxx
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJ)/filemgr.o hunspell-distributed/src/hunspell/filemgr.cxx
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJ)/hashmgr.o hunspell-distributed/src/hunspell/hashmgr.cxx
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJ)/hunspell.o hunspell-distributed/src/hunspell/hunspell.cxx
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJ)/hunzip.o hunspell-distributed/src/hunspell/hunzip.cxx
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJ)/phonet.o hunspell-distributed/src/hunspell/phonet.cxx
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJ)/replist.o hunspell-distributed/src/hunspell/replist.cxx
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJ)/strmgr.o hunspell-distributed/src/hunspell/strmgr.cxx
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJ)/suggestmgr.o hunspell-distributed/src/hunspell/suggestmgr.cxx

libhunspell: libobjects
	mkdir -p $(LIB)
	$(CXX) -fPIC -shared $(IDIR) -c -O3 -o $(OBJ)/hunspelld.o ./include/hunspelld.c
	ar rcs $(LIB)/libhunspell.a $(OBJ)/hunspelld.o $(OBJ)/affentry.o \
		$(OBJ)/affixmgr.o $(OBJ)/csutil.o $(OBJ)/dictmgr.o $(OBJ)/filemgr.o \
		$(OBJ)/hashmgr.o $(OBJ)/hunspell.o $(OBJ)/hunzip.o $(OBJ)/phonet.o \
		$(OBJ)/replist.o $(OBJ)/strmgr.o $(OBJ)/suggestmgr.o
	rm -rf $(OBJ)
