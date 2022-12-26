package cvg

import "fmt"

// Internally, values are all number slices.
// Their type will influence the behavior of operators.
type ValueType int

const (

	// Scalar Values

	valueInteger  ValueType = iota // 42
	valueFloat                     // 3.14
	valueVariable                  // x
	valuePrimop                    // + - * / < <= > >=

	// Heap Values

	valueTuple
	valueLambda
)

// Value[int | float64]{} blocks the getValue() until evaluated.
type Value[N int | float64] struct {
	fields    []N
	valueType ValueType
	evaluated chan struct{}
}

// Create a Value object with fields and type.
// If the fields slice is not nil, unblock the getValue() right away.
func (val *Value[N]) NewValue(fs []N, vt ValueType) {
	val.fields = fs
	val.valueType = vt
	val.evaluated = make(chan struct{})
	if fs != nil {
		close(val.evaluated)
	}
}

// Block until the value is evaluated.
func (val *Value[N]) getValue() []N {
	<-val.evaluated
	return val.fields
}

// Getter of the valueType.
func (val *Value[N]) getType() ValueType {
	return val.valueType
}

// Value is evalueated as fs with type vt.
func (val *Value[N]) is(fs []N, vt ValueType) {
	val.fields = fs
	val.valueType = vt
	close(val.evaluated)
}

// Zero value or an empty slice is the False of Verse.
func (val *Value[N]) isFail() bool {
	if len(val.fields) == 0 {
		return true
	} else {
		return false
	}
}

// String presentation of the value
func (val *Value[N]) Sprint() string {
	if val.valueType == valueInteger ||
		val.valueType == valueFloat {
		return fmt.Sprintf("%v", val.fields[0])
	}
	return "Unknown Value"
}
