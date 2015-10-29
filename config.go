package gonumbers

type uniter interface {
	setUnit(string)
}

type precisioner interface {
	setPrecision(int)
}

type separatorer interface {
	setSeparator(string)
}

type delimiterer interface {
	setDelimiter(string)
}

type prefixarer interface {
	setPrefix(string)
}

type areaCoder interface {
	setAreaCode(bool)
}

type extensioner interface {
	setExtension(string)
}

type countryCoder interface {
	setCountryCode(string)
}

type digitsSizer interface {
	setDigitsSize(int)
}

// Custom Types

type unitOption func(uniter)
type precisionOption func(precisioner)
type separatorOption func(separatorer)
type delimiterOption func(delimiterer)
type prefixOption func(prefixarer)
type areaCodeOption func(areaCoder)
type extensionOption func(extensioner)
type countryCodeOption func(countryCoder)
type digitsSizeOption func(digitsSizer)

// Functions for api

//It sets the unit for NumberToCurrency function
func Unit(u string) unitOption {
	return func(obj uniter) {
		obj.setUnit(u)
	}
}

// Sets precision of the number.
func Precision(p int) precisionOption {
	return func(obj precisioner) {
		obj.setPrecision(p)
	}
}

// Defines the separator for the number.
func Separator(s string) separatorOption {
	return func(obj separatorer) {
		obj.setSeparator(s)
	}
}

// Defines the delimiter for the number.
func Delimiter(d string) delimiterOption {
	return func(obj delimiterer) {
		obj.setDelimiter(d)
	}
}

// Defines if it prefix for converting numbers to human size.
//It could be either binary or si (default => binary).
func Prefix(p string) prefixOption {
	return func(obj prefixarer) {
		obj.setPrefix(p)
	}
}

// It wraps the area code parenthesis. E.g. (123) 444-5678
func AreaCode(ac bool) areaCodeOption {
	return func(obj areaCoder) {
		obj.setAreaCode(ac)
	}
}

// It appends the extension to phone numbers. E.g. ...-6789 x 4545
func Extension(ex string) extensionOption {
	return func(obj extensioner) {
		obj.setExtension(ex)
	}
}

// It adds the country code to phone numbers. E.g +1(123)...
func CountryCode(cc string) countryCodeOption {
	return func(obj countryCoder) {
		obj.setCountryCode(cc)
	}
}

/*
It sets the number of digits for the second part of a phone number. e.g 1233334567 => gonumbers.DigitsSize(4) / 12-3333-4567 / gonumbers.DigitsSize(2) -> 1233-33-4567
*/
func DigitsSize(ds int) digitsSizeOption {
	return func(obj digitsSizer) {
		obj.setDigitsSize(ds)
	}
}

//Structs Inits

// It initializes a new NumberToCurrency to permorf on.
func ToCurrency() *NumberToCurrency {
	return &NumberToCurrency{
		unit:      "$",
		separator: ".",
		precision: 2,
		delimiter: ",",
	}

}

// It initializes a new NumberToDelimiter to permorf on.
func ToDelimiter() *NumberToDelimiter {
	return &NumberToDelimiter{
		separator: ".",
		delimiter: ",",
	}
}

// It initializes a new NumberToHumanSize to permorf on.
func ToHumanSize() *NumberToHumanSize {
	return &NumberToHumanSize{
		separator: ".",
		precision: 3,
		delimiter: "",
		prefix:    "binary",
	}
}

// It initializes a new NumberToHuman to permorf on.
func ToHuman() *NumberToHuman {
	return &NumberToHuman{
		separator: ".",
		precision: 3,
	}
}

// It initializes a new NumberToPercentage to permorf on.
func ToPercentage() *NumberToPercentage {
	return &NumberToPercentage{
		separator: ".",
		precision: 3,
		delimiter: ",",
	}
}

// It initializes a new NumberToPhone to permorf on.
func ToPhone() *NumberToPhone {
	return &NumberToPhone{
		delimiter:   "-",
		areaCode:    false,
		extension:   "",
		countryCode: "",
		digitsSize:  3,
	}
}
