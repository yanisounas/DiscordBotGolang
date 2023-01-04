package Utils

func KeyExists(m map[string]string, key string) (exists bool) {
	_, exists = m[key]
	return
}
