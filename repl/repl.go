package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/cdarne/monkey/lexer"
	"github.com/cdarne/monkey/parser"
)

const Prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, Prompt)
		ok := scanner.Scan()
		if !ok { // EOF or error
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(out, "error scanning input: %v", err)
			}
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if errors := p.Errors(); len(errors) > 0 {
			printParserErrors(out, errors)
			continue
		}

		fmt.Fprintf(out, "%s\n", program.String())
	}
}

func printParserErrors(out io.Writer, errors []string) {
	fmt.Fprintf(out, "Parsing errors: %s\n", strings.Join(errors, "\n"))
}
