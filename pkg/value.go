package cvg

type ValueType int

const (
	ValueScalar ValueType = iota
	ValueTuple
	ValueChoices
)

type Value[N int | float64] struct {
	value     []N
	valueType ValueType
	evaluated chan struct{}
}

func (iv *Value[N]) NewValue(v []N, vt ValueType) *Value[N] {
	intVal := &Value[N]{
		value:     v,
		valueType: vt,
		evaluated: make(chan struct{}),
	}
	if v != nil {
		close(intVal.evaluated)
	}
	return intVal
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
