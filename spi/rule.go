package spi

type Rule interface {
	When(conditions ...Condition) (bool, error)
	Then() error
}
