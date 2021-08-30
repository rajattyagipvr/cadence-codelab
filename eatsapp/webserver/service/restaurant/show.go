package restaurant

import (
	common "github.com/rajattyagipvr/cadence-codelab/eatsapp/webserver/service"
	"net/http"
)

func (h *RestaurantService) showOrders(w http.ResponseWriter, r *http.Request) {
	common.ViewHandler(w, r, &h.state)
}
