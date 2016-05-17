package spi

type Engine interface {
	AddListener(l listener.Listener)
	SetProvider(rp rule.RuleProvider)
	Run(in context.In) (context.Out, error)
}
