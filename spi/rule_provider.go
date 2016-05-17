package spi

type RuleProvider interface {
	Add(r Rule)
	Next() Rule
}
