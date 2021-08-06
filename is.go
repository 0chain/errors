package errors

/* Is - tells whether actual error is targer error
where, actual error can be either Error/withError
if actual error is wrapped error then if any internal error
matches the target error then function results in true
*/
func Is(actual error, target error) bool {
	if target == nil {
		return actual == target
	}
	switch targetError := target.(type) {
	case *Error:
		switch actualError := actual.(type) {
		case *Error:
			if actualError.Code == "" && targetError.Code == "" {
				return actualError.Msg == targetError.Msg
			} else {
				return actualError.Code == targetError.Code
			}
		case *withError:
			return Is(actualError.current, target) || Is(actualError.previous, target)
		default:
			return false
		}
	default:
		return actual == target
	}
}
