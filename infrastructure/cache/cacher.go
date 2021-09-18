package cache

type Cacher interface {
	Get(key string) interface{}
	Put(key string, val interface{}, ttl uint64) error
	IsExist(key string) bool
	Delete(key string) error
	Flush() error
}
