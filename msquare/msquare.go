package msquare

import (
	"fmt"
	"log"

	"github.com/kbknapp/gok/math"
)

type MSquare struct {
	math.Matrix
	Num  int
	List []int
}

func NewMSquare(s int, p bool) *MSquare {
	m := MSquare{math.NewMatrix(s), 0, make([]int, s*s)}
	if p {
		m.M = math.NPrimes(s * s)
	} else {
		for i, _ := range m.M {
			m.M[i] = i + 1
		}
	}
	sum := 0
	for i, _ := range m.M {
		sum += m.M[i]
	}
	copy(m.List, m.M)
	m.Num = int(sum / m.Size)
	log.Print("Square created\n")
	log.Printf("Magic Num=%d\n", m.Num)
	log.Printf("Square=%v\n", m.M)
	log.Printf("Nums=%v\n", m.List)

	return &m
}

func (s *MSquare) Solve() (solved bool) {
	for solved {
		i := s.Size - 1
		for i > -1 {
			for s.sumIndices(s.Rows[i]) != s.Num {
				if ok := s.incrementIndices(); ok {
					return false
				} else {
					i = s.Size - 1
				}
			}
			i -= 1
		}
		solved = true
		for _, col := range s.Cols {
			for s.sumIndices(col) != s.Num {
				if ok := s.incrementIndices(); ok {
					return false
				} else {
					solved = false
					break
				}
			}
		}
		if solved {
			for _, diag := range s.Diags {
				for s.sumIndices(diag) != s.Num {
					if ok := s.incrementIndices(); ok {
						return false
					} else {
						solved = false
						break
					}
				}
			}
		}
	}

	return true
}

func (s *MSquare) sumIndices(indices []int) (sum int) {
	for _, ind := range indices {
		sum += s.M[ind]
	}
	return
}

func (s *MSquare) incrementIndices() bool {
	s_len := s.Size * s.Size
	s_dup := make([]int, s_len)
	copy(s_dup, s.M)
	r_i := s_len - 1
	n_i := s_len

	for {
		for n_i == s_len {
			n_i, _ := indexOf(s_dup, s_dup[r_i])
			n_i += 1
			if n_i == s_len {
				s_dup[r_i] = 0
				r_i -= 1
				if r_i < 0 {
					return true
				}
			}
		}
		for {
			if !isIn(s_dup, s.List[n_i]) {
				break
			}
			n_i += 1
			if n_i == s_len {
				s_dup[r_i] = 0
				r_i -= 1
				break
			} else {
				s_dup[r_i] = s.List[n_i]
				if isIn(s_dup, 0) {
					r_i += 1
					n_i = 0
				} else {
					break
				}
			}
		}
	}
	copy(s.M, s_dup)
	return false
}

func isIn(seq []int, elem int) bool {
	for _, s := range seq {
		if s == elem {
			return true
		}
	}
	return false
}

func indexOf(seq []int, elem int) (int, bool) {
	for i, _ := range seq {
		if seq[i] == elem {
			return i, true
		}
	}
	return -1, false
}

func (s *MSquare) Print() {
	s.printSep()

	for i, _ := range s.M {
		if (i+1)%s.Size == 0 {
			fmt.Printf("|%s|\n", s.getCellString(s.M[i]))
			s.printSep()
		} else {
			fmt.Printf("|%s", s.getCellString(s.M[i]))
		}
	}
}

func (ms *MSquare) getCellString(v int) string {
	s := ""
	switch {
	case v < 10:
		s = fmt.Sprintf("  %d ", v)
	case v < 100:
		s = fmt.Sprintf(" %d ", v)
	case v < 1000:
		s = fmt.Sprintf(" %d", v)
	case v < 10000:
		s = fmt.Sprintf("%d", v)
	default:
		s = "    "
	}
	return s
}

func (s *MSquare) printSep() {
	for i := 0; i < s.Size; i++ {
		fmt.Print("+----")
	}
	fmt.Println("+")
}
