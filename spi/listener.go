package spi

type Listener interface {
	Before() error
	After() error
	Success() error
	Failure() error
}
