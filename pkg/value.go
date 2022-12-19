package cvg

type ValueType int

const (
	ValueScalar ValueType = iota
	ValueTuple
	ValueChoices
)

type Value[T int | float64] struct {
	value     []T
	valueType ValueType
	evaluated chan struct{}
}

func (iv *Value[T]) NewValue(v []T, vt ValueType) *Value[T] {
	intVal := &Value[T]{
		value:     v,
		valueType: vt,
		evaluated: make(chan struct{}),
	}
	if v != nil {
		close(intVal.evaluated)
	}
	return intVal
}

func (iv *Value[T]) getValue() []T {
	<-iv.evaluated
	return iv.value
}

func (iv *Value[T]) getType() ValueType {
	return iv.valueType
}

func (iv *Value[T]) Is(value []T) {
	iv.value = value
	close(iv.evaluated)
}
