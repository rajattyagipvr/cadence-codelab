package restaurant

import (
	"net/http"
	common "trying/webserver/service"
)

func (h *RestaurantService) showOrders(w http.ResponseWriter, r *http.Request) {
	common.ViewHandler(w, r, &h.state)
}
