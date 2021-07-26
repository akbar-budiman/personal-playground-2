package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/akbar-budiman/personal-playground-2/es"
	"github.com/akbar-budiman/personal-playground-2/service"
)

type Resolver struct {
	Es    es.EsClient
	Redis service.RedisClient
	Crdb  service.CrdbClient
}
