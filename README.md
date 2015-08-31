#GoNumbers

GoNumbers is an implementation of numbers helpers for Go-lang.

##Number to Currency:

Options: unit, precision, separator, delimiter

```go
gonumbers.NumberToCurrency(1234567890.50) // "$1,234,567,890.50"

gonumbers.NumberToCurrency(1234567890.506, "precision:2") // "$1,234,567,890.51"

gonumbers.NumberToCurrency(1234567890.50, "precision:2", "unit:CAD$") // "CAD$1.234.567.890,50"

gonumbers.NumberToCurrency(1234567890.50, "precision:2", "unit:$", "separator.", "delimiter:") // "$1234567890.50"
```

##Number to Delimiter:

Options: separator, delimiter

```go
gonumbers.NumberToDelimiter(1234567890) // "1,234,567,890"

gonumbers.NumberToDelimiter(1234567890.506) // "1,234,567,890.506"

gonumbers.NumberToDelimiter(1234567890.34, "separator:-", "delimiter::") // "1:234:567:890-34"

gonumbers.NumberToDelimiter(1234567890.34, "separator:.", "delimiter:") // "1234567890.34"
```

##Number to Human:

Options: separator, delimiter

```go
gonumbers.NumberToHuman(1234) // "1.234 Thousand"

gonumbers.NumberToHuman(12345) // "12.3 Thousand"

gonumbers.NumberToHuman(1234, "separator:,") // "1,234 Thousand"

gonumbers.NumberToHuman(1234567) // "1.234 Million"

gonumbers.NumberToHuman(12345678) // "12.3 Million"

gonumbers.NumberToHuman(123456789) // "123.4 Million"

gonumbers.NumberToHuman(123000000) // "123 Million"

gonumbers.NumberToHuman(1234, "precision:4") // "1234"

gonumbers.NumberToHuman(1234, "precision:5") // "0.01234"

gonumbers.NumberToHuman(1234, "precision:5", "separator:,") // "0,01234"
```

This is a work in progress.