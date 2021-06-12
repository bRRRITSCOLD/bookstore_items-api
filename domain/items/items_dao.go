package items_domain

import (
	elasticsearch_client "bookstore_items-api/clients/elasticsearch"
	queries_domain "bookstore_items-api/domain/queries"
	"encoding/json"
	"fmt"
	"strings"

	errors_utils "github.com/bRRRITSCOLD/bookstore_utils-go/errors"
)

const (
	ITEMS_ELASTICSEARCH_INDEX   = "items"
	ITEM_ELASTICSEARCH_DOC_TYPE = "item"
)

func (i *Item) Save() errors_utils.APIError {
	result, err := elasticsearch_client.Client.Index(
		ITEMS_ELASTICSEARCH_INDEX,
		ITEM_ELASTICSEARCH_DOC_TYPE,
		i,
	)
	if err != nil {
		return errors_utils.NewInternalServerAPIError("error when trying to save item", err)
	}

	i.ItemID = result.Id

	return nil
}

func (i *Item) Get() errors_utils.APIError {
	itemId := i.ItemID

	getResult, getErr := elasticsearch_client.Client.Get(
		ITEMS_ELASTICSEARCH_INDEX,
		ITEM_ELASTICSEARCH_DOC_TYPE,
		i.ItemID,
	)
	if getErr != nil {
		if strings.Contains(getErr.Error(), "404") {
			return errors_utils.NewNotFoundAPIError(fmt.Sprintf("no item with id %s found", i.ItemID), nil)
		}
		return errors_utils.NewInternalServerAPIError(fmt.Sprintf("error when trying to get item with id %s", i.ItemID), getErr)
	}

	bytes, marshalJsonErr := getResult.Source.MarshalJSON()
	if marshalJsonErr != nil {
		return errors_utils.NewInternalServerAPIError("error when trying to marshal database result", marshalJsonErr)
	}

	if unmarshalErr := json.Unmarshal(bytes, &i); unmarshalErr != nil {
		return errors_utils.NewInternalServerAPIError("error when trying to unmarshal database result", unmarshalErr)
	}

	i.ItemID = itemId

	return nil
}

func (i *Item) Search(query queries_domain.EsQuery) ([]Item, errors_utils.APIError) {
	result, err := elasticsearch_client.Client.Search(ITEMS_ELASTICSEARCH_INDEX, query.Build())
	if err != nil {
		return nil, errors_utils.NewInternalServerAPIError("error when trying to search documents", err)
	}

	items := make([]Item, result.TotalHits())
	for i, hit := range result.Hits.Hits {
		bytes, marshalJsonErr := hit.Source.MarshalJSON()
		if marshalJsonErr != nil {
			return nil, errors_utils.NewInternalServerAPIError("error when trying to marshal json", marshalJsonErr)
		}

		var item Item
		if unmarshalErr := json.Unmarshal(bytes, &item); unmarshalErr != nil {
			return nil, errors_utils.NewInternalServerAPIError("error when trying to unmarshal json", unmarshalErr)
		}
		item.ItemID = hit.Id
		items[i] = item
	}

	if len(items) == 0 {
		return nil, errors_utils.NewNotFoundAPIError("no items found matching given criteria", nil)
	}

	return items, nil
}
