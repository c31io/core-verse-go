package cvg

type ExprType int

const (
	exprValueInt ExprType = iota
	exprValueFloat
	exprSequence
	exprScope
	exprFail
	exprAll // for-do
	exprOne // if-else
	exprChoices
	exprApplication
	exprUnify
)

type Expression struct {
	inter      *Interpreter
	valueInt   *Value[int]
	valueFloat *Value[float64]
	scope      Scope
	exprType   ExprType
	outerExpr  *Expression
	innerExprs []Expression

	// context

	inExpression  bool
	inApplication bool
	inScope       bool
	inChoices     bool
}

// Float the binding points of variables.
func (expr *Expression) Floater() {}

// Evaluate by rewriting.
func (expr *Expression) Rewriter() {
	// brute force different paths
	switch expr.exprType {

	case exprValueInt:
		expr.inter.print(expr.valueInt.Sprint())

	case exprValueFloat:
		expr.inter.print(expr.valueFloat.Sprint())

	case exprSequence:

	case exprScope:

	case exprFail:

	case exprAll:

	case exprOne:

	case exprChoices:

	case exprApplication:

	case exprUnify:

	default:
		expr.inter.print("Unknow Value")
	}
}
