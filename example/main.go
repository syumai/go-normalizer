package main

import (
	"fmt"
	"os"

	"github.com/syumai/go-normalizer"
)

type User struct {
	Email     string `normalize:"narrow"`
	FirstName string `normalize:"narrow,capital"`
	LastName  string `normalize:"narrow,capital"`
	WebSite   string `normalize:"pascal"`
}

func (u *User) Normalize() error {
	err := defaultNormalizer.Normalize(u)
	if err != nil {
		return err
	}
	return nil
}

var defaultNormalizer = normalizer.New()

func main() {
	u := &User{
		Email:     "ｓｙｕｍａｉ@syumai.net",
		FirstName: "ｓｙｕ",
		LastName:  "ｍａｉ",
		WebSite:   "syumai_net",
	}

	err := u.Normalize()
	if err != nil {
		fmt.Printf("unexpected error: %v", err)
		os.Exit(1)
	}

	fmt.Println(u)

	// &User{
	//	Email:     "syumai@syumai.net",
	//	FirstName: "Syu",
	//	LastName:  "Mai",
	//	WebSite:   "SyumaiNet",
	// }
}
