package elasticsearch

import (
	"context"
	"github.com/olivere/elastic/v7"
	"ipw-clean-arch/internal/model"
	"strconv"
)

func NewElasticClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		panic(err)
	}
	return client, nil
}

func IndexUser(user model.User) error {
	client, err := NewElasticClient()
	if err != nil {
		return err
	}
	_, err = client.Index().
		Index("users").
		Id(strconv.Itoa(user.ID)).
		BodyJson(user).
		Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
