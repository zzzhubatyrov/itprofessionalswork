package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/olivere/elastic/v7"
	"ipw-clean-arch/internal/model"
	"ipw-clean-arch/pkg/elasticsearch"
	"log"
	"net/http"
)

func (h *Handler) searchUser(c *fiber.Ctx) error {
	query := c.Params("tag")
	log.Println("Получен запрос:", query)
	client, err := elasticsearch.NewElasticClient()
	if err != nil {
		return err
	}
	searchResult, err := client.Search().
		Index("users").
		Query(elastic.NewMatchQuery("tag", query)).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Failed to search for users: %v", err))
	}
	var users []model.User
	for _, hit := range searchResult.Hits.Hits {
		var user model.User
		if err := json.Unmarshal(hit.Source, &user); err != nil {
			log.Printf("Failed to unmarshal ElasticSearch hit: %v", err)
			continue
		}
		users = append(users, user)
	}
	return c.JSON(users)
}

func (h *Handler) esData(c *fiber.Ctx) error {
	client, err := elasticsearch.NewElasticClient()
	if err != nil {
		log.Fatal(err)
	}

	// Поиск и получение всех документов в индексе myindex
	searchResult, err := client.Search().
		Index("users").
		Query(elastic.NewMatchAllQuery()).
		Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a total of %d documents in the index 'myindex'\n", searchResult.Hits.TotalHits.Value)
	for _, hit := range searchResult.Hits.Hits {
		fmt.Printf("Document ID: %s\n", hit.Id)
		fmt.Printf("Document Source: %s\n", hit.Source)
		fmt.Println()
	}
	return c.JSON("")
}
