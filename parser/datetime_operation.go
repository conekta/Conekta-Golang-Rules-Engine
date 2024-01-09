package parser

import (
	"time"
)

const timeLayout = "2006-01-02T15:04"

type DateTimeOperation struct {
	NullOperation
}

func (o *DateTimeOperation) getTime(operand Operand) (time.Time, error) {
	switch opVal := operand.(type) {
	case string:
		return time.Parse(timeLayout, opVal)
	case time.Time:
		return opVal, nil
	default:
		return time.Time{}, newErrInvalidOperand(operand, opVal)
	}
}

func (o *DateTimeOperation) get(left Operand, right Operand) (time.Time, time.Time, error) {
	if left == nil {
		return time.Time{}, time.Time{}, ErrEvalOperandMissing
	}
	leftVal, err := o.getTime(left)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	rightVal, err := o.getTime(right)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	return leftVal, rightVal, nil

}

func (o *DateTimeOperation) EQ(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l.Equal(r), nil
}

func (o *DateTimeOperation) NE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return !l.Equal(r), nil
}

func (o *DateTimeOperation) GT(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l.After(r), nil
}

func (o *DateTimeOperation) LT(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l.Before(r), nil
}

func (o *DateTimeOperation) GE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l.After(r) || l.Equal(r), nil
}

func (o *DateTimeOperation) LE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l.Before(r) || l.Equal(r), nil
}
