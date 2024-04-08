package lifecycle

type AfterObjectInject interface {
	AfterObjectInjectAction() error
}
