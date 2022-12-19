package cvg

type ValueType int

const (
	ValueScalar ValueType = iota
	ValueTuple
	ValueChoices
)

// Empty list is zero value.
type Value[N int | float64] struct {
	value     []N
	valueType ValueType
	evaluated chan struct{}
}

func (iv *Value[N]) NewValue(v []N, vt ValueType) *Value[N] {
	value := &Value[N]{
		value:     v,
		valueType: vt,
		evaluated: make(chan struct{}),
	}
	if v != nil {
		close(value.evaluated)
	}
	return value
}

func (iv *Value[N]) getValue() []N {
	<-iv.evaluated
	return iv.value
}

func (iv *Value[N]) getType() ValueType {
	return iv.valueType
}

func (iv *Value[N]) Is(value []N) {
	iv.value = value
	close(iv.evaluated)
}
