package fak


func SetIfNotExists[K comparable, V any](m map[K]V, k K, v V) bool {
	if m == nil {
		m = map[K]V{k:v}
	}else{
		if _, exists := m[k]; exists {
			return false
		}
		m[k] = v
	}
	return true
}
