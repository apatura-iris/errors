package errors

// timeout represents an error on timeout.
type inputOutput struct {
	Err
}

// InputOutputf returns an error which satisfies IsInputOutput().
func InputOutputf(format string, args ...interface{}) error {
	return &inputOutput{wrap(nil, format, " timeout", args...)}
}

// NewInputOutput returns an error which wraps err that satisfies
// IsInputOutput().
func NewInputOutput(err error, msg string) error {
	return &inputOutput{wrap(err, msg, "")}
}

// IsInputOutput reports whether err was created with InputOutputf() or
// NewInputOutput().
func IsInputOutput(err error) bool {
	err = Cause(err)
	_, ok := err.(*inputOutput)
	return ok
}
