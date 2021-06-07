package items_controllers

import (
	items_domain "bookstore_items-api/domain/items"
	items_service "bookstore_items-api/services/items"
	"fmt"
	"net/http"

	"github.com/bRRRITSCOLD/bookstore_oauth-go/oauth"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	CreateItem(w http.ResponseWriter, r *http.Request)
	GetItem(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (iC *itemsController) CreateItem(w http.ResponseWriter, r *http.Request) {
	authenticateUserRequestErr := oauth.AuthenticateRequest(r)
	if authenticateUserRequestErr != nil {
		// TODO: return error to caller
		return
	}

	item := items_domain.Item{
		Seller: oauth.GetCallerID(r),
	}

	result, err := items_service.ItemsService.Create(item)
	if err != nil {
		// TODO: return error to caller
		return
	}

	// TODO: return result to caller
	fmt.Println(result)
}

func (iC *itemsController) GetItem(w http.ResponseWriter, r *http.Request) {

}
