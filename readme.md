Gohun
=====
Gohun is a wrapper around a [patched version of Hunspell](https://github.com/nathanjsweet/hunspell-distributed),
which allows Hunspell to be initialized with buffers rather than file locations, thus allowing the dictionaries to be stored
 in more flexible ways.

Installation
------------
You must have pkg-config installed, and you have to include $GOPATH/pkgconfig in its paths to search.
Then you have to run make just once to install the library. Thereafter, as long as pkg-config remains aware of the new path,
you should only have to use the go build tools.
```sh
export PKG_CONFIG_PATH=${PKG_CONFIG_PATH}:${GOPATH}/pkgconfig
cd ./gohun
make
go install gohun
```

Documentation
--------------
```sh
godoc gohun
```
