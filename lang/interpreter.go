package fishy

import (
	"fmt"
)

type EvalEnviroment struct {
	Functions map[string]func(*EvalEnviroment, []Expr) (Expr, error)
	Variables map[string]Expr
}

func NewEvalEnviroment() EvalEnviroment {

	functions := make(map[string]func(*EvalEnviroment, []Expr) (Expr, error))
	variables := make(map[string]Expr)

	return EvalEnviroment{Functions: functions, Variables: variables}

}

func (env *EvalEnviroment) LoadBuildIns() {

	env.Functions["println"] = println
	env.Functions["add"] = add
	env.Functions["http"] = httpFunc
	env.Functions["let"] = let
	env.Functions["fun"] = fun

}

func (env *EvalEnviroment) EvalAll(exprs []Expr) error {
	for _, expr := range exprs {
		_, err := env.Eval(expr)
		if err != nil {
			return err
		}
	}

	return nil
}

func (env *EvalEnviroment) Eval(expr Expr) (Expr, error) {

	switch expr.Type {
	case ExprFuncall:
		fn, ok := env.Functions[expr.AsFuncall.Name]

		if !ok {
			return Expr{}, fmt.Errorf("Function: %s not found \n", expr.AsFuncall.Name)

		}

		return fn(env, expr.AsFuncall.Args)
	case ExprVar:

		vars, ok := env.Variables[expr.AsVar]

		if !ok {
			return Expr{}, fmt.Errorf("Variable: %s not found \n", expr.AsVar)
		}

		return vars, nil

	default:
		return expr, nil
	}
}
