package main

import (
	"./ctx"
)

type Rule interface {
	When() (bool, error)
	Then() error
}

type Engine interface {
	AddListener(l Listener)
	AddRule(r Rule)
	Run(in ctx.In) (ctx.Out, error)
}

type Listener interface {
	Before() error
	After() error
	Success() error
	Failure() error
}
