# pystruct
Pystruct is a module like python struct module

### Install package

``` bash
> go get github.com/WesleiRamos/pystruct
```

### Usage:

`Pack` returns byte array<br>
`Unpack` returns interface array<br>
<br>
NOTICE: Tested only with small values

``` go
package main

import "fmt"
import "reflect"
import "github.com/WesleiRamos/pystruct"

func main() {
	h := "kk eae man"
	b := pystruct.Pack("!hs", len(h), h)

	fmt.Printf("%q\n", b) // "\x00\nkk eae man"

	u, err := pystruct.Unpack("!h", b[:2])
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", b[2:2+reflect.ValueOf(u[0]).Int()]) // kk eae man
}
```
