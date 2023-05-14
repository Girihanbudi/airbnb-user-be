package elasticsearch

import (
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/log"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

const Instance string = "Elastic Search"

// global elastic search client declaration
var Client *elasticsearch.Client

func InitElasticSearch() {

	config := elasticsearch.Config{
		Addresses: env.CONFIG.Elastic.Addresses,
		Username:  env.CONFIG.Elastic.Username,
		Password:  env.CONFIG.Elastic.Password,
	}

	client, err := elasticsearch.NewClient(config)
	if err != nil {
		log.Fatal(Instance, "connection error", err)
	}

	log.Event(Instance, fmt.Sprintf("connected to %v", env.CONFIG.Elastic.Addresses))
	client = client
}
