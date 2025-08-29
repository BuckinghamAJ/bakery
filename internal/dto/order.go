package dto

import (
	"fmt"
)

type FoodOrder struct {
	Id          int
	Name        string
	Cost        float64
	ImgPath     string
	Description string
}

func (fo FoodOrder) FormatCost() string {
	return fmt.Sprintf("$%.2f", fo.Cost)
}

func (fo FoodOrder) doesIdMatch(id int) bool {
	return id == fo.Id
}

type OrderInCart struct {
	Orders    []FoodOrder
	TotalCost float64
}

func (oic *OrderInCart) AddToCart(order FoodOrder) {
	oic.Orders = append(oic.Orders, order)
	oic.TotalCost += order.Cost
}

func (oic OrderInCart) calculateTotalCost() float64 {
	var total float64 = 0
	for _, order := range oic.Orders {
		total += order.Cost
	}
	return total
}

func (oic OrderInCart) DisplayTotalCost() string {
	return fmt.Sprintf("$%.2f", oic.calculateTotalCost())
}

func (oic OrderInCart) Condensed() *ViewCartOrders {
	// Does it make more sense to do a hash?
	condensedOrders := &ViewCartOrders{
		orders: map[int]CondensedOrder{},
	}

	for _, order := range oic.Orders {
		if entry, ok := condensedOrders.orders[order.Id]; ok {
			entry.amount += 1
		}

		condensedOrders.orders[order.Id] = CondensedOrder{
			amount: 1,
			order:  order,
		}
	}

	return condensedOrders
}

type CondensedOrder struct {
	amount int
	order  FoodOrder
}

type ViewCartOrders struct {
	orders map[int]CondensedOrder
}
