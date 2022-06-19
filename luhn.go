// SPDX-License-Identifier: BSD-3-Clause

// Package luhn implements Luhn algorithm check digit calculation, and number
// validation.
package luhn

import (
	"strconv"
)

// InputError is returned if input contains malformed data.
type InputError int

func (err InputError) Error() string {
	switch int(err) {
	case -1:
		return "input truncated"
	default:
		return "invalid character at position " + strconv.FormatInt(int64(err), 10)
	}
}

// Checksum returns check digit calculated using Luhn algorithm for
// the provided input string that consists of sequence of digits. If
// input data is invalid, an InputError will be returned.
func Checksum(data string) (int, error) {
	if len(data) == 0 {
		return 0, InputError(-1)
	}
	parity := (len(data) - 1) & 1
	sum := 0
	for i, r := range data {
		if r < '0' || '9' < r {
			return 0, InputError(i)
		}
		d := int(r - '0')
		if i&1 == parity {
			d <<= 1
			if d > 9 {
				d -= 9
			}
		}
		if int(^uint(0)>>1)-d < sum {
			return 0, InputError(i)
		}
		sum += d
	}
	c := sum % 10
	if c == 0 {
		return c, nil
	}
	return 10 - c, nil
}

// Validate checks if a sequence of digits ends with a correct check digit
// calculated using Luhn algorithm.
func Validate(data string) bool {
	c, err := Checksum(data[:len(data)-1])
	return err == nil && byte(c+'0') == data[len(data)-1]
}
