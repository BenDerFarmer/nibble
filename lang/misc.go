package fishy

import (
	"fmt"
	"strings"
)

var TokenTypeName = map[TokenType]string{
	TokenInvalid: "TokenInvalid",
	TokenSym:     "TokenSym",
	TokenNum:     "TokenNum",
	TokenStr:     "TokenStr",
	TokenOParen:  "TokenOParen",
	TokenCParen:  "TokenCParen",
	TokenComma:   "TokenComma",
}

func (t TokenType) String() string {
	return TokenTypeName[t]
}

func (t Token) String() string {
	return fmt.Sprintf("{Type: %s, Text: %s}", t.TokenType, string(t.Text))
}

func (funcall *Funcall) String() string {
	var result strings.Builder
	fmt.Fprintf(&result, "%s(", funcall.Name)
	for i, arg := range funcall.Args {
		if i > 0 {
			fmt.Fprintf(&result, ", ")
		}
		fmt.Fprintf(&result, "%s", arg.String())
	}
	fmt.Fprintf(&result, ")")
	return result.String()
}

func (expr Expr) String() string {
	switch expr.Type {
	case ExprVoid:
		return "Void"
	case ExprInt:
		return fmt.Sprintf("Int: %d", expr.AsInt)
	case ExprStr:
		return fmt.Sprintf("Str: %s", expr.AsStr)
	case ExprVar:
		return fmt.Sprintf("Var: %s", expr.AsVar)
	case ExprFuncall:
		return "Fun:" + expr.AsFuncall.String()
	}
	panic("unreachable")
}
