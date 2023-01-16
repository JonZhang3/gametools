package tools

func If(condition bool, trueValue, falseValue any) any {
	if condition {
		return trueValue
	}
	return falseValue
}
