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
	ntc := new(gonumbers.NumberToCurrency)
	ntc.Options(gonumbers.Precision(2))
	ntc_final := ntc.Perform(1234567890.506)
	ftm.Println(ntc_final) // "$1,234,567,890.51"

	ntd := new(gonumbers.NumberToDelimiter)
	ntd.Options(gonumbers.Separator("."), gonumbers.Delimiter(","))
	ntd_final := ntd.Perform(1234567890.506)
	ftm.Println(ntd_final) // "1,234,567,890.506"

	nth := new(gonumbers.NumberToHuman)
	nth.Options(gonumbers.Separator(","), gonumbers.Precision(2))
	nth_final := nth.Perform(489939)
	ftm.Println(nth_final) // "489,9 Thousand"
	
	ntp := new(gonumbers.NumberToPercentage)
	ntp.Options(gonumbers.Separator(","), gonumbers.Delimiter("."))
	ntp_final := ntp.Perform(1000)
	ftm.Println(ntp_final) // "1.000,000%"

	ntph := new(gonumbers.NumberToPhone)
	ntph.Options(gonumbers.CountryCode("1"))
	ntph_final, err := ntph.Perform(1235551234)
	if err != nil{
		fmt.Println(err)
	}
	ftm.Println(ntph_final) // "+1-123-555-1234"

	nths := new(gonumbers.NumberToHumanSize)
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
ntc2 := new(gonumbers.NumberToCurrency)
ntc2.Perform(1234567890.50) // "$1,234,567,890.50"

ntc3 := new(gonumbers.NumberToCurrency)
ntc3.Options(gonumbers.Precision(2))
ntc3.Perform(1234567890.506) // "$1,234,567,890.51"

ntc4 := new(gonumbers.NumberToCurrency)
ntc4.Options(gonumbers.Precision(2), gonumbers.Unit("$"), gonumbers.Separator("."))
ntc4.Perform(1234567890) // "$1,234,567,890.00"

ntc5 := new(gonumbers.NumberToCurrency)
ntc5.Options(gonumbers.Precision(3), gonumbers.Unit("CAD$"), gonumbers.Separator("."), gonumbers.Delimiter(","))
ntc5.Perform(1234567890.506) // "CAD$1,234,567,890.506"

ntc6 := new(gonumbers.NumberToCurrency)
ntc6.Options(gonumbers.Precision(2), gonumbers.Separator("."), gonumbers.Delimiter(""))
ntc6.Perform(1234567890.50) // "$1234567890.50"

ntc7 := new(gonumbers.NumberToCurrency)
ntc7.Options(gonumbers.Precision(2), gonumbers.Separator(","), gonumbers.Delimiter("."))
ntc7.Perform(1234567890.506) // "$1.234.567.890,51"
```

##Number to Delimiter:

**Options:**

*	**gonumbers.Delimiter(string):** 
*	**gonumbers.Separator(string):** 

```go
ntd1 := new(gonumbers.NumberToDelimiter)
ntd1.Perform(1234567890) // "1,234,567,890"

ntd2 := new(gonumbers.NumberToDelimiter)
ntd2.Options(gonumbers.Separator("."), gonumbers.Delimiter(","))
ntd2.Perform(1234567890.506) // "1,234,567,890.506"

ntd3 := new(gonumbers.NumberToDelimiter)
ntd3.Options(gonumbers.Separator("-"), gonumbers.Delimiter(":"))
ntd3.Perform(1234567890.34) // "1:234:567:890-34"

ntd4 := new(gonumbers.NumberToDelimiter)
ntd4.Options(gonumbers.Separator("."), gonumbers.Delimiter(""))
ntd4.Perform(1234567890.34) // "1234567890.34"

ntd5 := new(gonumbers.NumberToDelimiter)
ntd5.Options(gonumbers.Separator(""), gonumbers.Delimiter(""))
ntd5.Perform(1234567890.34) // "123456789034"
```

##Number to Human:

**Options:**

*	**gonumbers.Precision(int):** 
*	**gonumbers.Separator(string):** 

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

##Number to Human Size:

**Options:**

*	**gonumbers.Delimiter(string):** Defines the delimiter for the number (default => ".")
*	**gonumbers.Precision(int):** Sets precision of the number (default => 3).
*	**gonumbers.Prefix(string):** Defines if it is binary or si (default => binary)

```go
gonumbers.NumberToHumanSize(333) // "333 Bytes"

gonumbers.NumberToHumanSize(1234) // "1.21 KB"

gonumbers.NumberToHumanSize(1234567890) // "1.15 GB"

gonumbers.NumberToHumanSize(1234567, "precision:2") // "1.2 MB"

gonumbers.NumberToHumanSize(1234567, "separator:,", "precision:2") // "1,2 MB"

gonumbers.NumberToHumanSize(524288000, "precision:5") // "500 MB"
```

##Number to Percentage:

**Options:**

*	**gonumbers.Precision(int):** 
*	**gonumbers.Separator(string):** 
*  **gonumbers.Delimiter(string):**

```go
gonumbers.NumberToPercentage(100) // "100.000%"

gonumbers.NumberToPercentage("98") // "98.000%"

gonumbers.NumberToPercentage(100, "precision:0") // "100%"

gonumbers.NumberToPercentage(1000) // "1,000.000%"

gonumbers.NumberToPercentage(1000, "delimiter:.", "separator:,") // "1.000,000%"

gonumbers.NumberToPercentage(302.24398923423, "precision:5") // "302.24399%"

gonumbers.NumberToPercentage("hahahah") // "0.000%"
```

##Number to Phone:

**Options:**

*	**gonumbers.Delimiter(string):** Defines the delimiter for the number. Default is "-"
*	**gonumbers.AreaCode(bool):** Prints the area code in ()
*	**gonumbers.Extension(string):** Appends extension to phone number. Ex. +1(123) 555-6789 x 4545
*	**gonumbers.CountryCode(string):** Prepends country code to phone number (ex. +1)
*	**gonumbers.DigitsSize(int):** Sets the number of digits of the second part of the phone (digits_count:5). Ex. 1234555556789 => 1234-55555-6789

```go
gonumbers.NumberToPhone(5551234) // "555-1234"

gonumbers.NumberToPhone(1235551234) // "123-555-1234"

gonumbers.NumberToPhone(1235551234, "area_code:true") // (123) 555-1234

gonumbers.NumberToPhone(1235551234, "area_code:true", "country_code:1") // "+1(123) 555-1234"

gonumbers.NumberToPhone(1235551234, "country_code:1") // "+1-123-555-1234"

gonumbers.NumberToPhone(1234555556789, "area_code:true", "country_code:1", "extension:4545") // +1(123455) 555-6789 x 4545

gonumbers.NumberToPhone(1234555556789, "area_code:true", "country_code:1", "extension:4545", "digits_size:5") // "+1(1234) 55555-6789 x 4545"

gonumbers.NumberToPhone(1234555556789, "area_code:true", "country_code:55", "digits_size:5", "delimiter:,") // "+55(1234) 55555,6789"
```

**This is a work in progress.**