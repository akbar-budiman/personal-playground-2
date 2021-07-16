package main

var (
	redisAddress = "localhost:6379"
)

func main() {
	InitializeLocalRedisConnectionPool(redisAddress)

	InitializeLocalCrdbPool()
	InitializeLocalCrdbDdl()

	RegisterConsumer()

	RegisterUserRouter()
}
