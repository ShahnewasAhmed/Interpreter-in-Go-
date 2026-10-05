package repl

import (
	"bufio"
	"fmt"
	"interpreter-go/01/lexer"
	"interpreter-go/01/token"
	"io"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	// io.Reader means the source from where we read input.
	// io.Writer means the destination where we write output.
	// They are interfaces, so the input/output can come from different sources, like terminal, file, buffer, etc.

	// Scanner reads the user's input line by line.
	Scanner := bufio.NewScanner(in)

	for {
		// Show the prompt and wait for the user to type something.
		fmt.Print(PROMPT)

		// Scan() reads the next line from the input.
		// It returns true if a line was successfully scanned, and false if there is no more input or scanning fails.
		Scanned := Scanner.Scan()
		if !Scanned {
			return
		}

		// Get the line that the user just typed.
		line := Scanner.Text()

		// Convert the input line into a Lexer so we can turn the source code into tokens.
		l := lexer.New(line)

		// Keep asking the Lexer for the next token until we reach EOF.
		// Every call to NextToken() gives us the next token.
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			// the for loop (tok := l.NextToken() - initialized the tok with the first Token, checked the condition tok.Type != token.EOF; then the loop body prints that tok, and then in the update part tok = l.NextToken() gets the next Token.
			// This continues until we get EOF.)
			// initialization; condition; update
			fmt.Printf("%+v\n", tok)
		}
	}
}
