package server

import (
	"buckingham_bakery/internal/dto"
	"encoding/gob"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

var SESSION_NAME = "cart-session"
var CART_KEY = "cart"

func init() {
	gob.Register(dto.CondensedOrder{})
	gob.Register(dto.ViewCartOrders{})
}

func saveCartToSession(w http.ResponseWriter, r *http.Request, cart *dto.ViewCartOrders) error {
	session, err := store.Get(r, SESSION_NAME)
	if err != nil {
		return err
	}

	session.Values[CART_KEY] = cart
	return session.Save(r, w)
}

func getCartFromSession(r *http.Request) (*dto.ViewCartOrders, error) {
	session, err := store.Get(r, SESSION_NAME)

	if err != nil {
		return dto.NewCartOrders(), err
	}

	if cart, ok := session.Values[CART_KEY]; ok {
		if c, ok := cart.(dto.ViewCartOrders); ok {
			return &c, nil
		}
	}

	return dto.NewCartOrders(), nil

}

func clearCartFromSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, SESSION_NAME)
	if err != nil {
		return err
	}

	session.Values[CART_KEY] = nil
	return session.Save(r, w)
}
