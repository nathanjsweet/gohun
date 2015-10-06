Gohun
=====
Gohun exposes all of hunspell's functionality and it does away with hunspell's need for files, allowing you to pass raw buffers to create dictionary objects. This, obviously, makes it easier to run a distributed spell checking program that can rely on an SSOT, like a database.

Warning
-------
Gohun now only works with golang version 1.5+, as it has changed to use the new $SRCDIR replacement feature of CGO. I'm not thrilled about this change, but the base golang docker image doesn't come with `pkg-config` (that I can see), and that would be more confusing than just forcing everybody to 1.5+.

Installation
------------
```sh
cd ./gohun
make
```

Documentation
--------------
```sh
godoc gohun
```
