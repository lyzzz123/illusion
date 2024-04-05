package lifecycle

type BeforeInitProperty interface {
	BeforeInitPropertyAction() error
}
