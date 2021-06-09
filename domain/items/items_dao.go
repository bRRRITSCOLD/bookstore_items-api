package items_domain

import (
	elasticsearch_client "bookstore_items-api/clients/elasticsearch"

	errors_utils "github.com/bRRRITSCOLD/bookstore_utils-go/errors"
)

const (
	ITEMS_ELASTICSEARCH_INDEX   = "items"
	ITEM_ELASTICSEARCH_DOC_TYPE = "item"
)

func (i *Item) Save() *errors_utils.APIError {
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
