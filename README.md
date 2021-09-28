# Knight's Travails
---

##

## About

Knight's Travails is an application that proposes to solve the problem of finding the shortest path between two squares on a chessboard using only valid knight/camel moves.

The goal is to return the next moves to achieve the desired target square.

The current version only consider that only one piece (knight or camel) exists on the chessboard.

##
---

## Features

- Input values validation
- Accepts chessboard notation both on lowercase or uppercase (e.g. a1 or A1)
- Choose from available chess pieces: Knight or Camel (fairy chess)

##
---

## Installing

There is already a compiled binary for **[Linux](https://github.com/Caik/go-knights-travails/blob/main/dist/knights-travails)** and another one for **[Windows](https://github.com/Caik/go-knights-travails/blob/main/dist/knights-travails.exe)** on the **dist/** directory.
So you only have to download the appropriate binary and run on your machine.

PS: You may need to give execution permission to the binary after downloading it:

 ```bash
# giving execution permission on linux
chmod +x ./knights-travails
```

If you have **Go** configured on your environment, you can build your own binaries as well:

```bash
# building a MacOS on AMD64 binary
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -ldflags '-extldflags "-static" -s -w' -o ./knights-travails-darwin-amd64 cmd/knights-travails/main.go
```

##
---

## Usage

Example:

```bash
# getting the shortest path from A1 to B2 using the knight
./knights-travails A1 B2 knight
```

##
---

## Tests

Running all tests:

```bash
make test
```

or:

```bash
go test ./... -cover
```

##
---

## Authors

* Carlos Henrique Severino (**carloshenrique.dev@gmail.com**)