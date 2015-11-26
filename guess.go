package type2

type arc struct {
	next  int
	score float32
}

type dfa struct {
	states  [][]int
	arcs    []arc
	state   int
	score   float32
	charset string
}

func new_dfa(st [][]int, ar []arc, charset string) *dfa {
	return &dfa{st, ar, 0, 1.0, charset}
}

func (d *dfa) alive() bool {
	return d.state >= 0
}

func (d *dfa) next(c int) {
	if d.alive() {
		next := d.states[d.state][c]
		if next < 0 {
			d.state = -1
		} else {
			d.state = d.arcs[next].next
			d.score *= d.arcs[next].score
		}
	}
}

func guess_jp(buf []byte) string {
	eucj := new_dfa(eucjStates, eucjArcs, "EUC-JP")
	sjis := new_dfa(sjisStates, sjisArcs, "Shift_JIS")
	utf8 := new_dfa(utf8States, utf8Arcs, "UTF-8")
    isoj := new_dfa(nil, nil, "ISO2022-JP")

	buflen := len(buf)
	for i := 0; i < buflen; i++ {
		c := int(buf[i])
		if c == 0x1b {
			if i < buflen-1 {
				c = int(buf[i+1])
				i = i + 1
				if c == '$' || c == '(' {
					return isoj.charset
				}
			}
		}
		if eucj.alive() {
			if !sjis.alive() && !utf8.alive() {
				return eucj.charset
			}
			eucj.next(c)
		}
		if sjis.alive() {
			if !eucj.alive() && !utf8.alive() {
				return sjis.charset
			}
			sjis.next(c)
		}
		if utf8.alive() {
			if !sjis.alive() && !eucj.alive() {
				return utf8.charset
			}
			utf8.next(c)
		}
		if !eucj.alive() && !sjis.alive() && !utf8.alive() {
			return ""
		}
	}

	var top *dfa
	if eucj.alive() {
		top = eucj
	}
	if utf8.alive() {
		if top != nil {
			if top.score <= utf8.score {
				top = utf8
			}
		} else {
			top = utf8
		}
	}
	if sjis.alive() {
		if top != nil {
			if top.score <= sjis.score {
				top = sjis
			}
		} else {
			top = sjis
		}
	}
	switch top {
	case eucj:
		return eucj.charset
	case utf8:
		return utf8.charset
	case sjis:
		return sjis.charset
	default:
		return ""
	}
}
