/*
Gohun exposes all of hunspell's functionality and it does away with hunspell's need for files, allowing you to pass raw buffers to create dictionary objects. This, obviously, makes it easier to run a distributed spell checking program that can rely on an SSOT, like a database.

Installation

	make

Warning

Gohun requires golang 1.5+ to build, refer to the repository reamde for details on why this is the case.
*/
package gohun
