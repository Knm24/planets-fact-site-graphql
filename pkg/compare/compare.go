package compare

type Func[T any, P any] func(val1, val2 T, opt P) *bool

func IsLarger[T any, P any](val1, val2 T, opt P, fns []Func[T, P]) bool {
	for _, fn := range fns {
		if boolPtr := fn(val1, val2, opt); boolPtr != nil {
			return *boolPtr
		}
	}

	return false
}
