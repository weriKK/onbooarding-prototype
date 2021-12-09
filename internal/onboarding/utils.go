package onboarding

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

func GetConfigName(r *http.Request) (configName string) {
	if r.Method == http.MethodPost {
		configName = path.Base(r.RequestURI)
	} else {
		configName = path.Base(path.Dir(r.RequestURI))
	}

	return configName
}

func GetConfigID(r *http.Request) (configID string) {
	configID = mux.Vars(r)["configurationId"]

	return configID
}

func ReadAllConfigs(configKeyPrefix string) []interface{} {
	return storage.ReadAll(configKeyPrefix)
}

func UpdateConfig(configKey string, configValue interface{}) {
	storage.Update(configKey, configValue)

}

func BuildConfigKey(configName, configID string) string {
	return storage.BuildKey(configName, configID)
}
