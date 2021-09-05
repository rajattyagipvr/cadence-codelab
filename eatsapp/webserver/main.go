package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	//"github.com/rajattyagipvr/cadence-codelab/common"
	//"github.com/rajattyagipvr/cadence-codelab/eatsapp/webserver/service"
	//"github.com/rajattyagipvr/cadence-codelab/eatsapp/webserver/service/courier"
	//"github.com/rajattyagipvr/cadence-codelab/eatsapp/webserver/service/eats"
	//"github.com/rajattyagipvr/cadence-codelab/eatsapp/webserver/service/restaurant"
	"trying/helper"
	"trying/webserver/service"
	"trying/webserver/service/courier"
	"trying/webserver/service/eats"
	"trying/webserver/service/restaurant"


)

func main() {

	// runtime := common.NewRuntime()
	// workflowClient, err := runtime.Builder.BuildCadenceClient()
	// if err != nil {
	// 	panic(err)
	// }
	var h helper.SampleHelper
	h.SetupServiceConfig()
	workflowClient, err := h.Builder.BuildCadenceClient()
	if err != nil {
		h.Logger.Error("Failed to build cadence client.", zap.Error(err))
		panic(err)
	}
	service.LoadTemplates()

	restaurant := restaurant.NewService(workflowClient, "assets/data/menu.yaml")

	http.Handle("/restaurant", restaurant)
	http.Handle("/courier", courier.NewService(workflowClient))
	http.Handle("/eats-orders", eats.NewService(workflowClient, restaurant.GetMenu()))
	http.Handle("/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/eats-menu", func(w http.ResponseWriter, r *http.Request) {
		service.ViewHandler(w, r, restaurant.GetMenu())
	})
	// setup & start server
	http.HandleFunc("/bistro", func(w http.ResponseWriter, r *http.Request) {
		service.ViewHandler(w, r, nil)
	})

	fmt.Println("Starting Webserver")
	http.ListenAndServe(":8090", nil)
}
