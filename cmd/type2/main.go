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

	var multiple = false
	if len(os.Args[1:]) > 1 {
		multiple = true
	}

	var result = ok
	for _, arg := range os.Args[1:] {
		t2 := type2.New(arg)
		if !t2.Typeable {
			result = ng
			fmt.Fprintln(os.Stderr, "指定されたファイルが見つかりません。")
			if multiple {
				fmt.Fprintln(os.Stderr, "処理中にエラーが発生しました: "+t2.Name)
			}
			continue
		}
		defer t2.Close()

		if multiple {
			fmt.Fprintf(os.Stderr, "%s\n\n\n", t2.Name)
		}
		if _, err := io.Copy(os.Stdout, t2.File); err != nil {
			fmt.Fprintln(os.Stderr, err)
			result = ng
			if multiple {
				fmt.Fprintln(os.Stderr, "処理中にエラーが発生しました: "+t2.Name)
			}
		}
	}
	os.Exit(result)
}
