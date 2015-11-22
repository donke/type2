package type2

import (
	"testing"
)

func TestGuessUTF8(t *testing.T) {
	utf8 := "\x67\x6f\xe3\x81\xaf\xe6\xa5\xbd\xe3\x81\x97"
	actual := guess_jp([]byte(utf8))
	if actual != Utf8 {
		t.Errorf("got %v", actual)
	}
}

func TestGuessEUCJP(t *testing.T) {
	eucjp := "\x67\x6f\xa4\xcf\xb3\xda\xa4\xb7"
	actual := guess_jp([]byte(eucjp))
	if actual != EucJp {
		t.Errorf("got %v", actual)
	}
}

func TestGuessSHIFJIS(t *testing.T) {
	sjis := "\x67\x6f\x82\xcd\x8a\x79\x82\xb5"
	actual := guess_jp([]byte(sjis))
	if actual != ShiftJis {
		t.Errorf("got %v", actual)
	}
}
