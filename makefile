OBJFlAGS=-fPIC
IDIR=-I./hunspell-distributed/src/hunspell/
OBJDIR=./obj

default: lib/libhunspell.a
	go install ./
	rm -rf obj
	rm -rf lib
obj/%.o: hunspell-distributed/src/hunspell/
	mkdir -p $(OBJDIR)
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJDIR)/affentry.o hunspell-distributed/src/hunspell/affentry.cxx
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJDIR)/affixmgr.o hunspell-distributed/src/hunspell/affixmgr.cxx	
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJDIR)/csutil.o hunspell-distributed/src/hunspell/csutil.cxx
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJDIR)/dictmgr.o hunspell-distributed/src/hunspell/dictmgr.cxx
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJDIR)/filemgr.o hunspell-distributed/src/hunspell/filemgr.cxx
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJDIR)/hashmgr.o hunspell-distributed/src/hunspell/hashmgr.cxx
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJDIR)/hunspell.o hunspell-distributed/src/hunspell/hunspell.cxx
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJDIR)/hunzip.o hunspell-distributed/src/hunspell/hunzip.cxx
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJDIR)/phonet.o hunspell-distributed/src/hunspell/phonet.cxx
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJDIR)/replist.o hunspell-distributed/src/hunspell/replist.cxx
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJDIR)/strmgr.o hunspell-distributed/src/hunspell/strmgr.cxx
	$(CXX) $(OBJFLAGS) -c -O3 -o $(OBJDIR)/suggestmgr.o hunspell-distributed/src/hunspell/suggestmgr.cxx
	$(CXX) -fPIC -shared $(IDIR) -c -O3 -o $(OBJDIR)/hunspelld.o ./include/hunspelld.c

lib/libhunspell.a: obj/%.o
	mkdir -p ./lib
	ar rcs $@ $(wildcard obj/*.o)
