package lifecycle

type AfterContainerInitProperty interface {
	AfterContainerInitPropertyAction(propertiesArray []map[string]string) error
}
