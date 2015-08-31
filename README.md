#GoNumbers

GoNumbers is an implementation of numbers helpers for Go-lang.

##Number to Currency:
```go
gonumbers.NumberToCurrency(1234567890.50) // "$1,234,567,890.50"

gonumbers.NumberToCurrency(1234567890.506) // "$1,234,567,890.51"

gonumbers.NumberToCurrency(1234567890.50, 2, "CAD$") // "CAD$1.234.567.890,50"

gonumbers.NumberToCurrency(1234567890.50, 2, "$", ".", "") // "$1234567890.50"
```

##Number to Delimiter:
```go
gonumbers.NumberToDelimiter(1234567890) // "1,234,567,890"

gonumbers.NumberToDelimiter(1234567890.506) // "1,234,567,890.506"

gonumbers.NumberToDelimiter(1234567890.34, "-", ":") // "1:234:567:890-34"

gonumbers.NumberToDelimiter(1234567890.34, ".", "") // "1234567890.34"
```

This is a work in progress.