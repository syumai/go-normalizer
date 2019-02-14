# go-normalizer

* A Go package to normalize struct fields by tags.

## Usage

```go
package main

import (
	"fmt"
	
	"github.com/syumai/go-normalizer"
)

type User struct {
	Email     string `normalize:"narrow"`
	FirstName string `normalize:"narrow,capital"`
	LastName  string `normalize:"narrow,capital"`
	WebSite   string `normalize:"pascal"`
}

func main() {
	u := &User{
		Email:     "ｓｙｕｍａｉ@syumai.net",
		FirstName: "ｓｙｕ",
		LastName:  "ｍａｉ",
		WebSite:   "syumai_net",
	}

    n = normalizer.New()
	n.Normalize(u)

	fmt.Println(u)

	// &User{
	//	Email:     "syumai@syumai.net",
	//	FirstName: "Syu",
	//	LastName:  "Mai",
	//	WebSite:   "SyumaiNet",
	// }
}
```

## Tag lists

```go
"lower":   `TEST => test`
"upper":   `test => TEST`
"capital": `test test => Test Test`
"title":   `test test => Test Test`
"snake":   `TestTest => test_test`
"camel":   `test_test => testTest`
"pascal":  `test_test => TestTest`
"kebab":   `test_test => test-test`
"widen":   `test => ｔｅｓｔ`
"narrow":   `ｔｅｓｔ => test`
```
