package cvg

type ExprType int

const (
	exprValue ExprType = iota
	exprSequence
	exprScope
	exprFail
	exprAll // for-do
	exprTuple
	exprOne // if-else
	exprChoices
	exprApplication
	exprUnion
)

type Expression struct {
	context   Context
	exprType  ExprType
	inChoices bool
}
