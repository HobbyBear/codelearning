package nfa

func isMatch(s string, p string) bool {
	nfa := New(p)
	return nfa.Match(s)
}

type NFA struct {
	final, start *State
}

const (
	pTypeSingleLetter    int = 1
	pTypeSingleAnyLetter int = 2
	pTypeAnyLetters      int = 3
	pTypeSpecialLetters  int = 4
)

type Edge struct {
	IsEpsilon bool
	Ch        string
	pType     int
}

func (e *Edge) match(ch rune) bool {
	if e.IsEpsilon {
		return true
	}
	switch e.pType {
	case pTypeAnyLetters, pTypeSingleAnyLetter:
		return true
	default:
		if string(ch) == e.Ch {
			return true
		}
		return false
	}
}

type State struct {
	IsEnd      bool
	Edges      []*Edge
	Edge2State map[*Edge]*State
}

func (s *State) addEdge(state1 *State, edge *Edge) {
	s.Edges = append(s.Edges, edge)
	s.Edge2State[edge] = state1
}

func End(curState *State) bool {
	if curState.IsEnd {
		return true
	}
	for _, edge := range curState.Edges {
		if edge.IsEpsilon {
			return End(curState.Edge2State[edge])
		}
	}
	return false
}

func (s *State) replaceState(state1 *State, newState *State) {
	for edge, st := range s.Edge2State {
		if st == state1 {
			s.Edge2State[edge] = newState
		}
	}
}

func New(words string) *NFA {
	start, end := newState(false), newState(true)
	start.addEdge(end, &Edge{IsEpsilon: true, Ch: ""})

	nfa := &NFA{
		final: end,
		start: start,
	}
	nfa.createNFA(words)
	return nfa
}

func newState(isend bool) *State {
	return &State{
		IsEnd:      isend,
		Edges:      make([]*Edge, 0),
		Edge2State: make(map[*Edge]*State),
	}
}

func (n *NFA) createNFA(words string) {
	// 构建nfa  单个字符  字符*  字符.
	start, final := n.start, n.final
	for index := 0; index < len(words); index++ {
		ch := words[index]
		switch ch {
		case '.':
			if index+1 < len(words) && words[index+1] == '*' {
				begin, end := newState(false), newState(true)
				s := newState(false)
				s.addEdge(end, &Edge{IsEpsilon: true, Ch: ".", pType: 0})
				s.addEdge(s, &Edge{IsEpsilon: false, Ch: ".", pType: pTypeAnyLetters})
				begin.addEdge(s, &Edge{IsEpsilon: true, Ch: ".", pType: 0})
				begin.addEdge(end, &Edge{IsEpsilon: true, Ch: ".", pType: 0})
				start.replaceState(final, begin)
				start, final = s, end
				index++
				continue
			}
			begin, end := newState(false), newState(true)
			begin.addEdge(end, &Edge{IsEpsilon: false, Ch: ".", pType: pTypeSingleAnyLetter})
			start.replaceState(final, begin)
			start, final = begin, end
		default:
			if index+1 < len(words) && words[index+1] == '*' {
				begin, end := newState(false), newState(true)
				s := newState(false)
				n.final.addEdge(s, &Edge{IsEpsilon: true, Ch: ".", pType: 0})
				s.addEdge(end, &Edge{IsEpsilon: true, Ch: ".", pType: 0})
				s.addEdge(s, &Edge{IsEpsilon: false, Ch: string(ch), pType: pTypeSpecialLetters})
				begin.addEdge(s, &Edge{IsEpsilon: true, Ch: ".", pType: 0})
				begin.addEdge(end, &Edge{IsEpsilon: true, Ch: ".", pType: 0})
				start.replaceState(final, begin)
				start, final = s, end
				index++
				continue
			}
			begin, end := newState(false), newState(true)
			begin.addEdge(end, &Edge{IsEpsilon: false, Ch: string(ch), pType: pTypeSingleLetter})
			start.replaceState(final, begin)
			start, final = begin, end
		}
	}
}

func (n *NFA) Match(inputs string) bool {
	index := matchHelper(n.start, []rune(inputs), 0)
	if index != len(inputs) {
		return false
	}
	return true
}

// 返回匹配到chars的数组长度， index为已经匹配到的索引长度
func matchHelper(state *State, chars []rune, index int) int {
	index3 := index
	if state == nil {
		return index3
	}
	if index == len(chars) {
		if End(state) {
			return len(chars)
		} else {
			return -1
		}
	}
	for _, edge := range state.Edges {
		//fmt.Println(edge.Ch)
		if edge.IsEpsilon {
			index3 = matchHelper(state.Edge2State[edge], chars, index)
			if index3 == len(chars) {
				return index3
			}
		} else if edge.match(chars[index]) {
			index3 = matchHelper(state.Edge2State[edge], chars, index+1)
			if index3 == len(chars) {
				return index3
			}
		}
	}
	return index3
}
