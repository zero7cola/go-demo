package cahce

type Store interface {
	Get(key string, clear bool) string

	Set(key string, value string) bool
}
