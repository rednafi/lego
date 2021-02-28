## Update & Installation

```bash
#!/usr/bin/env bash

# Added bash strict mode.
set -euo pipefail

# Remove old golang versions.
whereis go | xargs -n1 sudo rm -rf

# Download golang.
curl -OJL https://golang.org/dl/go1.16.linux-amd64.tar.gz

# Install.
sudo tar -C /usr/local -xzf go1.16.linux-amd64.tar.gz

# Show golang.
go version
```

## Add Export Path to Bash/Fish

In bash, add the following variable to your `~/.bashrc`.

```bash
export PATH=$PATH:/usr/local/go/bin
```

In fish, add the following variable to your `~/.config/fish/config.fish`

```fish
# Adds go path
set -x PATH $PATH:/usr/local/go/bin
```

## Project Nursery

* Create a folder named `go` in `~/`. This is the nursery where all of your `go` projects should live.
* The folder structure should look like the following:
  ```
  .
  ├── bin
  ├── pkg
  └── src
  ```

  Your project's source code should always live in the `src` directories. You should have a single top-level directory for
  each one of your projects. Those top-level directories can have additional packages (directories) and `*.go` files.

 ## Project Folder Structure

 This is the structure of a toy project. It's a simple calculator that does four basic integer operations:

 * Addition
 * Subtraction
 * Multiplication
 * Division (Floor division)

 Look carefully how the packages (directories) are structured inside the root directory of the project.

```
src
└── calc
    ├── addsub
    │   ├── add.go
    │   └── sub.go
    ├── muldiv
    │   ├── div.go
    │   └── mul.go
    ├── go.mod
    └── main.go

3 directories, 6 files
```

## Import Structure

Just take a look at the contents of the following files and take a look at how files can be imported.

* `calc/addsub/add`
  ```go
  package addsub

  // Add two integers.
  func Add(a, b int) (int, error) {
    return a + b, nil
  }
  ```

* `calc/addsub/sub`
  ```go
  package addsub

  // Sub two integers.
  func Sub(a, b int) (int, error) {
    return a - b, nil
  }

  ```
* `calc/muldiv/mul`
  ```go
  package muldiv

  // Mul two integers.
  func Mul(a, b int) (int, error) {
    return a * b, nil
  }

  ```
* `calc/muldiv/div`
  ```go
  package muldiv

  import "math"

  // Div two integers.
  func Div(a, b int) (float64, error) {
    return math.Floor(float64(a) / float64(b)), nil
  }
  ```
* `main.go`
  ```go
  package main

  import (
    "calc/addsub"
    "calc/muldiv"
    "fmt"
  )

  // CheckError does rudimentary error checking.
  func CheckError(e error) interface{} {
    if e != nil {
      return fmt.Errorf("an error happened")
    }
    return nil
  }

  func main() {

    r, err := addsub.Add(20, 44)
    CheckError(err)
    fmt.Printf("Value of add func is %v \n", r)

    r, err = addsub.Sub(20, 44)
    CheckError(err)
    fmt.Printf("Value of sub func is %v \n", r)

    r, err = muldiv.Mul(20, 44)
    CheckError(err)
    fmt.Printf("Value of mul func is %v \n", r)

    s, err := muldiv.Div(200, 44)
    CheckError(err)
    fmt.Printf("Value of div func is %v \n", s)
  }
  ```
Now if you run `go run main.go` you should be looking at the following output:

```
Value of add func is 64
Value of sub func is -24
Value of mul func is 880
Value of div func is 4
```


