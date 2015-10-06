/*
Gohun exposes all of hunspell's functionality and it does away with hunspell's need for files, allowing you to pass raw buffers to create dictionary objects. This, obviously, makes it easier to run a distributed spell checking program that can rely on an SSOT, like a database.

Installation

You must have pkg-config installed, and you have to include $GOPATH/pkgconfig in its paths to search. Then you have to run make just once to install the library. Thereafter, as long as pkg-config remains aware of the new path, you should only have to use the go build tools.

	export PKG_CONFIG_PATH=${PKG_CONFIG_PATH}:${GOPATH}/pkgconfig
	cd ./gohun
	make
	go install gohun

Warning

Gohun requires golang 1.5+ to build, refer to the repository reamde for details on why this is the case.
*/
package gohun
