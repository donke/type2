package type2

import (
	"testing"
)

func TestGuess(t *testing.T) {
	utf8 := "\x67\x6f\xe3\x81\xaf\xe6\xa5\xbd\xe3\x81\x97"
	actual := guess_jp([]byte(utf8))
	if actual != Utf8 {
		t.Errorf("got %v", actual)
	}
}
