package es

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/akbar-budiman/personal-playground-2/entity"
	"github.com/olivere/elastic/v7"
)

var (
	EsAddress                      = "http://127.0.0.1:9200"
	MyEsContext    context.Context = context.Background()
	MyEsClient     *elastic.Client
	EsIndexName    = "users"
	EsIndexSetting = `
	{
		"settings": {
		  "analysis": {
			"analyzer": {
			  "my_analyzer": {
				"tokenizer": "my_tokenizer"
			  }
			},
			"tokenizer": {
			  "my_tokenizer": {
				"type": "ngram",
				"token_chars": [
				  "letter"
				]
			  }
			}
		  }
		}
	  }
	`
)

func InitializeEsClient() {
	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}
	MyEsClient = client

	info, code, err := MyEsClient.Ping(EsAddress).Do(MyEsContext)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}

func InitializeIndex() {
	exists, err := MyEsClient.IndexExists(EsIndexName).Do(MyEsContext)
	if err != nil {
		panic(err)
	}
	if !exists {
		createIndex, err := MyEsClient.CreateIndex(EsIndexName).BodyString(EsIndexSetting).Do(MyEsContext)
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
			fmt.Println("Not acknowledged")
		}
	}
}

func InsertData(esUser *entity.EsUser) {
	_, err := MyEsClient.
		Index().
		Index(EsIndexName).
		Id(esUser.Name).
		BodyJson(esUser).
		Do(MyEsContext)
	if err != nil {
		panic(err)
	}
}

func FindDataBySearchKey(searchKey string) []*entity.EsUser {
	searchResult := findDataToEs(searchKey)
	response := convertToStruct(searchResult)
	return response
}

func findDataToEs(searchKey string) *elastic.SearchResult {
	searchSource := elastic.NewSearchSource()
	searchSource.Query(
		elastic.
			NewMatchQuery("searchable", searchKey).
			MinimumShouldMatch("100%"),
	)

	fmt.Println(searchSource.Source())

	searchService := MyEsClient.Search().Index(EsIndexName).SearchSource(searchSource)

	searchResult, err := searchService.Do(MyEsContext)
	if err != nil {
		panic(err)
	} else {
		return searchResult
	}
}

func convertToStruct(searchResult *elastic.SearchResult) []*entity.EsUser {
	var response = []*entity.EsUser{}
	for _, hit := range searchResult.Hits.Hits {
		var esUser entity.EsUser
		err := json.Unmarshal(hit.Source, &esUser)
		if err != nil {
			panic(err)
		}
		response = append(response, &esUser)
	}

	return response
}
