# gobgcolor
## Overview
gobgcolor is package to change background color of output word by golang.

## Install
```
go get -u github.com/yuzgit/gobgcolor
```
## How to use
```go
package main

import (
	"fmt"

	"github.com/yuzgit/gobgcolor"
)

func main() {
	fmt.Println(gobgcolor.Red("test"))
}
```