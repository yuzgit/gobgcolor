# gobgcolor
## Overview
gobgcolor is package to change background color of output word by golang.

## Install
```
go get -u github.com/yuzgit/gobgcolor
```
## How to use
### Code
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

### Output
<img src='https://user-images.githubusercontent.com/39664888/178088437-d03fd19a-77fa-40d1-b67b-90af53b60f50.png' width='40px'>