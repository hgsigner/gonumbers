#GoNumbers [![Build Status](https://travis-ci.org/hgsigner/gonumbers.svg?branch=master)](https://travis-ci.org/hgsigner/gonumbers) [![GoDoc](https://godoc.org/github.com/hgsigner/gonumbers?status.svg)](https://godoc.org/github.com/hgsigner/gonumbers)

GoNumbers is an implementation of number helpers for Go-lang. This project is inspired by the ruby library ActiveSupport::NumberHelper.

##Intalling:

```
$ go get github.com/hgsigner/gonumbers
```

##Usage:

###- Number to Currency:

```go
package main

import (
	"fmt"
	"github.com/hgsigner/gonumbers"
)

func main(){
	ntc := gonumbers.ToCurrency()
	ntc.Options(gonumbers.Precision(2))
	ntc_final := ntc.Perform(1234567890.506)
	ftm.Println(ntc_final) // "$1,234,567,890.51"
}
```

**Options:**

```go
gonumbers.Unit(string) 

gonumbers.Precision(int) 

gonumbers.Separator(string) 

gonumbers.Delimiter(string) 
```

###- Number to Delimiter:

```go
package main

import (
	"fmt"
	"github.com/hgsigner/gonumbers"
)

func main(){
	ntd := gonumbers.ToDelimiter()
	ntd.Options(gonumbers.Separator("."), gonumbers.Delimiter(","))
	ntd_final := ntd.Perform(1234567890.506)
	ftm.Println(ntd_final) // "1,234,567,890.506"
}
```

**Options:**

```go
gonumbers.Delimiter(string)

gonumbers.Separator(string) 
```

###- Number to Human:

```go
package main

import (
	"fmt"
	"github.com/hgsigner/gonumbers"
)

func main(){
	nth := gonumbers.ToHuman()
	nth.Options(gonumbers.Separator(","), gonumbers.Precision(2))
	nth_final := nth.Perform(489939)
	ftm.Println(nth_final) // "489,9 Thousand"
}
```

**Options:**

```go
gonumbers.Precision(int) 

gonumbers.Separator(string) 
```

###- Number to Human Size:

```go
package main

import (
	"fmt"
	"github.com/hgsigner/gonumbers"
)

func main(){
	nths := gonumbers.ToHumanSize()
	nths_final, err := nths.Perform(1235551234)
	if err != nil{
		fmt.Println(err)
	}
	ftm.Println(nths_final) // "1.15 GB"
}
```

**Options:**

```go
gonumbers.Delimiter(string) //Defines the delimiter for the number (default => ".")

gonumbers.Precision(int) // Sets precision of the number (default => 3).

gonumbers.Prefix(string) // Defines if it is binary or si (default => binary)
```

###- Number to Percentage:

```go
package main

import (
	"fmt"
	"github.com/hgsigner/gonumbers"
)

func main(){
	ntp := gonumbers.ToPercentage()
	ntp.Options(gonumbers.Separator(","), gonumbers.Delimiter("."))
	ntp_final := ntp.Perform(1000)
	ftm.Println(ntp_final) // "1.000,000%"
}
```

**Options:**

```go
gonumbers.Precision(int):

gonumbers.Separator(string):
 
gonumbers.Delimiter(string):
```

###- Number to Phone:

```go
package main

import (
	"fmt"
	"github.com/hgsigner/gonumbers"
)

func main(){
	ntph := gonumbers.ToPhone()
	ntph.Options(gonumbers.CountryCode("1"))
	ntph_final, err := ntph.Perform(1235551234)
	if err != nil{
		fmt.Println(err)
	}
	ftm.Println(ntph_final) // "+1-123-555-1234"
}
```

**Options:**

```go
gonumbers.Delimiter(string) // Defines the delimiter for the number. Default is "-"

gonumbers.AreaCode(bool) // Prints the area code in parenthesis.

gonumbers.Extension(string) // Appends extension to phone number. Ex. +1(123) 555-6789 x 4545

gonumbers.CountryCode(string) // Prepends country code to phone number (ex. +1)

gonumbers.DigitsSize(int) // Sets the number of digits of the second part of the phone (digits_count:5). Ex. 1234555556789 => 1234-55555-6789
```

- - -
For more information, please refer to the [docs.](https://godoc.org/github.com/hgsigner/gonumbers)
- - -
##Licensing
You can use or modify it to meet your needs.