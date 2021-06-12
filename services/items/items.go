package items_service

import (
	items_domain "bookstore_items-api/domain/items"
	queries_domain "bookstore_items-api/domain/queries"

	errors_utils "github.com/bRRRITSCOLD/bookstore_utils-go/errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items_domain.Item) (*items_domain.Item, errors_utils.APIError)
	Get(string) (*items_domain.Item, errors_utils.APIError)
	Search(queries_domain.EsQuery) ([]items_domain.Item, errors_utils.APIError)
}

type itemsService struct{}

func (s *itemsService) Create(item items_domain.Item) (*items_domain.Item, errors_utils.APIError) {
	if saveErr := item.Save(); saveErr != nil {
		return &item, saveErr
	}

	return &item, nil
}

func (s *itemsService) Get(itemId string) (*items_domain.Item, errors_utils.APIError) {
	item := items_domain.Item{
		ItemID: itemId,
	}

	if saveErr := item.Get(); saveErr != nil {
		return &item, saveErr
	}

	return &item, nil
}

func (s *itemsService) Search(query queries_domain.EsQuery) ([]items_domain.Item, errors_utils.APIError) {
	return nil, errors_utils.NewInternalServerAPIError("implement me", nil)
}
