package onboarding

import (
	"encoding/json"
	"net/http"
)

func defaultReadHandler(w http.ResponseWriter, r *http.Request) {
	k := storage.BuildKey(GetConfigName(r), GetConfigID(r))
	if v, err := storage.Read(k); err == nil {
		json.NewEncoder(w).Encode(v)
	} else {
		w.Write([]byte(err.Error()))
	}
}

func defaultDeleteHandler(w http.ResponseWriter, r *http.Request) {
	k := storage.BuildKey(GetConfigName(r), GetConfigID(r))
	storage.Delete(k)

	w.WriteHeader(http.StatusNoContent)
}

func defaultReadAllHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func defaultUpdateHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}
