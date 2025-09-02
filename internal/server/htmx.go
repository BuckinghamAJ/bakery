package server

import (
	"net/http"
)

// HX-Trigger -> allows you to trigger client-side events
// HX-Trigger-After-Settle -> allows you to trigger client-side events after the settle step
// HX-Trigger-After-Swap -> allows you to trigger client-side events after the swap step
//

func HtmxTrigger(eventName string, w http.ResponseWriter) {
	w.Header().Add("HX-Trigger", eventName)
}

func HtmxTriggerAfterSettle(eventName string, w http.ResponseWriter) {
	w.Header().Add("HX-Trigger-After-Settle", eventName)
}

func HtmxTriggerAfterSwap(eventName string, w http.ResponseWriter) {
	w.Header().Add("HX-Trigger-After-Swap", eventName)
}
