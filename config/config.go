package config

var configMap = make(map[string]interface{})

func init() {
	configMap["port"] = "8080"
}

func GetConfig(key string, defaultValue interface{}) interface{} {
	value, ok := configMap[key]
	if !ok {
		return defaultValue
	}
	return value
}
func SetConfig(key string, value interface{}) {
	configMap[key] = value
}
