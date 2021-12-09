package app

import (
	"encoding/json"
	"net/http"
	"oblib/internal/onboarding"
)

type MontePolicyConfigJSON struct {
	AfID                   string `json:"afId"`
	AllowLocationReporting bool   `json:"allowLocationReporting"`
}

func CreateMontePolicyHandler(w http.ResponseWriter, r *http.Request) {
	var payload MontePolicyConfigJSON
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	apiName := onboarding.GetConfigName(r)
	configID := payload.AfID
	storageKey := onboarding.BuildConfigKey(apiName, configID)

	onboarding.UpdateConfig(storageKey, payload)

	w.WriteHeader(http.StatusOK)
}
