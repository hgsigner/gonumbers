#GoNumbers

GoNumbers is an implementation of numbers helpers for Go-lang.

```go
gonumbers.NumberToCurrency(1234567890.50) // "$1,234,567,890.50"

gonumbers.NumberToCurrency(1234567890.506, 2) // "$1,234,567,890.51"

gonumbers.NumberToCurrency(1234567890.50, 2, "CAD$") // "CAD$1.234.567.890,50"

gonumbers.NumberToCurrency(1234567890.50, 2, "$", ".", "") // "$1234567890.50"
```

This is a work in progress.