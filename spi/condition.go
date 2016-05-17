package spi

type Condition interface {
	eval() (bool, error)
}
