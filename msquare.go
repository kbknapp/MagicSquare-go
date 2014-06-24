package main

import (
	"fmt"
	"log"

	"github.com/kbknapp/gok/math"
)

type MSquare struct {
	math.Matrix
	Num int
}

func NewMSquare(s int, p bool) *MSquare {
	m = MSquare{math.NewMatrix(s)}
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
	m.Num, ok = (sum / m.Size).(int)
	if !ok {
		log.Fatalf("Cannot get magic number")
	}
	log.Printf("Square created\nMagic Num=%d\nSquare=%v\n", m.Num, m.M)
	return m
}

func (s *MSquare) Solve() (solved bool) {
	return
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
