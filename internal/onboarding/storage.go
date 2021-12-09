package onboarding

import (
	"fmt"
	"strings"
)

type simpleStorage struct {
	m map[string]interface{}
}

var storage simpleStorage

func init() {
	storage = simpleStorage{
		m: make(map[string]interface{}),
	}
}

func (s *simpleStorage) Update(k string, v interface{}) {
	s.m[k] = v
}

func (s *simpleStorage) Read(k string) (interface{}, error) {
	if val, found := s.m[k]; found {
		return val, nil
	}

	return "", fmt.Errorf("%q not found in SimpleStorage", k)
}

func (s *simpleStorage) ReadAll(keyPrefix string) []interface{} {
	var vals []interface{}

	for k, v := range s.m {
		if strings.HasPrefix(k, keyPrefix) {
			vals = append(vals, v)
		}
	}

	return vals
}

func (s *simpleStorage) Delete(k string) {
	delete(s.m, k)
}

func (s *simpleStorage) BuildKey(apiName, configID string) string {
	return apiName + "-" + configID
}
