package lifecycle

type BeforeContainerInitProperty interface {
	BeforeContainerInitPropertyAction() error
}
