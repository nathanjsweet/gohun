SOFLAGS=-fPIC
IDIR=-I./hunspell-distributed/src/hunspell/
LIBS=-L./build/ -lhunspell -lstdc++

libobjects:
	mkdir -p build
	mkdir -p build/obj
	$(CXX) $(SOFLAGS) -c -o build/obj/affentry.o hunspell-distributed/src/hunspell/affentry.cxx
	$(CXX) $(SOFLAGS) -c -o build/obj/affixmgr.o hunspell-distributed/src/hunspell/affixmgr.cxx	
	$(CXX) $(SOFLAGS) -c -o build/obj/csutil.o hunspell-distributed/src/hunspell/csutil.cxx
	$(CXX) $(SOFLAGS) -c -o build/obj/dictmgr.o hunspell-distributed/src/hunspell/dictmgr.cxx
	$(CXX) $(SOFLAGS) -c -o build/obj/filemgr.o hunspell-distributed/src/hunspell/filemgr.cxx
	$(CXX) $(SOFLAGS) -c -o build/obj/hashmgr.o hunspell-distributed/src/hunspell/hashmgr.cxx
	$(CXX) $(SOFLAGS) -c -o build/obj/hunspell.o hunspell-distributed/src/hunspell/hunspell.cxx
	$(CXX) $(SOFLAGS) -c -o build/obj/hunzip.o hunspell-distributed/src/hunspell/hunzip.cxx
	$(CXX) $(SOFLAGS) -c -o build/obj/phonet.o hunspell-distributed/src/hunspell/phonet.cxx
	$(CXX) $(SOFLAGS) -c -o build/obj/replist.o hunspell-distributed/src/hunspell/replist.cxx
	$(CXX) $(SOFLAGS) -c -o build/obj/strmgr.o hunspell-distributed/src/hunspell/strmgr.cxx
	$(CXX) $(SOFLAGS) -c -o build/obj/suggestmgr.o hunspell-distributed/src/hunspell/suggestmgr.cxx

libhunspell: libobjects
	$(CXX) $(SOFLAGS) $(IDIR) -c -o build/obj/hunspelld.o lib/hunspelld.c
	ar rcs build/libhunspell.a build/obj/hunspelld.o build/obj/affentry.o \
		build/obj/affixmgr.o build/obj/csutil.o build/obj/dictmgr.o build/obj/filemgr.o \
		build/obj/hashmgr.o build/obj/hunspell.o build/obj/hunzip.o build/obj/phonet.o \
		build/obj/replist.o build/obj/strmgr.o build/obj/suggestmgr.o

test.c: libhunspell
	$(CXX) $(IDIR) -o lib/test lib/test.c $(LIBS)
clean:
	rm -rf build
	rm -f lib/test
