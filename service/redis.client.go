package service

type RedisClient interface {
	SetValue(key string, value []byte) error
	GetValue(key string) ([]byte, error)
}

type RedisClientImpl struct {
}

func (redisClient *RedisClientImpl) SetValue(key string, value []byte) error {
	return SetValue(key, value)
}

func (redisClient *RedisClientImpl) GetValue(key string) ([]byte, error) {
	return GetValue(key)
}
