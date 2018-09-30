package topp

import (
	"bufio"
	"github.com/lpuig/scopelecspi/parsetop/stat"
)

type chapter struct {
	Found func(rs *bufio.Scanner) bool
	Parse func(s *stat.Stat, rs *bufio.Scanner) error
}

type ParserDef struct {
	Chapters       []chapter
	CurrentChapter int
	Stat           stat.Stat
}

func (b *ParserDef) Found(rs *bufio.Scanner) bool {
	return b.Chapters[0].Found(rs)
}

func (b *ParserDef) Parse(rs *bufio.Scanner) (stat.Stat, error) {
	b.CurrentChapter = 0
	for {
		err := b.Chapters[b.CurrentChapter].Parse(&b.Stat, rs)
		if err != nil {
			return stat.Stat{}, err
		}
		b.CurrentChapter++
		if b.CurrentChapter == len(b.Chapters) {
			break
		}
		for rs.Scan() && !b.Chapters[b.CurrentChapter].Found(rs) {
			continue
		}
		if rs.Err() != nil {
			break
		}
	}
	return b.Stat, rs.Err()
}
