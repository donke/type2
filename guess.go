package type2

const (
    EucJp = iota
    ShiftJis
    Utf8
    Iso2022Jp
    Unknown
)

type arc struct {
	next  int
	score float32
}

type dfa struct {
	states [][]int
	arcs   []arc
	state  int
	score  float32
}

func new_dfa(st [][]int, ar []arc) *dfa {
	return &dfa{st, ar, 0, 1.0}
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

func guess_jp(buf []byte) int {
	eucj := new_dfa(eucjStates, eucjArcs)
	sjis := new_dfa(sjisStates, sjisArcs)
	utf8 := new_dfa(utf8States, utf8Arcs)

	buflen := len(buf)
	for i := 0; i < buflen; i++ {
		c := int(buf[i])
		if c == 0x1b {
			if i < buflen-1 {
				c = int(buf[i+1])
				i = i + 1
				if c == '$' || c == '(' {
					return Iso2022Jp
				}
			}
		}
		if eucj.alive() {
			if !sjis.alive() && !utf8.alive() {
				return EucJp
			}
			eucj.next(c)
		}
		if sjis.alive() {
			if !eucj.alive() && !utf8.alive() {
				return ShiftJis
			}
			sjis.next(c)
		}
		if utf8.alive() {
			if !sjis.alive() && !eucj.alive() {
				return Utf8
			}
			utf8.next(c)
		}
		if !eucj.alive() && !sjis.alive() && !utf8.alive() {
			return Unknown
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
		return EucJp
	case utf8:
		return Utf8
	case sjis:
		return ShiftJis
	default:
		return Unknown
	}
}
