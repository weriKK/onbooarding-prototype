package onboarding

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

type Onboarder struct {
	r *mux.Router
}

func New(r *mux.Router) *Onboarder {
	return &Onboarder{r: r}

}

func RegisterConfigurationPath(r *mux.Router, apiName string, externalCreateHandler http.HandlerFunc) {

	apiRoot := "/onboarding/v1"

	apiPath := path.Join(apiRoot, apiName)
	apiPathWithId := path.Join(apiPath, "{configurationId}")

	r.HandleFunc(apiPath, externalCreateHandler).Methods(http.MethodPost)
	r.HandleFunc(apiPath, defaultReadAllHandler).Methods(http.MethodGet)
	r.HandleFunc(apiPathWithId, defaultReadHandler).Methods(http.MethodGet)
	r.HandleFunc(apiPathWithId, defaultUpdateHandler).Methods(http.MethodPut)
	r.HandleFunc(apiPathWithId, defaultDeleteHandler).Methods(http.MethodDelete)
}
