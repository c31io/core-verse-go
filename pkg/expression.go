package cvg

import "sync"

type ExprType int

const (
	exprUnknown ExprType = iota
	exprValueInt
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
		if expr.outerExpr == nil {
			expr.inter.print(expr.valueInt.Sprint())
		}

	case exprValueFloat:
		if expr.outerExpr == nil {
			expr.inter.print(expr.valueFloat.Sprint())
		}

	case exprSequence:
		// after all evaluated
		var wg sync.WaitGroup
		wg.Add(len(expr.innerExprs))
		for index := range expr.innerExprs {
			go func(index int) {
				defer wg.Done()
				expr.innerExprs[index].Rewriter()
			}(index)
		}
		wg.Wait()
		// become the last inner expression
		last := expr.innerExprs[len(expr.innerExprs)-1]
		expr.exprType = last.exprType
		expr.valueInt = last.valueInt
		expr.valueFloat = last.valueFloat
		expr.Rewriter()

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
