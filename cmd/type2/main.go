package main

import (
	"fmt"
	"io"
	"os"

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
		fi, err := os.Stat(arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, arg+":", err)
			result = ng
			continue
		}
		if fi.IsDir() {
			continue
		}

		r, err := os.Open(arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, arg+":", err)
			result = ng
			continue
		}
		defer r.Close()

		if _, err = io.Copy(os.Stdout, r); err != nil {
			fmt.Fprintln(os.Stderr, err)
			result = ng
		}
	}
	os.Exit(result)
}
