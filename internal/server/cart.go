package server

import (
	web "buckingham_bakery/cmd/web/templates/components"
	"buckingham_bakery/internal/dto"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func PutCartWebHandler(cartId int, w http.ResponseWriter, r *http.Request) {

	fmt.Println("I am in PutCartWebHandler")
	// Need to figure out how to get existing number of card orders
	cartOrders := &web.OrdersInCart

	orderToAdd, err := findOrderBy(cartId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	cartOrders.AddToCart(*orderToAdd)

	// TODO: Add OrderBased on Order ID in Cart
	component := web.Cart()
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
