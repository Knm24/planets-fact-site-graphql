package utils

func ToPtr[V any](v V) *V {
	return &v
}

func ToValue[V any](p *V) V {
	var v V
	if p != nil {
		v = *p
	}
	return v
}

func ConvertPtrToPtr[V any, P any](ptr *V, cnv func(ptr *V) *P) *P {
	if ptr == nil {
		return nil
	}

	return cnv(ptr)
}

func ConvertValueToPtr[V any, P any](ptr *V, cnv func(v V) *P) *P {
	if ptr == nil {
		return nil
	}

	return cnv(ToValue(ptr))
}
