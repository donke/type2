package main

import (
	"fmt"
	"io"
	"os"

	"github.com/donke/type2"
	"github.com/donke/wildp"
)

const (
	ok = 0
	ng = 1
)

func main() {
	os.Args = wildp.Args
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "コマンドの構文が誤っています。")
		os.Exit(ng)
	}

	var result = ok
	for _, arg := range os.Args[1:] {
		t2 := type2.New(arg)
		if !t2.Typeable {
			result = ng
			continue
		}
		defer t2.Close()

		if _, err := io.Copy(os.Stdout, t2.File); err != nil {
			fmt.Fprintln(os.Stderr, err)
			result = ng
		}
	}
	os.Exit(result)
}
