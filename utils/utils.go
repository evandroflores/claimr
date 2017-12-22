package utils

// IfThenElse to add inline condition as Golang does not have ternary ifelse
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
