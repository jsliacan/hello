hello
======

Simple module to test new Go modules (from version 1.11).

# Example

Create a folder for your example project somewhere off the `GOPATH`. In terminal:

    $ cd ~
    $ mkdir example

Create a file (the only one in this package with the following contents).

```go
package main

include (
    github.com/jsliacan/hello
)

func main() {
    hello.Greet()
}
```
Save the file and from terminal initialise the go mod file in the `example` folder.

    $ go mod init example

Now you should see `go.mod` file in `example` folder. It only contains one line:

    module example

Build your package.

    $ go build example

It should fetch all dependencies and update the `go.mod` file to something like:

```
module example

require github.com/jsliacan/hello v0.0.0-20190402124957-bcaea9885fa7
```

Another file was created, `go.sum`. It contains something like this:

```
github.com/jsliacan/hello v0.0.0-20190402124957-bcaea9885fa7 h1:G4XfuefRPp/UKxIlqmS0tc1lTlAV8RIEDwRIQLTRiiU=
github.com/jsliacan/hello v0.0.0-20190402124957-bcaea9885fa7/go.mod h1:fIw98EwoHE/VtlWRPsAbYYVtIdbxEHL27PJzmZtzDxI=
```
Anyway, the folder should now contain the binary of your `example` and when you run it you get:

```
Hello stranger!
```

