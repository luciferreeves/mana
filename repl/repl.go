package repl

import (
	"bufio"
	"fmt"
	"io"
	"mana/lexer"
	"mana/parser"
	"mana/evaluator"
)

// PROMPT is the prompt for the REPL.
const PROMPT = ">>> "
const MANA_START = `
███╗░░░███╗░█████╗░███╗░░██╗░█████╗░
████╗░████║██╔══██╗████╗░██║██╔══██╗
██╔████╔██║███████║██╔██╗██║███████║
██║╚██╔╝██║██╔══██║██║╚████║██╔══██║
██║░╚═╝░██║██║░░██║██║░╚███║██║░░██║
╚═╝░░░░░╚═╝╚═╝░░╚═╝╚═╝░░╚══╝╚═╝░░╚═╝
`                                                                                

func Start(in io.Reader, out io.Writer) {
	var scanner *bufio.Scanner = bufio.NewScanner(in)
	io.WriteString(out, MANA_START + "\n")

	for {
		fmt.Fprint(out, PROMPT)
		var scanned bool = scanner.Scan()

		if !scanned {
			return
		}

		var line string = scanner.Text()
		var l *lexer.Lexer = lexer.New(line)
		var p *parser.Parser = parser.New(l)

		var program = p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "ParseError:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t" + msg + "\n")
	}
}
