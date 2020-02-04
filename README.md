# Nubank Challenge

Challenge documentation for Eduardo Enriquez submission


### Prerequisites

The following tools are needed in order to run this project

```
go version go1.12.1 darwin/amd64

list of dependencies 
====================
bytes
encoding
encoding/base64
encoding/binary
encoding/json
errors
fmt
gopkg.in/validator.v2
internal/bytealg
internal/cpu
internal/fmtsort
internal/poll
internal/race
internal/syscall/unix
internal/testlog
io
io/ioutil
log
math
math/bits
os
path/filepath
reflect
regexp
regexp/syntax
runtime
runtime/internal/atomic
runtime/internal/math
runtime/internal/sys
sort
strconv
strings
sync
sync/atomic
syscall
time
unicode
unicode/utf16
unicode/utf8
unsafe
```

### Installing

in order to first install the project you will need to set up Golang 1.12.1 in your local machine

Golang installation

```
brew install go
```

create directory for the project

```
mkdir ~/go/src/github.com/edenriquez/nubank-challenge
```

install go dependencies, for this you can install listed list of dependencies one by one with:

```
go get -u <dependency>
```

## Running the project

in order to run project you could run under `~/.go/src/github.com/edenriquez/nubank-challenge` in the following way:

```
go run main.go < test.txt
```

or use the binary file attached

```
./authorize < test.txt
```

### Running unit tests

In order to run unit tests you will need to run the following command:

```
go test ./... -cover
```


This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

