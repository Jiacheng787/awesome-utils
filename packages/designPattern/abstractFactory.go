package designPattern

type ICache interface {
	Get(key string) string
	Set(key, value string)
}

type Redis map[string]string

func NewRedis() ICache {
	return make(Redis)
}

func (r Redis) Get(key string) string {
	if value, ok := r[key]; ok {
		return value
	}
	return ""
}

func (r Redis) Set(key, value string) {
	r[key] = value
}

type MemCache map[string]string

func NewMemCache() ICache {
	return make(MemCache)
}

func (m MemCache) Get(key string) string {
	if value, ok := m[key]; ok {
		return value
	}
	return ""
}

func (m MemCache) Set(key, value string) {
	m[key] = value
}
