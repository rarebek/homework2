package repo

type RedisRepositoryStorage interface {
	Set(key, value string) error
	SetWithTTL(key, value string, seconds int64) error
	Get(key string) (interface{}, error)
}
