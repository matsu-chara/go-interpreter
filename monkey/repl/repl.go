package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/matsu-chara/go-interpreter/monkey/lexer"
	"github.com/matsu-chara/go-interpreter/monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		if scanner.Err() != nil {
			fmt.Printf("scan error occurred: %+v\n")
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
