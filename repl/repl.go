package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/itsmeberwyn/go-interpreter/lexer"
	"github.com/itsmeberwyn/go-interpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf("%v", PROMPT)
		
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
