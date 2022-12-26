package cvg

type ErrorParsingNumber struct{}
type ErrorInterUninitialized struct{}

func (e ErrorParsingNumber) Error() string {
	return "failed to parse number"
}
func (e ErrorInterUninitialized) Error() string {
	return "interpreter uninitialized"
}
