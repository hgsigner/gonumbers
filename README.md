#GoNumbers [![Build Status](https://travis-ci.org/hgsigner/gonumbers.svg?branch=master)](https://travis-ci.org/hgsigner/gonumbers)

GoNumbers is an implementation of number helpers for Go-lang. This project is inspired by the ruby library ActiveSupport::NumberHelper.

##Intalling:

```
$ go get github.com/hgsigner/gonumbers
```

##Usage:

```go
package main

import (
	"fmt"
	"github.com/hgsigner/gonumbers"
)

func main(){
	ntc := gonumbers.NewNumberToCurrency()
	ntc.Options(gonumbers.Precision(2))
	ntc_final := ntc.Perform(1234567890.506)
	ftm.Println(ntc_final) // "$1,234,567,890.51"

	ntd := gonumbers.NewNumberToDelimiter()
	ntd.Options(gonumbers.Separator("."), gonumbers.Delimiter(","))
	ntd_final := ntd.Perform(1234567890.506)
	ftm.Println(ntd_final) // "1,234,567,890.506"

	nth := gonumbers.NewNumberToHuman()
	nth.Options(gonumbers.Separator(","), gonumbers.Precision(2))
	nth_final := nth.Perform(489939)
	ftm.Println(nth_final) // "489,9 Thousand"
	
	ntp := gonumbers.NewNumberToPercentage()
	ntp.Options(gonumbers.Separator(","), gonumbers.Delimiter("."))
	ntp_final := ntp.Perform(1000)
	ftm.Println(ntp_final) // "1.000,000%"

	ntph := gonumbers.NewNumberToPhone()
	ntph.Options(gonumbers.CountryCode("1"))
	ntph_final, err := ntph.Perform(1235551234)
	if err != nil{
		fmt.Println(err)
	}
	ftm.Println(ntph_final) // "+1-123-555-1234"

	nths := gonumbers.NewNumberToHumanSize()
	nths_final, err := nths.Perform(1235551234)
	if err != nil{
		fmt.Println(err)
	}
	ftm.Println(nths_final) // "1.15 GB"
}
```

##Number to Currency:

**Options:**

*	**gonumbers.Unit(string):** 
*	**gonumbers.Precision(int):** 
*	**gonumbers.Separator(string):** 
*	**gonumbers.Delimiter(string):** 

```go
ntc2 := gonumbers.NumberToCurrency)
ntc2.Perform(1234567890.50) // "$1,234,567,890.50"

ntc3 := gonumbers.NewNumberToCurrency()
ntc3.Options(gonumbers.Precision(2))
ntc3.Perform(1234567890.506) // "$1,234,567,890.51"

ntc4 := gonumbers.NewNumberToCurrency()
ntc4.Options(gonumbers.Precision(2), gonumbers.Unit("$"), gonumbers.Separator("."))
ntc4.Perform(1234567890) // "$1,234,567,890.00"

ntc5 := gonumbers.NewNumberToCurrency()
ntc5.Options(gonumbers.Precision(3), gonumbers.Unit("CAD$"), gonumbers.Separator("."), gonumbers.Delimiter(","))
ntc5.Perform(1234567890.506) // "CAD$1,234,567,890.506"

ntc6 := gonumbers.NewNumberToCurrency()
ntc6.Options(gonumbers.Precision(2), gonumbers.Separator("."), gonumbers.Delimiter(""))
ntc6.Perform(1234567890.50) // "$1234567890.50"

ntc7 := gonumbers.NewNumberToCurrency()
ntc7.Options(gonumbers.Precision(2), gonumbers.Separator(","), gonumbers.Delimiter("."))
ntc7.Perform(1234567890.506) // "$1.234.567.890,51"
```

##Number to Delimiter:

**Options:**

*	**gonumbers.Delimiter(string):** 
*	**gonumbers.Separator(string):** 

```go
ntd1 := gonumbers.NewNumberToDelimiter()
ntd1.Perform(1234567890) // "1,234,567,890"

ntd2 := gonumbers.NewNumberToDelimiter()
ntd2.Options(gonumbers.Separator("."), gonumbers.Delimiter(","))
ntd2.Perform(1234567890.506) // "1,234,567,890.506"

ntd3 := gonumbers.NewNumberToDelimiter()
ntd3.Options(gonumbers.Separator("-"), gonumbers.Delimiter(":"))
ntd3.Perform(1234567890.34) // "1:234:567:890-34"

ntd4 := gonumbers.NewNumberToDelimiter()
ntd4.Options(gonumbers.Separator("."), gonumbers.Delimiter(""))
ntd4.Perform(1234567890.34) // "1234567890.34"

ntd5 := gonumbers.NewNumberToDelimiter()
ntd5.Options(gonumbers.Separator(""), gonumbers.Delimiter(""))
ntd5.Perform(1234567890.34) // "123456789034"
```

##Number to Human:

**Options:**

*	**gonumbers.Precision(int):** 
*	**gonumbers.Separator(string):** 

