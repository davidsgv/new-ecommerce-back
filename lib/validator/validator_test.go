package validator_test

import (
	"fmt"
	"testing"
	"validator"
)

type validatorTest struct {
	name      string
	tag       string
	value     interface{}
	msg       error
	validator validator.Validator
}

func getTestCases() []validatorTest {
	return []validatorTest{
		{
			name:      "tag id 0",
			tag:       "id",
			value:     0,
			msg:       fmt.Errorf(validator.InvalidId, 0),
			validator: *validator.GetValidator(),
		},
		{
			name:      "tag id negative",
			tag:       "id",
			value:     -1,
			msg:       fmt.Errorf(validator.InvalidId, -1),
			validator: *validator.GetValidator(),
		},
		{
			name:      "tag id succces",
			tag:       "id",
			value:     1,
			msg:       nil,
			validator: *validator.GetValidator(),
		},
		{
			name:      "min string fail",
			tag:       "min=10",
			value:     "fail",
			msg:       fmt.Errorf(validator.MinString, "fail", "10"),
			validator: *validator.GetValidator(),
		},
		{
			name:      "min string succes",
			tag:       "min=10",
			value:     "deberia pasar por que este string es grande",
			msg:       nil,
			validator: *validator.GetValidator(),
		},
		{
			name:      "max string succes",
			tag:       "max=5",
			value:     "ok",
			msg:       nil,
			validator: *validator.GetValidator(),
		},
	}
}

func TestGetValidator(t *testing.T) {
	validators := []*validator.Validator{}
	for i := 0; i < 10; i++ {
		validators = append(validators, validator.GetValidator())
	}

	validator := validators[0]
	for _, v := range validators {
		if validator != v {
			t.Errorf("expected %v, got %v", validator, v)
		}
	}

}

func TestValidate(t *testing.T) {
	testCases := getTestCases()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := tc.validator.Validate(tc.tag, tc.value)

			if err == nil && tc.msg == nil {
				return
			}

			//se valida si el error no es el esperado
			if err.Error() != tc.msg.Error() {
				t.Errorf("expected %v, got %v", tc.msg, err)
			}
		})
	}
}
