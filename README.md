Package `luhn` implements [Luhn algorithm][luhn-algo] check digit calculation,
and number validation. 

#### Check digit calculation

    checksum, err := luhn.Checksum("123412341234123")
    fmt.Printf("%d, %v", checksum, err)
    // Output:
    // 8, <nil>

#### Number validation

    valid := luhn.Validate("1234123412341238")
    fmt.Println(valid)
    // Output:
    // true

[luhn-algo]: https://en.wikipedia.org/wiki/Luhn_algorithm
