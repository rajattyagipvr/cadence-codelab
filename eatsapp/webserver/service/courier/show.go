package courier

import (
	common "github.com/rajattyagipvr/cadence-codelab/eatsapp/webserver/service"
	"net/http"
)

func (h *CourierService) showJobs(w http.ResponseWriter, r *http.Request) {
	common.ViewHandler(w, r, &h.DeliveryQueue)
}
