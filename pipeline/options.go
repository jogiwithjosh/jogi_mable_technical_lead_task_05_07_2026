package pipeline

type Option[T any] func(*Pipeline[T])

func WithStage[T any](
	stage Stage[T],
) Option[T] {
	return func(
		p *Pipeline[T],
	) {
		p.stages = append(
			p.stages,
			stage,
		)
	}
}

func WithConfig[T any](config Config) Option[T] {
	return func(
		p *Pipeline[T],
	) {
		p.config = config
	}
}

func (p *Pipeline[T]) WithObserver(
	o Observer,
) *Pipeline[T] {
	p.observer = o

	return p
}
