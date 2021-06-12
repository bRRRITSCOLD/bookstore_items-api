package app

import (
	items_controllers "bookstore_items-api/controllers/items"
	ping_controllers "bookstore_items-api/controllers/ping"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/ping", ping_controllers.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items", items_controllers.ItemsController.CreateItem).Methods(http.MethodPost)
	router.HandleFunc("/items/{itemId}", items_controllers.ItemsController.GetItem).Methods(http.MethodGet)
	router.HandleFunc("/items/search", items_controllers.ItemsController.SearchItems).Methods(http.MethodPost)
}
