package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/matsu-chara/go-interpreter/monkey/object"

	"github.com/matsu-chara/go-interpreter/monkey/evaluator"
	"github.com/matsu-chara/go-interpreter/monkey/lexer"
	"github.com/matsu-chara/go-interpreter/monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		if scanner.Err() != nil {
			fmt.Printf("scan error occurred: %+v\n", scanner.Err())
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
