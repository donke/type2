package main

import (
	"fmt"
	"io"
	"os"

	"github.com/donke/wildp"
)

func main() {
	os.Args = wildp.Args
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "コマンドの構文が誤っています。")
		os.Exit(1)
	}

	for _, arg := range os.Args[1:] {
		r, err := os.Open(arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, arg+":", err)
			os.Exit(1)
		}
		defer r.Close()

		if _, err = io.Copy(os.Stdout, r); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	os.Exit(0)
}
