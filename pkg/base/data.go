// most common data structure
package atomgo

type Name struct {
	first  string
	last   string
	middle string
}

type DayOfBirth struct {
	year  int
	month int
	day   int
}

type Phone struct {
	landline string
	mobile   string
}

type Address struct {
	country  string
	state    string
	city     string
	street   string
	zipcode  string
	timezone int
	ipv4     string
}

type Avatar struct {
	category string
	ethnic   string
	file     string
	url      string
}
