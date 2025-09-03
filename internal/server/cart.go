package server

import (
	web "buckingham_bakery/cmd/web/templates/components"
	"buckingham_bakery/internal/dto"
	"errors"
	"log"
	"net/http"
)

func GetSideCartOrders(w http.ResponseWriter, r *http.Request) {
	component := web.SideCartOrderList()
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering Cart: %e", err)
	}
}

func PutCartWebHandler(cartId int, w http.ResponseWriter, r *http.Request) {

	cartOrders := &web.OrdersInCart

	orderToAdd, err := findOrderBy(cartId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	cartOrders.AddToCart(*orderToAdd)

	component := web.NavCart()

	HtmxTrigger("update-sidecart", w)

	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering Cart: %e", err)
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
