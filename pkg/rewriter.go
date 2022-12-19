package cvg

type Rewriter interface {
	need()
	provide()
}
