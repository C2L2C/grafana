package eval

var _ Evaluator = new(any)

func Any(anyOf ...Evaluator) Evaluator {
	return any{anyOf: anyOf}
}

type any struct {
	anyOf []Evaluator
}

func (a any) Evaluate(permissions map[string]map[string]struct{}) (bool, error) {
	for _, e := range a.anyOf {
		ok, err := e.Evaluate(permissions)
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}
	return false, nil
}

func (a any) Inject(params map[string]string) error {
	for _, e := range a.anyOf {
		if err := e.Inject(params); err != nil {
			return err
		}
	}
	return nil
}

func (a any) Failed() []permission {
	var failed []permission
	for _, e := range a.anyOf {
		failed = append(failed, e.Failed()...)
	}
	return failed
}
