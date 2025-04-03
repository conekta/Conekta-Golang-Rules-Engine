package parser

type FloatOperation struct {
	BaseOperation
	NullOperation
}

func (o *FloatOperation) get(left Operand, right Operand) (float64, float64, error) {
	if isNil(left) {
		if !o.config.NilToZeroValue {
			return 0, 0, ErrEvalOperandMissing
		}
		left = 0
	}
	leftVal, err := toFloat(left)
	if err != nil {
		return 0, 0, err
	}
	rightVal, err := toFloat(right)
	return leftVal, rightVal, err
}

func (o *FloatOperation) EQ(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l == r, nil
}

func (o *FloatOperation) NE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l != r, nil
}

func (o *FloatOperation) GT(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l > r, nil
}

func (o *FloatOperation) LT(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l < r, nil
}

func (o *FloatOperation) GE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l >= r, nil
}

func (o *FloatOperation) LE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, nil
	}
	return l <= r, nil
}

func (o *FloatOperation) IN(left Operand, right Operand) (bool, error) {
	if isNil(left) {
		if !o.config.NilToZeroValue {
			return false, ErrEvalOperandMissing
		}
		left = 0
	}
	leftVal, err := toFloat(left)
	if err != nil {
		return false, err
	}
	rightVal, ok := right.([]float64)
	if !ok {
		return false, newErrInvalidOperand(right, rightVal)
	}
	for _, num := range rightVal {
		if num == leftVal {
			return true, nil
		}
	}
	return false, nil
}

func ToFloat64(unk interface{}) float64 {
	switch i := unk.(type) {
	case float64:
		return i
	case *float64:
		if i == nil {
			return 0
		}
		return *i
	case float32:
		return float64(i)
	case *float32:
		if i == nil {
			return 0
		}
		return float64(*i)
	case int64:
		return float64(i)
	case *int64:
		if i == nil {
			return 0
		}
		return float64(*i)
	case int:
		return float64(i)
	case *int:
		if i == nil {
			return 0
		}
		return float64(*i)
	case int32:
		return float64(i)
	case *int32:
		if i == nil {
			return 0
		}
		return float64(*i)
	default:
		return 0
	}
}
