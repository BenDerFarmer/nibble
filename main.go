package main

import (
	"fmt"
	"os"

	. "github.com/BenDerFarmer/nibble/lang"
)

func main() {

	data, err := os.ReadFile("example.nibble")
	if err != nil {
		panic(err)
	}

	lexer := NewLexer([]rune(string(data)))

	tokens, err := lexer.Lex()

	if err != nil {
		panic(err)
	}

	parser := NewParser(tokens)

	exprs := parser.ParseAll()

	env := NewEvalEnviroment()
	env.LoadBuildIns()
	env.Functions["say"] = say

	err = env.EvalAll(exprs)

	if err != nil {
		fmt.Println(err)
	}
}

// send message to discord or something like this
func say(env *EvalEnviroment, args []Expr) (Expr, error) {

	for _, arg := range args {
		val, err := env.Eval(arg)
		if err != nil {
			return Expr{}, err
		}

		switch val.Type {
		case ExprStr:
			fmt.Printf("%s", val.AsStr)
		case ExprInt:
			fmt.Printf("%d", val.AsInt)
		default:
			return Expr{}, fmt.Errorf("say() expects ints or strings as args")
		}
	}
	fmt.Printf("\n")

	return Expr{}, nil

}
