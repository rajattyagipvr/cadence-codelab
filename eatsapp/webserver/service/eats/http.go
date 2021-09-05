package eats

import (
	common "trying/webserver/service"
	"go.uber.org/cadence/client"
	s "go.uber.org/cadence/.gen/go/shared"
	"net/http"
)

type (
	// EatsService implements the handler for requests sent
	// to the Eats http service
	EatsService struct {
		menu   *common.Menu
		client client.Client
	}

	// EatsOrderListPage models the data to be displayed in response to
	// GET requests to the Eats service.
	EatsOrderListPage struct {
		ShowOrderExistError bool
		Orders              *s.ListOpenWorkflowExecutionsResponse
	}
)

const (
	cadenceTaskList = "cadence-bistro"
)

// NewService returns a new EatsService instance
func NewService(c client.Client, menu *common.Menu) *EatsService {
	return &EatsService{
		client: c,
		menu:   menu,
	}
}

func (h *EatsService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// PLACEHOLDER IMPL
		http.Error(w, "Not Implemented", http.StatusInternalServerError)
	case "POST":
		h.create(w, r)
		// PLACEHOLDER IMPL
		//http.Error(w, "Not Implemented", http.StatusInternalServerError)
	default:
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
