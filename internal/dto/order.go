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

type CondensedOrder struct {
	Amount int
	Order  FoodOrder
}

func (co CondensedOrder) totalSingleOrderCost() float64 {
	return float64(co.Amount) * co.Order.Cost
}

func (co CondensedOrder) FormatSingleOrderCost() string {
	return fmt.Sprintf("$%.2f", co.totalSingleOrderCost())
}

type ViewCartOrders struct {
	Orders   map[int]CondensedOrder
	Quantity int
}

func NewCartOrders() *ViewCartOrders {
	return &ViewCartOrders{Orders: map[int]CondensedOrder{}, Quantity: 0}
}

func (vco ViewCartOrders) CalculateTotalCost() float64 {
	var total float64 = 0
	for _, order := range vco.Orders {
		total += order.totalSingleOrderCost()
	}
	return total
}

func (vco *ViewCartOrders) AddToCart(order FoodOrder) {
	vco.Quantity += 1

	if entry, ok := vco.Orders[order.Id]; ok {
		entry.Amount += 1
		vco.Orders[order.Id] = entry
	} else {
		vco.Orders[order.Id] = CondensedOrder{
			Amount: 1,
			Order:  order,
		}
	}
}

func (vco ViewCartOrders) FormatTotalOrderCost() string {
	return fmt.Sprintf("$%.2f", vco.CalculateTotalCost())
}

var CondensedOrders = ViewCartOrders{
	Orders: map[int]CondensedOrder{},
}

type OrderInCart struct {
	Orders    []FoodOrder
	TotalCost float64
}

func (oic *OrderInCart) AddToCart(order FoodOrder) {
	oic.Orders = append(oic.Orders, order)
	oic.TotalCost += order.Cost

	if entry, ok := CondensedOrders.Orders[order.Id]; ok {
		entry.Amount += 1
		CondensedOrders.Orders[order.Id] = entry
	} else {
		CondensedOrders.Orders[order.Id] = CondensedOrder{
			Amount: 1,
			Order:  order,
		}
	}

	// fmt.Printf("Condensed Orders: %+v\n", CondensedOrders)

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

func (oic OrderInCart) Condensed() ViewCartOrders {
	for _, order := range oic.Orders {
		if entry, ok := CondensedOrders.Orders[order.Id]; ok {
			entry.Amount += 1
		}

		CondensedOrders.Orders[order.Id] = CondensedOrder{
			Amount: 1,
			Order:  order,
		}
	}

	return CondensedOrders
}
