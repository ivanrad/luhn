// SPDX-License-Identifier: BSD-3-Clause

package luhn

import (
	"errors"
	"testing"
)

func TestChecksum(t *testing.T) {
	tests := []struct {
		input    string
		checksum int
	}{
		{"0", 0},
		{"1", 8},
		{"12", 5},
		{"123", 0},
		{"1234", 4},
		{"853", 2},
		{"7992739871", 3},
		{"12345678901", 5},
		{"400360000000001", 4},
		{"123412341234123", 8},
	}

	for _, tc := range tests {
		got, err := Checksum(tc.input)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if got != tc.checksum {
			t.Errorf("got %v; expected %v", got, tc.checksum)
		}
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		input string
		valid bool
	}{
		{"00", true},
		{"18", true},
		{"125", true},
		{"1230", true},
		{"12344", true},
		{"8532", true},
		{"79927398713", true},
		{"123456789015", true},
		{"4003600000000014", true},
		{"1234123412341238", true},
		{"490154203237518", true},
		{"01", false},
		{"10", false},
		{"123456789010", false},
		{"123456789011", false},
		{"123456789012", false},
		{"123456789013", false},
		{"123456789014", false},
		{"123456789016", false},
		{"123456789017", false},
		{"123456789018", false},
		{"123456789019", false},
	}
	for _, tc := range tests {
		valid := Validate(tc.input)
		if valid != tc.valid {
			t.Errorf("input %q: got %v; expected %v", tc.input, valid, tc.valid)
		}
	}
}

func TestInvalidInput(t *testing.T) {
	tests := []struct {
		input string
		err   int
	}{
		{"", -1},
		{"799x7398713", 3},
		{"7992739871.", 10},
	}
	for _, tc := range tests {
		_, err := Checksum(tc.input)
		var inputErr InputError
		switch {
		case err == nil:
			t.Errorf("got nil; expected InputError: %v", tc.input)
		case errors.As(err, &inputErr):
			if int(inputErr) != tc.err {
				t.Errorf("InputError: got %v (%d); expected %v (%d)", inputErr, int(inputErr), tc.err, int(tc.err))
			}
		default:
			t.Errorf("unexpected error: %v", err)
		}
	}
}
