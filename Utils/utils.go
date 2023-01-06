package Utils

func KeyExists(m map[any]any, key any) (exists bool) {
	_, exists = m[key]

	return
}

func Contains(arr []string, s string) bool {
	for _, v := range arr {
		if s == v {
			return true
		}
	}
	return false
}
