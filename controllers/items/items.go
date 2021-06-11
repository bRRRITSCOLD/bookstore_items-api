package items_controllers

import (
	items_domain "bookstore_items-api/domain/items"
	items_service "bookstore_items-api/services/items"
	http_utils "bookstore_items-api/utils/http"
	"encoding/json"
	"io/ioutil"

	"net/http"

	"github.com/bRRRITSCOLD/bookstore_oauth-go/oauth"
	errors_utils "github.com/bRRRITSCOLD/bookstore_utils-go/errors"
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
		http_utils.RespondError(w, *authenticateUserRequestErr)
		return
	}

	callerId := oauth.GetCallerID(r)
	if callerId == 0 {
		apiErr := errors_utils.NewUnauthorizedAPIError("unauthorized for action", nil)
		http_utils.RespondError(w, apiErr)
		return
	}

	requestBody, readAllErr := ioutil.ReadAll(r.Body)
	if readAllErr != nil {
		apiErr := errors_utils.NewBadRequestAPIError("invalid request body", readAllErr)
		http_utils.RespondError(w, apiErr)
		return
	}
	defer r.Body.Close()

	var itemToCreate items_domain.Item
	if unmarshalErr := json.Unmarshal(requestBody, &itemToCreate); unmarshalErr != nil {
		apiErr := errors_utils.NewBadRequestAPIError("invalid item json body", unmarshalErr)
		http_utils.RespondError(w, apiErr)
		return
	}

	itemToCreate.Seller = callerId

	createdItem, createErr := items_service.ItemsService.Create(itemToCreate)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, createdItem)
}

func (iC *itemsController) GetItem(w http.ResponseWriter, r *http.Request) {

}
