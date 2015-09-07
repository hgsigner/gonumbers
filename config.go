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

type numberToDelimiterer interface {
	separatorer
	delimiterer
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

func Unit(u string) unitOption {
	return func(obj uniter) {
		obj.setUnit(u)
	}
}

func Precision(p int) precisionOption {
	return func(obj precisioner) {
		obj.setPrecision(p)
	}
}

func Separator(s string) separatorOption {
	return func(obj separatorer) {
		obj.setSeparator(s)
	}
}

func Delimiter(d string) delimiterOption {
	return func(obj delimiterer) {
		obj.setDelimiter(d)
	}
}

func Prefix(p string) prefixOption {
	return func(obj prefixarer) {
		obj.setPrefix(p)
	}
}

func AreaCode(ac bool) areaCodeOption {
	return func(obj areaCoder) {
		obj.setAreaCode(ac)
	}
}

func Extension(ex string) extensionOption {
	return func(obj extensioner) {
		obj.setExtension(ex)
	}
}

func CountryCode(cc string) countryCodeOption {
	return func(obj countryCoder) {
		obj.setCountryCode(cc)
	}
}

func DigitsSize(ds int) digitsSizeOption {
	return func(obj digitsSizer) {
		obj.setDigitsSize(ds)
	}
}
