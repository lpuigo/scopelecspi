package topp

import "bufio"

type chapter struct {
	Found func(rs *bufio.Scanner) bool
	Parse func(s *Stat, rs *bufio.Scanner) error
}

type ParserDef struct {
	Chapters       []chapter
	CurrentChapter int
	Stat           Stat
}

func (b *ParserDef) Found(rs *bufio.Scanner) bool {
	return b.Chapters[0].Found(rs)
}

func (b *ParserDef) Parse(rs *bufio.Scanner) (Stat, error) {
	b.CurrentChapter = 0
	for {
		err := b.Chapters[b.CurrentChapter].Parse(&b.Stat, rs)
		if err != nil {
			return Stat{}, err
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
