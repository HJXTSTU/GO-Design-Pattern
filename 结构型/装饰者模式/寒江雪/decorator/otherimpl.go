package decorator

type DecoratorFunc func(float64)float64

func DecFunc(dec DecoratorFunc)DecoratorFunc{
	return func(f float64) float64 {
		result := dec(f)
		return result
	}
}


