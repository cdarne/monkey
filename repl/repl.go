package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/cdarne/monkey/lexer"
	"github.com/cdarne/monkey/token"
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
