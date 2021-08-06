package errors

// Top since errors can be wrapped and stacked,
// it's necessary to get the top level error for tests and validations
func Top(err error) string {
	if err == nil {
		return ""
	}
	current, _ := UnWrap(err)
	return current.Error()
}
