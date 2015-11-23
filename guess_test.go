package type2

import (
	"testing"
)

func TestGuessUTF8(t *testing.T) {
	// goは楽し
	utf8 := "\x67\x6f\xe3\x81\xaf\xe6\xa5\xbd\xe3\x81\x97"
	actual := guess_jp([]byte(utf8))
	if actual != Utf8 {
		t.Errorf("got %v", actual)
	}
}

func TestGuessEUCJP(t *testing.T) {
	// goは楽し
	eucjp := "\x67\x6f\xa4\xcf\xb3\xda\xa4\xb7"
	actual := guess_jp([]byte(eucjp))
	if actual != EucJp {
		t.Errorf("got %v", actual)
	}
}

func TestGuessSHIFJIS(t *testing.T) {
	// goは楽し
	sjis := "\x67\x6f\x82\xcd\x8a\x79\x82\xb5"
	actual := guess_jp([]byte(sjis))
	if actual != ShiftJis {
		t.Errorf("got %v", actual)
	}
}

func TestGuessISO2022JP(t *testing.T) {
	// goは楽し
	jis := "\x67\x6f\x1b\x24\x42\x24\x4f\x33\x5a\x24\x37\x1b\x28\x42"
	actual := guess_jp([]byte(jis))
	if actual != Iso2022Jp {
		t.Errorf("got %v", actual)
	}
}
