package repl

import (
	"bufio"
	"fmt"
	"io"
	"xyn/lexer"
	"xyn/token"
)

const PROMPT = ">> "
const NEXT_LINE = ".. "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok.TypeAsString())
            
			if tok.Type == token.EXIT {
				return
			}
		}
	}
}
