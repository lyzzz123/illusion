package lifecycle

type AfterObjectDestroy interface {
	AfterObjectDestroyAction() error
}