```go
nth1 := gonumbers.NewNumberToHuman()
nth1.Perform(123) // "123"

nth2 := gonumbers.NewNumberToHuman()
nth2.Perform(1234) // "1.23 Thousand"

nth3 := gonumbers.NewNumberToHuman()
nth3.Perform(12345) // "12.3 Thousand"

nth4 := gonumbers.NewNumberToHuman()
nth4.Perform(1234567) // "1.23 Million"

nth5 := gonumbers.NewNumberToHuman()
nth5.Perform(1234567890) // "1.23 Billion"

nth6 := gonumbers.NewNumberToHuman()
nth6.Perform(1234567890123) // "1.23 Trillion"

nth7 := gonumbers.NewNumberToHuman()
nth7.Perform(1234567890123456) // "1.23 Quadrillion"

nth8 := gonumbers.NewNumberToHuman()
nth8.Options(gonumbers.Precision(2))
nth8.Perform(489939) // "490 Thousand"

nth9 := gonumbers.NewNumberToHuman()
nth9.Options(gonumbers.Precision(4))
nth9.Perform(489939) // "489.9 Thousand"

nth10 := gonumbers.NewNumberToHuman()
nth10.Options(gonumbers.Precision(4), gonumbers.Separator(","))
nth10.Perform(489939) // "489,9 Thousand"
```

##Number to Human Size:

**Options:**

*	**gonumbers.Delimiter(string):** Defines the delimiter for the number (default => ".")
*	**gonumbers.Precision(int):** Sets precision of the number (default => 3).
*	**gonumbers.Prefix(string):** Defines if it is binary or si (default => binary)

```go
nths := gonumbers.NewNumberToHumanSize()
nths.Perform(333) // "333 Bytes"

nths := gonumbers.NewNumberToHumanSize()
nths.Perform(1234) // "1.21 KB"

nths := gonumbers.NewNumberToHumanSize()
nths.Perform(1234567890) // "1.15 GB"

nths := gonumbers.NewNumberToHumanSize()
nths.Options(gonumbers.Precision(2))
nths.Perform(1234567) // "1.2 MB"

nths := gonumbers.NewNumberToHumanSize()
nths.Options(gonumbers.Precision(2), gonumbers.Separator(","))
nths.Perform(1234567) // "1,2 MB"

nths := gonumbers.NewNumberToHumanSize()
nths.Options(gonumbers.Precision(5))
nths.Perform(524288000) // "500 MB"
```

##Number to Percentage:

**Options:**

*	**gonumbers.Precision(int):** 
*	**gonumbers.Separator(string):** 
*  **gonumbers.Delimiter(string):**

```go
ntp := gonumbers.NewNumberToPercentage()
ntp.Perform(100) // "100.000%"

ntp := gonumbers.NewNumberToPercentage()
ntp.Perform("98") // "98.000%"

ntp := gonumbers.NewNumberToPercentage()
npt.Options(gonumbers.Precision(0))
ntp.Perform(100) // "100%"

ntp := gonumbers.NewNumberToPercentage()
ntp.Perform(1000) // "1,000.000%"

ntp := gonumbers.NewNumberToPercentage()
npt.Options(gonumbers.Delimiter("."), gonumbers.Separator(","))
ntp.Perform(1000) // "1.000,000%"

ntp := gonumbers.NewNumberToPercentage()
npt.Options(gonumbers.Precision(5))
ntp.Perform(302.24398923423) // "302.24399%"

ntp := gonumbers.NewNumberToPercentage()
ntp.Perform("hahahah") // "0.000%"
```

##Number to Phone:

**Options:**

*	**gonumbers.Delimiter(string):** Defines the delimiter for the number. Default is "-"
*	**gonumbers.AreaCode(bool):** Prints the area code in ()
*	**gonumbers.Extension(string):** Appends extension to phone number. Ex. +1(123) 555-6789 x 4545
*	**gonumbers.CountryCode(string):** Prepends country code to phone number (ex. +1)
*	**gonumbers.DigitsSize(int):** Sets the number of digits of the second part of the phone (digits_count:5). Ex. 1234555556789 => 1234-55555-6789

```go
ntph := gonumbers.NewNumberToPhone()
ntph.Perform(5551234) // "555-1234"

ntph := gonumbers.NewNumberToPhone()
ntph.Perform(1235551234) // "123-555-1234"

ntph := gonumbers.NewNumberToPhone()
ntph.Options(gonumbers.Areacode(true))
ntph.Perform(1235551234) // (123) 555-1234 

ntph := gonumbers.NewNumberToPhone()
ntph.Options(gonumbers.Areacode(true), gonumbers.CountryCode("1"))
ntph.Perform(1235551234) // "+1(123) 555-1234"

ntph := gonumbers.NewNumberToPhone()
ntph.Options(gonumbers.CountryCode("1"))
ntph.Perform(1235551234) // "+1-123-555-1234"

ntph := gonumbers.NewNumberToPhone()
ntph.Options(gonumbers.AreaCode(true), gonumbers.CountryCode("1"), gonumbers.Extension("4545"))
ntph.Perform(1235551234) // +1(123455) 555-6789 x 4545

ntph := gonumbers.NewNumberToPhone()
ntph.Options(gonumbers.AreaCode(true), gonumbers.CountryCode("1"), gonumbers.Extension("4545"), gonumbers.DigitsSize(5))
ntph.Perform(1235551234) // "+1(1234) 55555-6789 x 4545"

ntph := gonumbers.NewNumberToPhone()
ntph.Options(gonumbers.AreaCode(true), gonumbers.CountryCode("55"), gonumbers.Delimiter(","), gonumbers.DigitsSize(5))
ntph.Perform(1235551234) // "+55(1234) 55555,6789"
```

**This is a work in progress.**