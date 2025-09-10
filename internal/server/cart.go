package server

import (
	web "buckingham_bakery/cmd/web/templates/components"
	"buckingham_bakery/internal/dto"
	"errors"
	"log"
	"log/slog"
	"net/http"
)

func GetSideCartOrders(w http.ResponseWriter, r *http.Request) {
	cartOrders, _ := getCartFromSession(r)

	component := web.SideCartOrderList(cartOrders)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering Cart: %e", err)
	}
}

func PutCartWebHandler(cartId int, w http.ResponseWriter, r *http.Request) {

	cartOrders, _ := getCartFromSession(r)

	orderToAdd, err := findOrderBy(cartId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	cartOrders.AddToCart(*orderToAdd)
	saveCartToSession(w, r, cartOrders)

	component := web.NavCart()

	HtmxTrigger("update-sidecart", w)

	// TODO: This is not needed now.
	err = component.Render(r.Context(), w)
	if err != nil {
		slog.Error("Error rendering Cart", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}

func findOrderBy(id int) (*dto.FoodOrder, error) {
	storedOrders := &orders

	for _, order := range *storedOrders {
		if order.Id == id {
			return &order, nil
		}
	}

	return nil, errors.New("Id not found")
}
