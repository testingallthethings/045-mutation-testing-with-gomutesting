package mutation

type MyInt struct {
	v int
}

func (i MyInt) IsGreaterThan(v int) bool {
	return i.v > v
}

func (i MyInt) IsLessThan(v int) bool {
	return i.v < v
}

func NewMyInt(v int) MyInt {
	return MyInt{v}
}
