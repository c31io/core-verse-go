package cvg

import "sync"

type Scope struct {
	outerScope  *Scope
	innerScopes []*Scope
	variables   sync.Map
}
