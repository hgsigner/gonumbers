#GoNumbers

GoNumbers is an implementation of numbers helpers for Go-lang.

##Number to Currency:

**Options:** unit, precision, separator, delimiter

```go
gonumbers.NumberToCurrency(1234567890.50) // "$1,234,567,890.50"

gonumbers.NumberToCurrency(1234567890.506, "precision:2") // "$1,234,567,890.51"

gonumbers.NumberToCurrency(1234567890.50, "precision:2", "unit:CAD$") // "CAD$1.234.567.890,50"

gonumbers.NumberToCurrency(1234567890.50, "precision:2", "unit:$", "separator.", "delimiter:") // "$1234567890.50"
```

##Number to Delimiter:

**Options:** separator, delimiter

```go
gonumbers.NumberToDelimiter(1234567890) // "1,234,567,890"

gonumbers.NumberToDelimiter(1234567890.506) // "1,234,567,890.506"

gonumbers.NumberToDelimiter(1234567890.34, "separator:-", "delimiter::") // "1:234:567:890-34"

gonumbers.NumberToDelimiter(1234567890.34, "separator:.", "delimiter:") // "1234567890.34"
```

##Number to Human:

**Options:** separator, precision

```go
gonumbers.NumberToHuman(123) // "123"

gonumbers.NumberToHuman(1234) // "1.23 Thousand"

gonumbers.NumberToHuman(12345) // "12.3 Thousand"

gonumbers.NumberToHuman(1234567) // "1.23 Million"

gonumbers.NumberToHuman(1234567890) // "1.23 Billion"

gonumbers.NumberToHuman(1234567890123) // "1.23 Trillion"

gonumbers.NumberToHuman(1234567890123456) // "1.23 Quadrillion"

gonumbers.NumberToHuman(489939, "precision:2") // "490 Thousand"

gonumbers.NumberToHuman(489939, "precision:4") // "489.9 Thousand"

gonumbers.NumberToHuman(489939, "precision:4", "separator:,") // "489,9 Thousand"
```

##Number to Percentage:

**Options:** separator, precision, delimiter

```go
gonumbers.NumberToPercentage(100) // "100.000%"

gonumbers.NumberToPercentage("98") // "98.000%"

gonumbers.NumberToPercentage(100, "precision:0") // "100%"

gonumbers.NumberToPercentage(1000) // "1,000.000%"

gonumbers.NumberToPercentage(1000, "delimiter:.", "separator:,") // "1.000,000%"

gonumbers.NumberToPercentage(302.24398923423, "precision:5") // "302.24399%"

gonumbers.NumberToPercentage("hahahah") // "0.000%"
```

**This is a work in progress.**