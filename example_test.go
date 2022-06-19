// SPDX-License-Identifier: BSD-3-Clause

package luhn_test

import (
	"fmt"

	"github.com/ivanrad/luhn"
)

func ExampleChecksum() {
	checksum, err := luhn.Checksum("123412341234123")
	fmt.Printf("%d, %v", checksum, err)
	// Output:
	// 8, <nil>
}

func ExampleValidate() {
	valid := luhn.Validate("1234123412341238")
	fmt.Println(valid)
	// Output:
	// true
}
