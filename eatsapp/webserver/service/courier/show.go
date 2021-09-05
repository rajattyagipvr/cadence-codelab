package courier

import (
	common "trying/webserver/service"
	"net/http"
)

func (h *CourierService) showJobs(w http.ResponseWriter, r *http.Request) {
	common.ViewHandler(w, r, &h.DeliveryQueue)
}
