package main

import (
	"net/http"
	"onbooarding-prototype/internal/app"
	"onbooarding-prototype/internal/onboarding"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/test", app.TestHandler)

	onboarding.RegisterConfigurationPath(r, "niddpolicy", app.CreateNiddPolicyHandler)
	onboarding.RegisterConfigurationPath(r, "montepolicy", app.CreateMontePolicyHandler)

	go app.Run()

	http.ListenAndServe("0.0.0.0:8080", r)
}
