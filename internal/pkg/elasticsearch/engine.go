package elasticsearch

import (
	"airbnb-user-be/internal/pkg/log"
	"encoding/json"
	"fmt"
	"strings"

	"airbnb-user-be/internal/pkg/env"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func CreateIndex(mapping interface{}, indexNames ...string) (*esapi.Response, error) {
	indexNames = append([]string{env.CONFIG.Elastic.MainIndex}, indexNames...)
	index := createIndex(indexNames...)

	return Client.Indices.Create(index, Client.Indices.Create.WithBody(esutil.NewJSONReader(mapping)))
}

func Send(body interface{}, indexNames ...string) {
	b, err := json.Marshal(body)
	if err != nil {
		return
	}
	payload := string(b)

	id, err := gonanoid.New()
	if err != nil {
		return
	}

	indexNames = append([]string{env.CONFIG.Elastic.MainIndex}, indexNames...)
	index := createIndex(indexNames...)

	res, err := Client.Create(index, id, strings.NewReader(payload))
	if err != nil {
		msg := fmt.Sprintf("%s error indexing for document id=%s", res.Status(), id)
		log.Error(Instance, msg, err)
		return
	}
}

func createIndex(names ...string) string {
	separator := env.CONFIG.Elastic.Separator
	if separator == "" {
		separator = "_"
	}
	return strings.Join(names, separator)
}
