package format

type Format interface {
	// ToJSON converts the given object to a JSON string.
	ToJSON(obj interface{}) (string, error)
	// PrettyJSON converts the given object to a pretty-printed JSON string.
	PrettyJSON(obj interface{}) (string, error)
}
