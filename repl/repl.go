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
	var scanner *bufio.Scanner = bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		var scanned bool = scanner.Scan()
		
		if !scanned {
			return
		}

		var line string = scanner.Text()
		var l *lexer.Lexer = lexer.New(line)

		for tok := l.NextToken(); tok.Type != tokens.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
