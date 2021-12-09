package app

import (
	"fmt"
	"net/http"
	"onbooarding-prototype/internal/onboarding"
	"time"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Run() {
	for {
		fmt.Printf("%+v\n", onboarding.ReadAllConfigs("montepolicy"))
		time.Sleep(5 * time.Second)
	}
}
