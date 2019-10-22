package parser

import (
	"errors"
	"testing"
)

func CheckError(expeted error, actual error, t *testing.T) {
	switch {
	case expeted == nil && actual == nil:
		return
	case expeted != nil && actual == nil:
		t.Errorf("Expected: %s, got success\n", expeted.Error())
	case expeted == nil && actual != nil:
		t.Errorf("Expected: succcess, got %s\n", actual.Error())
	case expeted.Error() != actual.Error():
		t.Errorf("Expected: %s, got %s\n", expeted.Error(), actual.Error())
	}
}

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		rpn []string
		err error
		res int
	}{
		{
			[]string{"10", "-1", "+", "3", "*"},
			nil,
			27,
		},
		{
			[]string{"--10", "-1", "+", "3", "*"},
			errors.New("failed to cast token to int: strconv.Atoi: parsing \"--10\": invalid syntax"),
			27,
		},
		{
			[]string{"10", "-1", "+", "*"},
			errors.New("invailid rpn input"),
			27,
		},

		{
			[]string{"10", "-1", "+", "2"},
			errors.New("invailid rpn input"),
			27,
		},
	}

	for _, test := range tests {
		res, err := EvalRPN(test.rpn)
		if test.err == nil && err == nil {
			if res != test.res {
				t.Fatalf("failed to evaluate rpn string")
			}
		} else {
			CheckError(test.err, err, t)
		}
	}
}
