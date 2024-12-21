package fishy

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func println(env *EvalEnviroment, args []Expr) (Expr, error) {

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
			return Expr{}, fmt.Errorf("println() expects ints or strings as args")
		}
	}
	fmt.Printf("\n")

	return Expr{}, nil

}

func add(env *EvalEnviroment, args []Expr) (Expr, error) {

	result := 0

	for _, arg := range args {
		val, err := env.Eval(arg)
		if err != nil {
			return Expr{}, err
		}

		if val.Type != ExprInt {
			return Expr{}, fmt.Errorf("Function add() expects int as input")
		}

		result += val.AsInt
	}
	return Expr{Type: ExprInt, AsInt: result}, nil

}

func httpFunc(context *EvalEnviroment, args []Expr) (Expr, error) {
	var url strings.Builder

	for _, arg := range args {
		val, err := context.Eval(arg)
		if err != nil {
			return Expr{}, err
		}
		if val.Type != ExprStr {
			return Expr{}, fmt.Errorf("http() expects its arguments to be strings")
		}
		fmt.Fprint(&url, val.AsStr)
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return Expr{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Expr{}, err
	}

	return Expr{
		Type:  ExprStr,
		AsStr: string(body),
	}, nil

}

func let(env *EvalEnviroment, args []Expr) (Expr, error) {
	if len(args) != 2 {
		return Expr{}, fmt.Errorf("let() expects two arguments")
	}

	if args[0].Type != ExprVar {
		return Expr{}, fmt.Errorf("First argument of let() has to be variable name")
	}

	name := args[0].AsVar

	value, err := env.Eval(args[1])
	if err != nil {
		return Expr{}, err
	}

	env.Variables[name] = value

	return Expr{}, nil
}

func fun(env *EvalEnviroment, args []Expr) (Expr, error) {
	if len(args) < 2 {
		return Expr{}, fmt.Errorf("fun() expects at least two arguments")
	}

	if args[0].Type != ExprVar {
		return Expr{}, fmt.Errorf("First argument of fun() has to be variable name")
	}

	name := args[0].AsVar

	env.Functions[name] = func(ee *EvalEnviroment, e []Expr) (Expr, error) {

		if len(e) != 0 {
			return Expr{}, fmt.Errorf("%s does not expect arguments", name)
		}

		return Expr{}, ee.EvalAll(args[1:])
	}

	return Expr{}, nil
}
