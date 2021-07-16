package closure

func TestClosure() func(int) int {
	var step int = 0
	return func(_step int) int {
		step += _step
		return step
	}
}
