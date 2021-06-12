package elasticsearch_client

import (
	// "os"
	"context"
	"fmt"
	"time"

	logger_utils "github.com/bRRRITSCOLD/bookstore_utils-go/logger"
	"github.com/olivere/elastic"
)

const (
	envEsHosts = "ELASTIC_HOSTS"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(index string, docType string, id string) (*elastic.GetResult, error)
	// Get(string, string, string) (*elastic.GetResult, error)
	// Search(string, elastic.Query) (*elastic.SearchResult, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	log := logger_utils.GetLogger()

	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)
	if err != nil {
		panic(err)
	}

	Client.setClient(client)

}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index(index).
		Type(docType).
		BodyJson(doc).
		Do(ctx)

	if err != nil {
		logger_utils.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}

	return result, nil
}

func (c *esClient) Get(index string, docType string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().
		Index(index).
		Type(docType).
		Id(id).
		Do(ctx)
	if err != nil {
		logger_utils.Error(fmt.Sprintf("error when trying to get id %s", id), err)
		return nil, err
	}
	return result, nil
}

// func (c *esClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
// 	ctx := context.Background()
// 	result, err := c.client.Search(index).
// 		Query(query).
// 		RestTotalHitsAsInt(true).
// 		Do(ctx)
// 	if err != nil {
// 		logger_utils.Error(fmt.Sprintf("error when trying to search documents in index %s", index), err)
// 		return nil, err
// 	}
// 	return result, nil
// }
