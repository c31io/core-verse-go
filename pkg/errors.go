package cvg

type ErrorParsingNumber struct{}
type ErrorUninitialized struct{}
type ErrorUnknownExpression struct{}
type ErrorBracketsNotMatch struct{}

func (e ErrorParsingNumber) Error() string {
	return "failed to parse number"
}
func (e ErrorUninitialized) Error() string {
	return "interpreter uninitialized"
}
func (e ErrorUnknownExpression) Error() string {
	return "unknown expression"
}
func (e ErrorBracketsNotMatch) Error() string {
	return "brackets do not match"
}
