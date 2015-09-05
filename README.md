#GoNumbers [![Build Status](https://travis-ci.org/hgsigner/gonumbers.svg?branch=master)](https://travis-ci.org/hgsigner/gonumbers)

GoNumbers is an implementation of numbers helpers for Go-lang. This project is inspired by the ruby library ActiveSupport::NumberHelper.

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
	ntc.Options(ntc.Precision(2))
	ntc_final := ntc.Perform(1234567890.506)
	ftm.Println(ntc_final) // "$1,234,567,890.51"

	ntd := new(gonumbers.NumberToDelimiter)
	ntd.Options(ntd.Separator("."), ntd.Delimiter(","))
	ntd_final := ntd.Perform(1234567890.506)
	ftm.Println(ntd_final) // "1,234,567,890.506"

	nth := new(gonumbers.NumberToHuman)
	nth.Options(nth.Separator(","), nth.Precision(2))
	nth_final := nth.Perform(489939)
	ftm.Println(nth_final) // "489,9 Thousand"
	
	ntp := new(gonumbers.NumberToPercentage)
	ntp.Options(ntd.Separator(","), ntd.Delimiter("."))
	ntp_final := ntp.Perform(1000)
	ftm.Println(ntp_final) // "1.000,000%"

	ntph := new(gonumbers.NumberToPhone)
	ntph.Options(ntd.CountryCode("1"))
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

*	**unit:** 
*	**precision:** 
*	**separator:** 
*	**delimiter:** 

```go
ntc2 := new(NumberToCurrency)
ntc2.Perform(1234567890.50) // "$1,234,567,890.50"

ntc3 := new(NumberToCurrency)
ntc3.Options(ntc3.Precision(2))
ntc3.Perform(1234567890.506) // "$1,234,567,890.51"

ntc4 := new(NumberToCurrency)
ntc4.Options(ntc4.Precision(2), ntc4.Unit("$"), ntc4.Separator("."))
ntc4.Perform(1234567890) // "$1,234,567,890.00"

ntc5 := new(NumberToCurrency)
ntc5.Options(ntc5.Precision(3), ntc5.Unit("CAD$"), ntc5.Separator("."), ntc5.Delimiter(","))
ntc5.Perform(1234567890.506) // "CAD$1,234,567,890.506"

ntc6 := new(NumberToCurrency)
ntc6.Options(ntc6.Precision(2), ntc5.Separator("."), ntc5.Delimiter(""))
ntc6.Perform(1234567890.50) // "$1234567890.50"

ntc7 := new(NumberToCurrency)
ntc7.Options(ntc7.Precision(2), ntc7.Separator(","), ntc7.Delimiter("."))
ntc7.Perform(1234567890.506) // "$1.234.567.890,51"
```

##Number to Delimiter:

**Options:**

*	**separator:** 
*	**delimiter:** 

```go
ntd1 := new(NumberToDelimiter)
ntd1.Perform(1234567890) // "1,234,567,890"

ntd2 := new(NumberToDelimiter)
ntd2.Options(ntd2.Separator("."), ntd2.Delimiter(","))
ntd2.Perform(1234567890.506) // "1,234,567,890.506"

ntd3 := new(NumberToDelimiter)
ntd3.Options(ntd3.Separator("-"), ntd3.Delimiter(":"))
ntd3.Perform(1234567890.34) // "1:234:567:890-34"

ntd4 := new(NumberToDelimiter)
ntd4.Options(ntd4.Separator("."), ntd4.Delimiter(""))
ntd4.Perform(1234567890.34) // "1234567890.34"

ntd5 := new(NumberToDelimiter)
ntd5.Options(ntd5.Separator(""), ntd5.Delimiter(""))
ntd5.Perform(1234567890.34) // "123456789034"
```

##Number to Human:

**Options:**

*	**separator:** 
*	**precision:** 

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

*	**delimiter:** Defines the delimiter for the number (default => ".")
*	**precision:** Sets precision of the number (default => 3).

*	**prefix:** Defines if it is binary or si (default => binary)

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

* **separator:**
* **precision:**
* **delimiter:**

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

*	**delimiter:** Defines the delimiter for the number. Default is "-"
*	**area_code:** Prints the area code in ()
*	**extension:** Appends extension to phone number. Ex. +1(123) 555-6789 x 4545
*	**country_code:** Prepends country code to phone number (ex. +1)
*	**digits_size:** Sets the number of digits of the second part of the phone (digits_count:5). Ex. 1234555556789 => 1234-55555-6789

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