package Utils

func KeyExists(m map[any]any, key any) (exists bool) {
	_, exists = m[key]

	return
}
