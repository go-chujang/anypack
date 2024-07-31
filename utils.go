package anypack

func has0xPrefix(s string) bool {
	return len(s) >= 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X')
}

func ternary[T any](cond bool, trueCase, falseCase T) T {
	if cond {
		return trueCase
	} else {
		return falseCase
	}
}
