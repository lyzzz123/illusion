package lifecycle

type AfterInitProperty interface {
	AfterInitPropertyAction(propertiesArray []map[string]string) error
}
