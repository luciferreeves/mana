package repl

import (
	"bufio"
	"fmt"
	"io"
	"mana/lexer"
	"mana/parser"
)

// PROMPT is the prompt for the REPL.
const PROMPT = ">>> "
const MANA_ICON = `
███╗░░░███╗░█████╗░███╗░░██╗░█████╗░
████╗░████║██╔══██╗████╗░██║██╔══██╗
██╔████╔██║███████║██╔██╗██║███████║
██║╚██╔╝██║██╔══██║██║╚████║██╔══██║
██║░╚═╝░██║██║░░██║██║░╚███║██║░░██║
╚═╝░░░░░╚═╝╚═╝░░╚═╝╚═╝░░╚══╝╚═╝░░╚═╝
`                                                                                

func Start(in io.Reader, out io.Writer) {
	var scanner *bufio.Scanner = bufio.NewScanner(in)
	io.WriteString(out, MANA_ICON + "\n")

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

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Mana Encountered Parser Errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t" + msg + "\n")
	}
}
