package repl

import (
	"bufio"
	"fmt"
	"io"
	"mana/lexer"
	"mana/tokens"
)

// PROMPT is the prompt for the REPL.
const PROMPT = "=> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != tokens.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
