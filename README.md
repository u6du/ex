# ex

Usage

```
package main

import (
	"io/ioutil"
	"os"

	. "github.com/u6du/ex"
)

func main() {
	filepath := "test.txt"

	_, err := os.Stat(filepath)
	Panic(err)

	Panic(ioutil.WriteFile(filepath, []byte("test"), 0644))

}
```


