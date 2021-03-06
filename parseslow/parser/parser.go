package parser

import (
	"bufio"
	"io"
)

const (
	bufferSizeKb int = 256
)

type Block interface {
	FirstLineFound(line string) bool
	Parse(*Parser) (keepGoing, keepCurrentLine bool, err error)
}

type Parser struct {
	*bufio.Scanner
	keepCurrentLine bool
	err             error
}

func New(r io.Reader) *Parser {
	p := &Parser{Scanner: bufio.NewScanner(r)}
	p.Buffer(make([]byte, bufferSizeKb*1024/2), bufferSizeKb*1024)
	return p
}

func (p *Parser) ScanBlock(block Block) (keepGoing bool) {
	var err error
	for p.keepCurrentLine || p.Scan() {
		if !block.FirstLineFound(p.Text()) {
			continue
		}
		keepGoing, p.keepCurrentLine, err = block.Parse(p)
		if err != nil {
			p.err = err
		}
		return
	}
	if p.Scanner.Err() != nil {
		p.err = p.Scanner.Err()
		return false
	}
	p.err = nil
	p.keepCurrentLine = false
	return false
}

func (p *Parser) KeepCurrentLine() bool {
	return p.keepCurrentLine
}

func (p *Parser) BlockErr() error {
	return p.err
}
func (p *Parser) Err() error {
	p.err = p.Scanner.Err()
	return p.err
}
