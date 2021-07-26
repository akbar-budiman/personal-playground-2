package es

import "github.com/akbar-budiman/personal-playground-2/entity"

type EsClient interface {
	InsertData(esUser *entity.EsUser)
	FindDataBySearchKey(searchKey string)
}

type EsClientImpl struct {
}

func (esClient *EsClientImpl) InsertData(esUser *entity.EsUser) {
	InsertData(esUser)
}

func (esClient *EsClientImpl) FindDataBySearchKey(searchKey string) {
	FindDataBySearchKey(searchKey)
}
