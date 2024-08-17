# Go Bash

Run linux bash commands in Go.

## Installation

```shell script
go get github.com/tkdeng/gobash
```

## Usage

```go
import (
  "github.com/tkdeng/gobash"
)

func main(){
  // run bash command
  out, err := bash.Run([]string{`echo`, `test`}, "/pwd", []string{"ENV=var"})

  // run from default directory and no environment vars
  out, err := bash.Run([]string{`echo`, `test`}, "", nil)

  // run raw `bash -c` command
  bash.RunRaw(`echo "test"`, "", nil)

  // run command as specified user
  bash.RunUser(`echo "root test"`, "sudo", "", nil)
  bash.RunUser(`echo "user test"`, "user", "", nil)

  // pipe multiple commands `echo "test" | tee -a "./test.txt"`
  bash.Pipe(".", []string{"echo", "test"}, []string{"tee", "-a", "./test.txt"})
}

```
