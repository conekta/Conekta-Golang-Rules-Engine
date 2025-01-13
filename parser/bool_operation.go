package parser

type BoolOperation struct {
	BaseOperation
	NullOperation
}

func (o *BoolOperation) get(left Operand, right Operand) (bool, bool, error) {
	if isNil(left) {
		if !o.config.NilToZeroValue {
			return false, false, ErrEvalOperandMissing
		}
		left = false
	}
	leftVal, ok := left.(bool)
	if !ok {
		return false, false, newErrInvalidOperand(left, leftVal)
	}
	rightVal, ok := right.(bool)
	if !ok {
		return false, false, newErrInvalidOperand(right, rightVal)
	}
	return leftVal, rightVal, nil
}

func (o *BoolOperation) EQ(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l == r, nil
}

func (o *BoolOperation) NE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l != r, nil
}
