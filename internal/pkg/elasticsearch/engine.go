package elasticsearch

import (
	"airbnb-user-be/internal/pkg/log"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"airbnb-user-be/internal/pkg/env"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

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
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: id,
		Body:       strings.NewReader(payload),
		Refresh:    "true",
	}

	// Return an API response object from request
	res, err := req.Do(context.Background(), Client)
	if err != nil {
		log.Error(Instance, "indexing request error", err)
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		msg := fmt.Sprintf("%s error indexing for document id=%s", res.Status(), id)
		log.Error(Instance, "indexing with response error", errors.New(msg))
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
