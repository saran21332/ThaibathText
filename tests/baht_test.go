package tests

import (
	"bath_to_text/converter"
	"testing"

	"github.com/shopspring/decimal"
)

func TestThaiBahtText(t *testing.T) {
	tests := []struct {
		input    decimal.Decimal
		expected string
	}{
		{
			input:    decimal.NewFromFloat(0),
			expected: "ศูนย์บาทถ้วน",
		},
		{
			input:    decimal.NewFromFloat(1),
			expected: "หนึ่งบาทถ้วน",
		},
		{
			input:    decimal.NewFromFloat(10),
			expected: "สิบบาทถ้วน",
		},
		{
			input:    decimal.NewFromFloat(11),
			expected: "สิบเอ็ดบาทถ้วน",
		},
		{
			input:    decimal.NewFromFloat(21),
			expected: "ยี่สิบเอ็ดบาทถ้วน",
		},
		{
			input:    decimal.NewFromFloat(101),
			expected: "หนึ่งร้อยเอ็ดบาทถ้วน",
		},
		{
			input:    decimal.NewFromFloat(1005),
			expected: "หนึ่งพันห้าบาทถ้วน",
		},
		{
			input:    decimal.NewFromFloat(1000000),
			expected: "หนึ่งล้านบาทถ้วน",
		},
		{
			input:    decimal.NewFromFloat(1000000000000),
			expected: "หนึ่งล้านล้านบาทถ้วน",
		},
		{
			input:    decimal.NewFromFloat(33333.75),
			expected: "สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์",
		},
	}

	for _, test := range tests {
		result := converter.ThaiBahtText(test.input)

		if result != test.expected {
			t.Errorf(
				"input %s: expected %s but got %s",
				test.input.String(),
				test.expected,
				result,
			)
		}
	}
}
